package arbitre

import (
	"gomoku/window"
	//	"strings"
	// "fmt"
	"strconv"
	"github.com/gtalent/starfish/gfx"
	"os"
)

type	Player struct {
	Name	string
	Points	int
	
}

type	GomokuGame struct {
	Players	[2]Player
	Turn	bool
	End	bool
}

func	IsStoneAtPos(dat *window.Drawer, i, j int) bool {
	if i >= 0 && i <= 18 && j >= 0 && j <= 18 {
		return dat.Stones[i][j].Visible
	}
	return false
}

func	CheckAlignement(dat *window.Drawer, stone *window.Stone, i, j, ite int, del bool) bool {
	if IsStoneAtPos(dat, stone.Ipos + i, stone.Jpos + j) {
		if ite < 2 && dat.Stones[stone.Ipos + i][stone.Jpos + j].White != stone.White {
			iniI, iniJ := i, j
			if i > 0 { i++ } else if i < 0 { i-- }
			if j > 0 { j++ } else if j < 0 { j-- }
			if CheckAlignement(dat, stone, i, j, ite + 1, del) {
				if del { dat.Stones[stone.Ipos + iniI][stone.Jpos + iniJ].Visible = false }
				return true
			}
		} else if ite == 2 && dat.Stones[stone.Ipos + i][stone.Jpos + j].White == stone.White {
			return true
		}
	}
	return false
}

func	TakeTwoStones(dat *window.Drawer, game *GomokuGame, stone *window.Stone) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if CheckAlignement(dat, stone, i, j, 0, true) {
				if stone.White { game.Players[0].Points += 2
				} else { game.Players[1].Points += 2 }
			}
		}
	}
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
			TakeTwoStones(pane, game, st)
			AppearStone(pane, x, y, size)
			game.Turn = !game.Turn
		}
	} else {
		gfx.CloseDisplay()
		os.Exit(0)

	}
	if game.Players[0].Points < 10 && game.Players[1].Points < 10 { game.End = true }
	pane.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
}
