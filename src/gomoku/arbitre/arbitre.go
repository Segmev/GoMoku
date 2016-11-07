package arbitre

import (
	"gomoku/window"
	"strconv"
)

type Player struct {
	Name        string
	Points      int
	FiveAligned [][5]*window.Stone
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
	pane.End_res.DrawEnd = false
	game.Players[0].Points, game.Players[1].Points = 0, 0
	pane.Board_res.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Board_res.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.GameState = "menu"
	return true
}

func GetPlayerNb(game *GomokuGame, color bool) int {
	if color {
		return 1
	}
	return 0
}

func isElemInAlignedArray(s [][5]*window.Stone, e [5]*window.Stone) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
			i, j = augmentPos(i), augmentPos(j)
			if CheckAlignement(dat, stone, i, j, lim, ite+1, del) {
				if del {
					dat.Board_res.Stones[stone.Infos.Ipos+iniI][stone.Infos.Jpos+iniJ].Visible = false
				}
				return true
			}
		} else if !del && ite < lim && dat.Board_res.Stones[stone.Infos.Ipos+i][stone.Infos.Jpos+j].Color ==
			stone.Color {
			i, j = augmentPos(i), augmentPos(j)
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
					IsStoneAtPos(dat, stone.Infos.Ipos+i+i, stone.Infos.Jpos+j+j) &&
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

func ResetTeamInfos(dat *window.Drawer, color bool) {
	for x := range dat.Board_res.Stones {
		for _, st := range dat.Board_res.Stones[x] {
			if st.Color == color {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						st.Infos.TeamSt[1+i][1+j], st.Infos.OppoSt[1+i][1+j] = 0, 0
					}
				}
				st.Infos.Breakable = false
			}
		}
	}
}

func UpdateInfos(dat *window.Drawer, game *GomokuGame, color bool) {
	//ResetTeamInfos(dat, color)
	for x := range dat.Board_res.Stones {
		for y := range dat.Board_res.Stones[x] {
			totOpp, totTeam := 0, 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(dat, x, y) {
						if dat.Board_res.Stones[x][y].Color == color {
							dat.Board_res.Stones[x][y].Infos.TeamSt[1+j][1+i] =
								getInfosNbStonesDirection(dat, dat.Board_res.Stones[x][y],
									dat.Board_res.Stones[x][y].Color, i, j)
							totTeam += dat.Board_res.Stones[x][y].Infos.TeamSt[1+j][1+i]
						} else {
							dat.Board_res.Stones[x][y].Infos.OppoSt[1+j][1+i] =
								getInfosNbStonesDirection(dat, dat.Board_res.Stones[x][y],
									!dat.Board_res.Stones[x][y].Color, i, j)
							totOpp += dat.Board_res.Stones[x][y].Infos.OppoSt[1+j][1+i]
						}
					}
				}
			}
			dat.Board_res.Stones[x][y].Infos.TeamSt[1][1] = totTeam
			dat.Board_res.Stones[x][y].Infos.OppoSt[1][1] = totOpp
		}
	}
}

func CheckWinAlignment(dat *window.Drawer, game *GomokuGame, color bool) {
	game.Players[GetPlayerNb(game, color)].FiveAligned = game.Players[GetPlayerNb(game, color)].FiveAligned[:0]
	for x := range dat.Board_res.Stones {
		for y := range dat.Board_res.Stones[x] {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(dat, x, y) {
						if dat.Board_res.Stones[x][y].Color == color {
							if CheckAlignement(dat, dat.Board_res.Stones[x][y], i, j, 3, 0, false) {
								StonesTab := [5]*window.Stone{
									dat.Board_res.Stones[x][y],
									dat.Board_res.Stones[x+i][y+j],
									dat.Board_res.Stones[x+i+i][y+j+j],
									dat.Board_res.Stones[x+i+i+i][y+j+j+j],
									dat.Board_res.Stones[x+i+i+i+i][y+j+j+j+j]}
								game.Players[GetPlayerNb(game, color)].FiveAligned =
									append(game.Players[GetPlayerNb(game, color)].FiveAligned, StonesTab)
							}
						}
					}
				}
			}
		}
	}
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
	} else if game.Players[1].Points >= 10 {
		game.End = 2
		pane.WinnerColor = false
	}
}

func ThreeBlockNear(dat *window.Drawer, game *GomokuGame, st *window.Stone) int {
	cpt := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				if IsStoneAtPos(dat, st.Infos.Ipos+i, st.Infos.Jpos+j) {
					if st.Color == dat.Board_res.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Color {
						if dat.Board_res.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Infos.TeamSt[1+j][1+i] <= 1 {
							if dat.Board_res.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Infos.TeamSt[1+j][1+i] == 1 ||
								(IsStoneAtPos(dat, st.Infos.Ipos+i+i+i, st.Infos.Jpos+j+j+j) &&
									st.Color == dat.Board_res.Stones[st.Infos.Ipos+i+i+i][st.Infos.Jpos+j+j+j].Color &&
									dat.Board_res.Stones[st.Infos.Ipos+i+i+i][st.Infos.Jpos+j+j+j].Infos.TeamSt[1+j][1+i] == 0) {
								cpt += 1
							}
						}
					}
				} else if IsStoneAtPos(dat, st.Infos.Ipos+i+i, st.Infos.Jpos+j+j) {
					if st.Color == dat.Board_res.Stones[st.Infos.Ipos+i+i][st.Infos.Jpos+j+j].Color {
						if dat.Board_res.Stones[st.Infos.Ipos+i+i][st.Infos.Jpos+j+j].Infos.TeamSt[1+j][1+i] == 1 {
							cpt += 1
						}
					}
				}
			}
		}
	}
	return cpt
}

func isDraw(pane *window.Drawer, game *GomokuGame) {
	state := true
	for _, stonesCol := range pane.Board_res.Stones {
		for _, stone := range stonesCol {
			if stone.Visible == false {
				state = false
			}
		}
	}
	if state == true {
		game.End = 2
		pane.End_res.DrawEnd = true
	}
}

func CheckBreakableAlign(dat *window.Drawer, game *GomokuGame, color bool) bool {
	tot := 0
	for _, line := range game.Players[GetPlayerNb(game, color)].FiveAligned {
		cpt := 0
		for _, st := range line {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) &&
						st.Infos.TeamSt[i+1][j+1] == 1 &&
						(st.Infos.OppoSt[1+(-1*i)][1+(-1*j)] >= 1 ||
							(dat.Board_res.Stones[st.Infos.Ipos+j][st.Infos.Jpos+i].Infos.OppoSt[1+i][1+j] >= 1 &&
								!IsStoneAtPos(dat, st.Infos.Ipos+(-1*j), st.Infos.Jpos+(-1*i)))) {
						st.Infos.Breakable = true
						cpt = 1
					}
				}
			}
		}
		tot += cpt
	}
	if tot < len(game.Players[GetPlayerNb(game, game.Turn)].FiveAligned) {
		return true
	}
	return false
}

func GamePlay(pane *window.Drawer, game *GomokuGame, x, y, size int) {
	if game.End != 2 {
		st := IsStoneHere(pane, x, y, size)
		if st != nil && !st.Visible {
			st.Color = game.Turn
			if ThreeBlockNear(pane, game, st) == 2 {
				pane.Board_res.BadX, pane.Board_res.BadY = st.X, st.Y
				return
			} else {
				pane.Board_res.BadX, pane.Board_res.BadY = 0, 0
			}
			st.Visible = true
			TakeTwoStones(pane, game, st)
			UpdateInfos(pane, game, game.Turn)
			CheckWinAlignment(pane, game, game.Turn)
			if len(game.Players[GetPlayerNb(game, game.Turn)].FiveAligned) > 0 &&
				CheckBreakableAlign(pane, game, st.Color) {
				game.End = 2
				pane.WinnerColor = game.Turn
			}
			game.Turn = !game.Turn
		}
	}
	HasTakenEnoughStones(pane, game)
	isDraw(pane, game)
	if game.End == 2 {
		pane.GameState = "end"
	}
	pane.Board_res.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.Board_res.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.Turn = game.Turn
}
