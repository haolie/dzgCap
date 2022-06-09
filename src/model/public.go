package model

import (
	"image/color"
	"time"
)

func GetLDColor() color.Color {
	return color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	}
}

func GetMinTime() time.Time {
	return time.Date(2000, 1, 1, 1, 1, 1, 1, time.Now().Location())
}

func GetMaxTime() time.Time {
	return time.Date(2090, 1, 1, 1, 1, 1, 1, time.Now().Location())
}
