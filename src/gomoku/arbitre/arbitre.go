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
			bmap.ClearStone(&bmap.Map, i, j)
		}
	}
	game.End = 0
	pane.EndRes.DrawEnd = false
	game.Players[0].Points, game.Players[1].Points = 0, 0
	pane.BoardRes.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.BoardRes.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.GameState = "menu"
	pane.BoardRes.BadX, pane.BoardRes.BadY = 0, 0
	return true
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

func ValableCoor(i, j int) bool {
	return (i >= 0 && i <= 18 && j >= 0 && j <= 18)
}

func IsStoneAtPos(Map *[361]uint64, i, j int) bool {
	if ValableCoor(i, j) {
		return bmap.IsVisible(Map, i, j)
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

func CheckAlignement(Map *[361](uint64), x, y, i, j, lim, ite int, del bool) bool {
	if IsStoneAtPos(Map, x+i, y+j) {
		if del && ite < lim && bmap.IsWhite(Map, x+i, y+j) != bmap.IsWhite(Map, x, y) {
			iniI, iniJ := i, j
			i, j = augmentPos(i), augmentPos(j)
			if CheckAlignement(Map, x, y, i, j, lim, ite+1, del) {
				if del {
					bmap.ResetStone(Map, x+iniI, y+iniJ)
					//bmap.SetVisibility(Map, x+iniI, y+iniJ, false)
				}
				return true
			}
		} else if !del && ite < lim && bmap.IsWhite(Map, x+i, y+j) == bmap.IsWhite(Map, x, y) {
			i, j = augmentPos(i), augmentPos(j)
			if CheckAlignement(Map, x, y, i, j, lim, ite+1, del) {
				return true
			}
		} else if ite == lim && bmap.IsWhite(Map, x+i, y+j) == bmap.IsWhite(Map, x, y) {
			return true
		}
	}
	return false
}

func TakeTwoStones(Map *[361](uint64), game *GomokuGame, x, y int, color bool) bool {
	hap := false
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if CheckAlignement(Map, x, y, i, j, 2, 0, true) {
				if bmap.IsWhite(Map, x, y) {
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

func getInfosNbStonesDirection(Map *[361]uint64, x, y int, color bool, i, j int) int {
	cpt := 0
	for IsStoneAtPos(Map, x+i, y+j) &&
		bmap.IsWhite(Map, x+i, y+j) == color {
		cpt += 1
		i, j = augmentPos(i), augmentPos(j)
	}
	return cpt
}

func CheckBreakable(dat *window.Drawer, stone *window.Stone) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) && stone.Infos.TeamSt[1+j][1+i] == 1 &&
				((IsStoneAtPos(&bmap.Map, stone.Infos.Ipos+i, stone.Infos.Jpos+j) &&
					IsStoneAtPos(&bmap.Map, stone.Infos.Ipos+i+i, stone.Infos.Jpos+j+j) &&
					bmap.IsWhite(&bmap.Map, stone.Infos.Ipos+i+i, stone.Infos.Jpos+j+j) != stone.Color) ||
					stone.Infos.OppoSt[1+(-1*j)][1+(-1*i)] >= 1) {
				stone.Infos.Breakable = true
				bmap.SetBreakable(&bmap.Map, stone.Infos.Ipos, stone.Infos.Jpos, true)
				return true
			}
		}
	}
	stone.Infos.Breakable = false
	bmap.SetBreakable(&bmap.Map, stone.Infos.Ipos, stone.Infos.Jpos, false)
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
				bmap.SetBreakable(&bmap.Map, st.Infos.Ipos, st.Infos.Jpos, false)
			}
		}
	}
}

func UpdateInfos(Map *[361](uint64), game *GomokuGame, color bool) {
	//ResetTeamInfos(dat, color)
	for x := 0; x <= 18; x++ {
		for y := 0; y <= 18; y++ {
			UpdateThreeGroups(Map, game, x, y, color)
			totOpp, totTeam := 0, 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(&bmap.Map, x, y) {
						if bmap.IsWhite(&bmap.Map, x, y) == color {
							bmap.SetNbTeamAt(&bmap.Map, x, y, 1+j, 1+i, uint64(getInfosNbStonesDirection(Map, x, y,
								bmap.IsWhite(&bmap.Map, x, y), i, j)))
							totTeam += bmap.GetNbT(Map, x, y, 1+j, 1+i)

						} else {
							bmap.SetNbOppoAt(&bmap.Map, x, y, 1+j, 1+i, uint64(getInfosNbStonesDirection(Map, x, y,
								!bmap.IsWhite(&bmap.Map, x, y), i, j)))
							totOpp += bmap.GetNbO(Map, x, y, 1+j, 1+i)
						}
					}
				}
			}
			bmap.SetNbTeamAt(&bmap.Map, x, y, 1, 1, uint64(totTeam))
			bmap.SetNbOppoAt(&bmap.Map, x, y, 1, 1, uint64(totOpp))
		}
	}
}

func AppearStone(dat *window.Drawer, x, y, size int) bool {
	stone := IsStoneHere(dat, x, y, size)
	if stone != nil {
		stone.Visible = true
		bmap.SetVisibility(&bmap.Map, stone.Infos.Ipos, stone.Infos.Jpos, true)
		return true
	}
	return false
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

func ThreeBlockNear(Map *[361]uint64, game *GomokuGame, x, y int, color bool) bool {
	if UpdateThreeGroups(Map, game, x, y, color) {
		return true
	}
	ret := false
	bmap.SetVisibility(Map, x, y, true)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				a, b := i, j
				for c := 0; c <= 2; c++ {
					if IsStoneAtPos(Map, x+a, y+b) &&
						color == bmap.IsWhite(Map, x+a, y+b) {
						if UpdateThreeGroups(Map, game, x+a, y+b, color) {
							bmap.SetVisibility(&bmap.Map, x, y, false)
							ret = true
						}
					}
					a, b = augmentPos(a), augmentPos(b)
				}
			}
		}
	}
	return ret
}

func updateThreeGroupLoop(Map *[361]uint64, color bool, x, y, dirI, dirJ, cptHowManyThreeGroups, cptFourGroups int) (int, int) {
	cpt := 0
	if !(dirI == 0 && dirJ == 0) {
		i, j := dirI, dirJ
		end := 2
		if (i <= 0 && j <= 0) || (i == 1 && j == -1) {
			if IsStoneAtPos(Map, x-dirI, y-dirJ) && color == bmap.IsWhite(Map, x-dirI, y-dirJ) {
				cpt++
				end = 1
			}
		}
		for c := 0; c <= end; c++ {
			if IsStoneAtPos(Map, x+i, y+j) {
				if color == bmap.IsWhite(Map, x+i, y+j) {
					cpt++
				} else {
					c = 3
					cpt = 0
				}
			}
			i, j = augmentPos(i), augmentPos(j)
		}
		if cpt == 2 && !IsStoneAtPos(Map, i+x, j+y) {
			cptHowManyThreeGroups++
		}
		if cpt == 3 {
			cptFourGroups++
		}
	}
	return cptHowManyThreeGroups, cptFourGroups
}

func UpdateThreeGroups(Map *[361]uint64, game *GomokuGame, x, y int, color bool) bool {
	cptHowManyThreeGroups, cptFourGroups := 0, 0
	for dirI := -1; dirI <= 0; dirI++ {
		for dirJ := -1; dirJ <= 0; dirJ++ {
			cptHowManyThreeGroups, cptFourGroups =
				updateThreeGroupLoop(Map, color, x, y, dirI, dirJ, cptHowManyThreeGroups, cptFourGroups)
		}
	}
	cptHowManyThreeGroups, cptFourGroups =
		updateThreeGroupLoop(Map, color, x, y, -1, 1, cptHowManyThreeGroups, cptFourGroups)
	if cptFourGroups > 0 {
		bmap.SetInFourGroup(Map, x, y, true)
	} else {
		bmap.SetInFourGroup(Map, x, y, false)
	}
	if cptHowManyThreeGroups > 0 {
		bmap.SetInThreeGroup(Map, x, y, true)
	} else {
		bmap.SetInThreeGroup(Map, x, y, false)
	}
	return cptHowManyThreeGroups == 2
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

func CheckWinAlignment(dat *window.Drawer, Map *[361](uint64), game *GomokuGame, color bool) {
	game.Players[GetPlayerNb(game, color)].FiveAligned = game.Players[GetPlayerNb(game, color)].FiveAligned[:0]
	for x := range dat.BoardRes.Stones {
		for y := range dat.BoardRes.Stones[x] {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(Map, x, y) {
						if bmap.IsWhite(Map, x, y) == color {
							if CheckAlignement(Map, x, y, i, j, 3, 0, false) {
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

func CheckBreakableAlign(Map *[361]uint64, game *GomokuGame, color bool) bool {
	tot := 0
	for _, line := range game.Players[GetPlayerNb(game, color)].FiveAligned {
		cpt := 0
		for _, st := range line {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) &&
						bmap.GetNbT(Map, st.Infos.Ipos, st.Infos.Jpos, i+1, j+1) == 1 &&
						((bmap.GetNbO(Map, st.Infos.Ipos, st.Infos.Jpos, 1+(-1*i), 1+(-1*j)) >= 1) ||
							(bmap.GetNbO(Map, st.Infos.Ipos+j, st.Infos.Jpos+i, 1+i, 1+j) >= 1 &&
								!IsStoneAtPos(Map, st.Infos.Ipos+(-1*j), st.Infos.Jpos+(-1*i)))) {
						bmap.SetBreakable(Map, st.Infos.Ipos, st.Infos.Jpos, true)
						cpt = 1
					} else if IsStoneAtPos(Map, st.Infos.Ipos, st.Infos.Jpos) {
						bmap.SetBreakable(Map, st.Infos.Ipos, st.Infos.Jpos, false)
					}
				}
			}
		}
		tot += cpt
	}
	if tot < len(game.Players[GetPlayerNb(game, color)].FiveAligned) {
		return true
	}
	return false
}

func ApplyRules(game *GomokuGame, Map *[361](uint64), i, j int, color bool, rule1, rule2 bool) bool {
	bmap.SetColor(Map, i, j, color)
	if rule2 && ThreeBlockNear(Map, game, i, j, color) == true {
		return false
	}
	bmap.SetVisibility(Map, i, j, true)
	TakeTwoStones(Map, game, i, j, color)
	UpdateInfos(Map, game, game.Turn)
	return true
}

func GamePlay(pane *window.Drawer, game *GomokuGame, x, y, size int) {
	if game.End != 2 {
		st := IsStoneHere(pane, x, y, size)
		if st != nil && !bmap.IsVisible(&bmap.Map, st.Infos.Ipos, st.Infos.Jpos) {
			if !ApplyRules(game, &bmap.Map, st.Infos.Ipos, st.Infos.Jpos, game.Turn, pane.OptionsRes.Op1, pane.OptionsRes.Op2) {
				pane.BoardRes.BadX, pane.BoardRes.BadY = st.X, st.Y
				return
			}
			CheckWinAlignment(pane, &bmap.Map, game, game.Turn)
			if len(game.Players[GetPlayerNb(game, game.Turn)].FiveAligned) > 0 {
				if !pane.OptionsRes.Op1 || CheckBreakableAlign(&bmap.Map, game, game.Turn) {
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
