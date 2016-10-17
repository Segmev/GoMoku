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
	Turn		bool
	End		int
	GameType	bool
}

func	(game *GomokuGame) Restart(pane *window.Drawer) bool {
	for i := range(pane.Stones) {
		for j := range(pane.Stones[i]) {
			pane.Stones[i][j].Visible = false
		}
	}
	game.Players[0].Points, game.Players[1].Points = 0, 0
	pane.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	return true
}

func	IsStoneAtPos(dat *window.Drawer, i, j int) bool {
	if i >= 0 && i <= 18 && j >= 0 && j <= 18 {
		return dat.Stones[i][j].Visible
	}
	return false
}

func	CheckAlignement(dat *window.Drawer, stone *window.Stone, i, j, lim, ite int, del bool) bool {
	if IsStoneAtPos(dat, stone.Ipos + i, stone.Jpos + j) {
		if del && ite < lim && dat.Stones[stone.Ipos + i][stone.Jpos + j].White != stone.White {
			iniI, iniJ := i, j
			if i > 0 { i++ } else if i < 0 { i-- }
			if j > 0 { j++ } else if j < 0 { j-- }
			if CheckAlignement(dat, stone, i, j, lim, ite + 1, del) {
				if del { dat.Stones[stone.Ipos + iniI][stone.Jpos + iniJ].Visible = false }
				return true
			}
		} else if !del && ite < lim && dat.Stones[stone.Ipos + i][stone.Jpos + j].White == stone.White {
			if i > 0 { i++ } else if i < 0 { i-- }
			if j > 0 { j++ } else if j < 0 { j-- }
			if CheckAlignement(dat, stone, i, j, lim, ite + 1, del) {
				return true
			}
		} else if ite == lim && dat.Stones[stone.Ipos + i][stone.Jpos + j].White == stone.White {
			return true
		}
	}
	return false
}

func	TakeTwoStones(dat *window.Drawer, game *GomokuGame, stone *window.Stone) bool {
	hap := false
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if CheckAlignement(dat, stone, i, j, 2, 0, true) {
				if stone.White {
					game.Players[0].Points += 2
				} else {
					game.Players[1].Points += 2
				}
				hap = true
			}
		}
	}
	return hap
}

func	CheckWinAlignment(dat *window.Drawer, game *GomokuGame, color bool) bool {
	for x := range(dat.Stones) {
		for y := range(dat.Stones[x]) {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) &&
						dat.Stones[x][y].Visible &&
						dat.Stones[x][y].White == color &&
						CheckAlignement(dat, dat.Stones[x][y], i, j, 3, 0, false) {
						return true
					}
				}
			}
		}
	}
	return false
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
	if game.End != 2 {
		st := IsStoneHere(pane, x, y, size)
		if st != nil && !st.Visible {
			st.White = game.Turn
			if TakeTwoStones(pane, game, st) {
				if game.End == 1 {
					t := CheckWinAlignment(pane, game, !game.Turn)
					if !t {
						game.End = 0
					}
				}				
			}
			AppearStone(pane, x, y, size)
			if CheckWinAlignment(pane, game, game.Turn) {
				game.End = 1
			} else if game.End == 1 {
				game.End = 2
			}
			game.Turn = !game.Turn
		}
	}
	
	if game.End == 2 {
		gfx.CloseDisplay()
		os.Exit(0)
	}
	
	if game.Players[0].Points >= 10 || game.Players[1].Points >= 10 { game.End = 2 }
	pane.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.Turn = game.Turn
}
