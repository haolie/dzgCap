package gameArea

import (
	"dzgCap/src/model"
)

func CreateFactory(modelKey string, gt model.IGameTable) func(r model.Rect) model.IGameArea {
	return func(r model.Rect) model.IGameArea {
		return &GameArea{
			gameModelKey: modelKey,
			gTable:       gt,
			rect:         r,
		}
	}
}
