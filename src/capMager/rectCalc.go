package capMager

import (
	"image"
	"image/color"

	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

const con_dis_num = 50

type line struct {
	x1, y1, x2, y2 int
}

// 纵向识别线条
func lineFromImageL(img image.Image, c color.Color) (lines []line) {

	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y
	pMap := make(map[int][]int, width/con_dis_num)
	for i := 1; i*con_dis_num < width; i++ {
		for j := 0; j < height; j++ {
			if imageTool.CompareColor(img.At(i*con_dis_num, j), c) {
				if _, exists := pMap[j]; !exists {
					pMap[j] = make([]int, 0, 8)
				}

				pMap[j] = append(pMap[j], i*con_dis_num)
			}
		}
	}

	for y, pList := range pMap {
		if len(pList) < 2 {
			continue
		}

		lines = append(lines, lineFromPListL(img, c, y, pList)...)
	}

	return
}

func lineFromPListL(img image.Image, c color.Color, y int, plist []int) (lines []line) {

	imgSize := img.Bounds().Size()

	var tempLine *line
	for _, x := range plist {
		// y 已在上一条线以内
		if tempLine != nil && tempLine.x2 >= x {
			continue
		}

		tempLine = &line{x, y, x, y}

		// 寻找线条开始点
		for i := 1; x-i >= 0; i++ {
			// 到达边界或颜色不同 停止查找
			if !imageTool.CompareColor(img.At(x-i, y), c) {
				break
			}

			tempLine.x1 = x - i
		}

		// 寻找线条开终点
		for i := 1; x+i < imgSize.X; i++ {
			// 到达边界或颜色不同 停止查找
			if !imageTool.CompareColor(img.At(x+i, y), c) {
				break
			}

			tempLine.x2 = x + i
		}

		// 开始点和结束点相同 不成线条
		if tempLine.x1 == tempLine.x2 {
			tempLine = nil
			continue
		}

		lines = append(lines, *tempLine)
	}

	return
}

// 横向识别线条
func lineFromImageH(img image.Image, c color.Color) []line {
	return nil
}

func FindMinRect(img image.Image, c color.Color) (rect model.Rect, exists bool) {
	list, exists := FindRect(img, c)
	if !exists {
		return
	}

	rect = list[0]
	for i := 1; i < len(list); i++ {
		if rect.W <= list[i].W && rect.H <= list[i].H {
			continue
		}

		rect = list[i]
	}

	return rect, true
}

func FindRect(img image.Image, c color.Color) (rectList []model.Rect, exists bool) {
	// 线条
	lineList := lineFromImageL(img, c)
	if len(lineList) < 2 {
		return nil, false
	}

	// 通过对比两条垂直位移的线条确定rect
	for i := 0; i < len(lineList)-1; i++ {
		for j := i + 1; j < len(lineList); j++ {

			if lineList[i].x1 == lineList[j].x1 && lineList[i].x2 == lineList[j].x2 {
				tempRect := model.Rect{
					X: lineList[i].x1,
					Y: lineList[i].y1,
					W: lineList[i].x2 - lineList[i].x1 + 1,
					H: lineList[j].y1 - lineList[i].y1 + 1,
				}
				if tempRect.H < 0 {
					tempRect.Y = lineList[j].y1
					tempRect.H = lineList[i].y1 - lineList[j].y1 + 1
				}

				rectList = append(rectList, tempRect)
				exists = true
			}
		}
	}

	return
}
