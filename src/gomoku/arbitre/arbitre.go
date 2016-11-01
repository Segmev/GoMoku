package arbitre

import (
	"gomoku/window"
	"strconv"
)

type Player struct {
	Name        string
	Points      int
	FiveAligned []window.Stone
	TwoAligned  []window.Stone
}

type GomokuGame struct {
	Players  [2]Player
	Turn     bool
	End      int
	GameType bool
}

func (game *GomokuGame) Restart(pane *window.Drawer) bool {
	for i := range pane.Board_res.Stones {
		for j := range pane.Board_res.Stones[i] {
			pane.Board_res.Stones[i][j].Visible = false
		}
	}
	game.End = 0
	game.Players[0].Points, game.Players[1].Points = 0, 0
	pane.Board_res.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Board_res.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.GameState = "menu"
	return true
}

func IsStoneAtPos(dat *window.Drawer, i, j int) bool {
	if i >= 0 && i <= 18 && j >= 0 && j <= 18 {
		return dat.Board_res.Stones[i][j].Visible
	}
	return false
}

func augmentPos(pos int) int {
	if pos > 0 {
		return pos + 1
	} else if pos < 0 {
		return pos - 1
	}
	return pos
}

func CheckAlignement(dat *window.Drawer, stone *window.Stone, i, j, lim, ite int, del bool) bool {
	if IsStoneAtPos(dat, stone.Infos.Ipos+i, stone.Infos.Jpos+j) {
		if del && ite < lim && dat.Board_res.Stones[stone.Infos.Ipos+i][stone.Infos.Jpos+j].Color != stone.Color {
			iniI, iniJ := i, j
			i = augmentPos(i)
			j = augmentPos(j)
			if CheckAlignement(dat, stone, i, j, lim, ite+1, del) {
				if del {
					dat.Board_res.Stones[stone.Infos.Ipos+iniI][stone.Infos.Jpos+iniJ].Visible = false
				}
				return true
			}
		} else if !del && ite < lim && dat.Board_res.Stones[stone.Infos.Ipos+i][stone.Infos.Jpos+j].Color ==
			stone.Color {
			i = augmentPos(i)
			j = augmentPos(j)
			if CheckAlignement(dat, stone, i, j, lim, ite+1, del) {
				return true
			}
		} else if ite == lim && dat.Board_res.Stones[stone.Infos.Ipos+i][stone.Infos.Jpos+j].Color == stone.Color {
			return true
		}
	}
	return false
}

func TakeTwoStones(dat *window.Drawer, game *GomokuGame, stone *window.Stone) bool {
	hap := false
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if CheckAlignement(dat, stone, i, j, 2, 0, true) {
				if stone.Color {
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

func getInfosNbStonesDirection(dat *window.Drawer, st *window.Stone, color bool, i, j int) int {
	cpt := 0
	for IsStoneAtPos(dat, st.Infos.Ipos+i, st.Infos.Jpos+j) &&
		dat.Board_res.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Color == color {
		cpt += 1
		i, j = augmentPos(i), augmentPos(j)
	}
	return cpt
}

func CheckBreakable(dat *window.Drawer, stone *window.Stone) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) && stone.Infos.TeamSt[1+j][1+i] == 1 &&
				((IsStoneAtPos(dat, stone.Infos.Ipos+i, stone.Infos.Jpos+j) &&
					dat.Board_res.Stones[stone.Infos.Ipos+i+i][stone.Infos.Jpos+j+j].Color != stone.Color) ||
					stone.Infos.OppoSt[1+(-1*j)][1+(-1*i)] >= 1) {
				stone.Infos.Breakable = true
				return true
			}
		}
	}
	stone.Infos.Breakable = false
	return false
}

func CheckWinAlignment(dat *window.Drawer, game *GomokuGame, color bool) bool {
	ret := false
	for x := range dat.Board_res.Stones {
		for y := range dat.Board_res.Stones[x] {
			totOpp, totTeam := 0, 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && dat.Board_res.Stones[x][y].Visible {
						if dat.Board_res.Stones[x][y].Color == color {
							dat.Board_res.Stones[x][y].Infos.TeamSt[1+j][1+i] =
								getInfosNbStonesDirection(dat, dat.Board_res.Stones[x][y],
									dat.Board_res.Stones[x][y].Color, i, j)
							totTeam += dat.Board_res.Stones[x][y].Infos.TeamSt[1+j][1+i]
							CheckBreakable(dat, dat.Board_res.Stones[x][y])
							if !dat.Board_res.Stones[x][y].Infos.Breakable &&
								CheckAlignement(dat, dat.Board_res.Stones[x][y], i, j, 3, 0, false) {
								ret = true
							} else {
								dat.Board_res.Stones[x][y].Infos.OppoSt[1+j][1+i] =
									getInfosNbStonesDirection(dat, dat.Board_res.Stones[x][y],
										!dat.Board_res.Stones[x][y].Color, i, j)
								totOpp += dat.Board_res.Stones[x][y].Infos.OppoSt[1+j][1+i]
							}
						}
					}
				}
			}
			dat.Board_res.Stones[x][y].Infos.TeamSt[1][1] = totTeam
			dat.Board_res.Stones[x][y].Infos.OppoSt[1][1] = totOpp
		}
	}
	return ret
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
	for i := range dat.Board_res.Stones {
		for j := range dat.Board_res.Stones[i] {
			if x >= dat.Board_res.Stones[i][j].X &&
				x <= dat.Board_res.Stones[i][j].X+size*2 &&
				y >= dat.Board_res.Stones[i][j].Y &&
				y <= dat.Board_res.Stones[i][j].Y+size*2 {
				return dat.Board_res.Stones[i][j]
			}
		}
	}
	return nil
}

func HasTakenEnoughStones(pane *window.Drawer, game *GomokuGame) {
	if game.Players[0].Points >= 10 {
		game.End = 2
		pane.WinnerColor = true
	}
	if game.Players[1].Points >= 10 {
		game.End = 2
		pane.WinnerColor = false
	}
}

func GamePlay(pane *window.Drawer, game *GomokuGame, x, y, size int) {
	if game.End != 2 {
		st := IsStoneHere(pane, x, y, size)
		if st != nil && !st.Visible {
			st.Color = game.Turn
			st.Visible = true
			if TakeTwoStones(pane, game, st) {
				if game.End == 1 {
					t := CheckWinAlignment(pane, game, !game.Turn)
					if !t {
						game.End = 0
					}
				}
			}
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
