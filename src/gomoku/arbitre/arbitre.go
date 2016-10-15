package arbitre

import (
	"gomoku/window"
//	"strings"
)

type Player struct {
	Name	string
	Points	int
	
}

type GomokuGame struct {
	Player1	Player
	Player2	Player
	Turn	bool
	End	bool
}


func AppearStone(dat *window.Drawer, x, y, size int) bool {
	stone := IsStoneHere(dat, x, y, size)
	if stone != nil {
		stone.Visible = true
		return true
	}
	return false
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

func GamePlay(pane *window.Drawer, game *GomokuGame, x, y, size int) {
	if !game.End {
		st := IsStoneHere(pane, x, y, size)
		if st  != nil {
			st.White = game.Turn
			AppearStone(pane, x, y, size)
			game.Turn = !game.Turn
		}
	}
}
