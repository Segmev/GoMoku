package arbitre

import (
	"gomoku/window"
	//	"strings"
	// "fmt"
	"strconv"
	// "github.com/gtalent/starfish/gfx"
	// "os"
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
	for i := range(pane.Board_res.Stones) {
		for j := range(pane.Board_res.Stones[i]) {
			pane.Board_res.Stones[i][j].Visible = false
		}
	}
	game.Players[0].Points, game.Players[1].Points = 0, 0
	pane.Board_res.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Board_res.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.GameState = "menu"
	return true
}

func	IsStoneAtPos(dat *window.Drawer, i, j int) bool {
	if i >= 0 && i <= 18 && j >= 0 && j <= 18 {
		return dat.Board_res.Stones[i][j].Visible
	}
	return false
}

func	CheckAlignement(dat *window.Drawer, stone *window.Stone, i, j, lim, ite int, del bool) bool {
	if IsStoneAtPos(dat, stone.Ipos + i, stone.Jpos + j) {
		if del && ite < lim && dat.Board_res.Stones[stone.Ipos + i][stone.Jpos + j].White != stone.White {
			iniI, iniJ := i, j
			if i > 0 { i++ } else if i < 0 { i-- }
			if j > 0 { j++ } else if j < 0 { j-- }
			if CheckAlignement(dat, stone, i, j, lim, ite + 1, del) {
				if del { dat.Board_res.Stones[stone.Ipos + iniI][stone.Jpos + iniJ].Visible = false }
				return true
			}
		} else if !del && ite < lim && dat.Board_res.Stones[stone.Ipos + i][stone.Jpos + j].White ==
			stone.White {
			if i > 0 { i++ } else if i < 0 { i-- }
			if j > 0 { j++ } else if j < 0 { j-- }
			if CheckAlignement(dat, stone, i, j, lim, ite + 1, del) {
				return true
			}
		} else if ite == lim && dat.Board_res.Stones[stone.Ipos + i][stone.Jpos + j].White == stone.White {
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
	for x := range(dat.Board_res.Stones) {
		for y := range(dat.Board_res.Stones[x]) {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) &&
						dat.Board_res.Stones[x][y].Visible &&
						dat.Board_res.Stones[x][y].White == color &&
						CheckAlignement(dat, dat.Board_res.Stones[x][y], i, j, 3, 0, false) {
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
	for i := range(dat.Board_res.Stones) {
		for j := range(dat.Board_res.Stones[i]) {
			if x >= dat.Board_res.Stones[i][j].X &&
				x <= dat.Board_res.Stones[i][j].X + size * 2 &&
				y >= dat.Board_res.Stones[i][j].Y &&
				y <= dat.Board_res.Stones[i][j].Y + size * 2 {
				return dat.Board_res.Stones[i][j]
			}
		}
	}
	return nil
}

func	HasTakenEnoughStones(pane *window.Drawer, game *GomokuGame) {
	if game.Players[0].Points >= 10 {
		game.End = 2
		pane.WinnerColor = false
	}
	if game.Players[1].Points >= 10 {
		game.End = 2
		pane.WinnerColor = true
	}
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
			if game.End == 1 {
				game.End = 2
			} else if CheckWinAlignment(pane, game, game.Turn) {
				game.End = 1
				pane.WinnerColor = game.Turn
			} 
			game.Turn = !game.Turn
		}
	}
	HasTakenEnoughStones(pane, game)
	if game.End == 2 {
		pane.GameState = "end"
	}
	pane.Board_res.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Board_res.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.Turn = game.Turn
}
