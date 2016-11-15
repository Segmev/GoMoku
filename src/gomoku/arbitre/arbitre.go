package arbitre

import (
	"gomoku/bmap"
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
	for i := range pane.BoardRes.Stones {
		for j := range pane.BoardRes.Stones[i] {
			pane.BoardRes.Stones[i][j].Visible = false
			bmap.ClearStone(i, j)
		}
	}
	game.End = 0
	pane.EndRes.DrawEnd = false
	game.Players[0].Points, game.Players[1].Points = 0, 0
	pane.BoardRes.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.BoardRes.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
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
		return bmap.IsVisible(i, j)
		//return dat.BoardRes.Stones[i][j].Visible
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
		if del && ite < lim && bmap.IsWhite(stone.Infos.Ipos+i, stone.Infos.Jpos+j) != stone.Color {
			iniI, iniJ := i, j
			i, j = augmentPos(i), augmentPos(j)
			if CheckAlignement(dat, stone, i, j, lim, ite+1, del) {
				if del {
					dat.BoardRes.Stones[stone.Infos.Ipos+iniI][stone.Infos.Jpos+iniJ].Visible = false
					bmap.SetVisibility(stone.Infos.Ipos+iniI, stone.Infos.Jpos+iniJ, false)
				}
				return true
			}
		} else if !del && ite < lim && bmap.IsWhite(stone.Infos.Ipos+i, stone.Infos.Jpos+j) == stone.Color {
			i, j = augmentPos(i), augmentPos(j)
			if CheckAlignement(dat, stone, i, j, lim, ite+1, del) {
				return true
			}
		} else if ite == lim && bmap.IsWhite(stone.Infos.Ipos+i, stone.Infos.Jpos+j) == stone.Color {
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
		bmap.IsWhite(st.Infos.Ipos+i, st.Infos.Jpos+j) == color {
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
					bmap.IsWhite(stone.Infos.Ipos+i+i, stone.Infos.Jpos+j+j) != stone.Color) ||
					stone.Infos.OppoSt[1+(-1*j)][1+(-1*i)] >= 1) {
				stone.Infos.Breakable = true
				bmap.SetBreakable(stone.Infos.Ipos, stone.Infos.Jpos, true)
				return true
			}
		}
	}
	stone.Infos.Breakable = false
	bmap.SetBreakable(stone.Infos.Ipos, stone.Infos.Jpos, false)
	return false
}

func ResetTeamInfos(dat *window.Drawer, color bool) {
	for x := range dat.BoardRes.Stones {
		for _, st := range dat.BoardRes.Stones[x] {
			if st.Color == color {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						st.Infos.TeamSt[1+i][1+j], st.Infos.OppoSt[1+i][1+j] = 0, 0
					}
				}
				st.Infos.Breakable = false
				bmap.SetBreakable(st.Infos.Ipos, st.Infos.Jpos, false)
			}
		}
	}
}

func UpdateInfos(dat *window.Drawer, game *GomokuGame, color bool) {
	//ResetTeamInfos(dat, color)
	for x := range dat.BoardRes.Stones {
		for y := range dat.BoardRes.Stones[x] {
			totOpp, totTeam := 0, 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(dat, x, y) {
						if bmap.IsWhite(x, y) == color {
							dat.BoardRes.Stones[x][y].Infos.TeamSt[1+j][1+i] =
								getInfosNbStonesDirection(dat, dat.BoardRes.Stones[x][y],
									bmap.IsWhite(x, y), i, j)
							bmap.SetNbTeamAt(x, y, 1+j, 1+i, uint64(dat.BoardRes.Stones[x][y].Infos.TeamSt[1+j][1+i]))
							totTeam += dat.BoardRes.Stones[x][y].Infos.TeamSt[1+j][1+i]
						} else {
							dat.BoardRes.Stones[x][y].Infos.OppoSt[1+j][1+i] =
								getInfosNbStonesDirection(dat, dat.BoardRes.Stones[x][y],
									!bmap.IsWhite(x, y), i, j)
							bmap.SetNbOppoAt(x, y, 1+j, 1+i, uint64(dat.BoardRes.Stones[x][y].Infos.OppoSt[1+j][1+i]))
							totOpp += dat.BoardRes.Stones[x][y].Infos.OppoSt[1+j][1+i]
						}
					}
				}
			}
			bmap.SetNbTeamAt(x, y, 1, 1, uint64(totTeam))
			bmap.SetNbOppoAt(x, y, 1, 1, uint64(totOpp))
			dat.BoardRes.Stones[x][y].Infos.TeamSt[1][1] = totTeam
			dat.BoardRes.Stones[x][y].Infos.OppoSt[1][1] = totOpp
		}
	}
}

func CheckWinAlignment(dat *window.Drawer, game *GomokuGame, color bool) {
	game.Players[GetPlayerNb(game, color)].FiveAligned = game.Players[GetPlayerNb(game, color)].FiveAligned[:0]
	for x := range dat.BoardRes.Stones {
		for y := range dat.BoardRes.Stones[x] {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(dat, x, y) {
						if dat.BoardRes.Stones[x][y].Color == color {
							if CheckAlignement(dat, dat.BoardRes.Stones[x][y], i, j, 3, 0, false) {
								StonesTab := [5]*window.Stone{
									dat.BoardRes.Stones[x][y],
									dat.BoardRes.Stones[x+i][y+j],
									dat.BoardRes.Stones[x+i+i][y+j+j],
									dat.BoardRes.Stones[x+i+i+i][y+j+j+j],
									dat.BoardRes.Stones[x+i+i+i+i][y+j+j+j+j]}
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
		bmap.SetVisibility(stone.Infos.Ipos, stone.Infos.Jpos, true)
		return true
	}
	return false
}

func IsStoneHere(dat *window.Drawer, x, y, size int) *window.Stone {
	for i := range dat.BoardRes.Stones {
		for _, st := range dat.BoardRes.Stones[i] {
			if x >= st.X && x <= st.X+size*2 &&
				y >= st.Y && y <= st.Y+size*2 {
				return st
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
					if st.Color == dat.BoardRes.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Color {
						if dat.BoardRes.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Infos.TeamSt[1+j][1+i] <= 1 {
							if dat.BoardRes.Stones[st.Infos.Ipos+i][st.Infos.Jpos+j].Infos.TeamSt[1+j][1+i] == 1 ||
								(IsStoneAtPos(dat, st.Infos.Ipos+i+i+i, st.Infos.Jpos+j+j+j) &&
									st.Color == dat.BoardRes.Stones[st.Infos.Ipos+i+i+i][st.Infos.Jpos+j+j+j].Color &&
									dat.BoardRes.Stones[st.Infos.Ipos+i+i+i][st.Infos.Jpos+j+j+j].Infos.TeamSt[1+j][1+i] == 0) {
								cpt += 1
							}
						}
					}
				} else if IsStoneAtPos(dat, st.Infos.Ipos+i+i, st.Infos.Jpos+j+j) {
					if st.Color == dat.BoardRes.Stones[st.Infos.Ipos+i+i][st.Infos.Jpos+j+j].Color {
						if dat.BoardRes.Stones[st.Infos.Ipos+i+i][st.Infos.Jpos+j+j].Infos.TeamSt[1+j][1+i] == 1 {
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
	for _, stonesCol := range pane.BoardRes.Stones {
		for _, stone := range stonesCol {
			if stone.Visible == false {
				state = false
			}
		}
	}
	if state == true {
		game.End = 2
		pane.EndRes.DrawEnd = true
	}
}

func CheckBreakableAlign(dat *window.Drawer, game *GomokuGame, color bool) bool {
	if dat.OptionsRes.Op1 == false {
		return true
	}
	tot := 0
	for _, line := range game.Players[GetPlayerNb(game, color)].FiveAligned {
		cpt := 0
		for _, st := range line {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) &&
						st.Infos.TeamSt[i+1][j+1] == 1 &&
						(st.Infos.OppoSt[1+(-1*i)][1+(-1*j)] >= 1 ||
							(dat.BoardRes.Stones[st.Infos.Ipos+j][st.Infos.Jpos+i].Infos.OppoSt[1+i][1+j] >= 1 &&
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
			bmap.SetColor(st.Infos.Ipos, st.Infos.Jpos, game.Turn)
			if pane.OptionsRes.Op2 && ThreeBlockNear(pane, game, st) == 2 {
				pane.BoardRes.BadX, pane.BoardRes.BadY = st.X, st.Y
				return
			} else {
				pane.BoardRes.BadX, pane.BoardRes.BadY = 0, 0
			}
			st.Visible = true
			bmap.SetVisibility(st.Infos.Ipos, st.Infos.Jpos, true)
			TakeTwoStones(pane, game, st)
			UpdateInfos(pane, game, game.Turn)
			CheckWinAlignment(pane, game, game.Turn)
			if len(game.Players[GetPlayerNb(game, game.Turn)].FiveAligned) > 0 {
				if CheckBreakableAlign(pane, game, st.Color) {
					game.End = 2
					pane.WinnerColor = game.Turn
				}
			}
			game.Turn = !game.Turn
		}
	}
	HasTakenEnoughStones(pane, game)
	isDraw(pane, game)
	if game.End == 2 {
		pane.GameState = "end"
	}
	pane.BoardRes.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.BoardRes.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.Turn = game.Turn
}
