package arbitre

import (
	"gomoku/window"
	//	"strings"
	"fmt"
)

type	Player struct {
	Name	string
	Points	int
	
}

type	GomokuGame struct {
	Player1	Player
	Player2	Player
	Turn	bool
	End	bool
}

func	IsStoneAtPos(dat *window.Drawer, i, j int) bool {
	if i >= 0 && i <= 18 && j >= 0 && j <= 18 {
		return dat.Stones[i][j].Visible
	}
	return false
}

func	TakeTwoStones(dat *window.Drawer, stone *window.Stone) {
	
	fmt.Println(stone.Ipos, stone.Jpos, (IsStoneAtPos(dat, stone.Ipos, stone.Jpos + 1)))
}

func	AppearStone(dat *window.Drawer, x, y, size int) bool {
	stone := IsStoneHere(dat, x, y, size)
	if stone != nil {
		stone.Visible = true
		return true
	}
	return false
}

func	IsStoneHere(dat *window.Drawer, x, y, size int) *window.Stone {
	for i := range(dat.Stones) {
		for j := range(dat.Stones[i]) {
			if x >= dat.Stones[i][j].X && x <= dat.Stones[i][j].X + size * 2 &&
				y >= dat.Stones[i][j].Y && y <= dat.Stones[i][j].Y + size * 2 {
				return dat.Stones[i][j]
			}
		}
	}
	return nil
}

func	GamePlay(pane *window.Drawer, game *GomokuGame, x, y, size int) {
	if !game.End {
		st := IsStoneHere(pane, x, y, size)
		if st != nil && !st.Visible {
			st.White = game.Turn
			TakeTwoStones(pane, st)
			AppearStone(pane, x, y, size)
			game.Turn = !game.Turn
//			fmt.Println(st.Ipos, st.Jpos)
		}
	}
}
