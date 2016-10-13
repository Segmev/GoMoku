package arbitre

import (
	"gomoku/window"
)

func AppearStone(dat *window.Drawer, x, y, size int) {
	for i := range(dat.Stones) {		
		if x >= dat.Stones[i].X && x <= dat.Stones[i].X + size * 2 &&
			y >= dat.Stones[i].Y && y <= dat.Stones[i].Y + size * 2 {
			dat.Stones[i].Visible = true
		}
	}
}

func IsStoneHere(dat *window.Drawer, x, y, size int) *window.Stone {
	for i := range(dat.Stones) {
		if x >= dat.Stones[i].X && x <= dat.Stones[i].X + size * 2 &&
			y >= dat.Stones[i].Y && y <= dat.Stones[i].Y + size * 2 {
			return dat.Stones[i]
		}
	}
	return nil
}
