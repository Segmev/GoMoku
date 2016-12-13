package arbitre

import (
	"gomoku/bmap"
	"gomoku/window"
	"strconv"
)

type Coor struct {
	X, Y int
}

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

var XMin, XMax, YMin, YMax int

func (game *GomokuGame) Restart(pane *window.Drawer) bool {
	for i := range pane.BoardRes.Stones {
		for j := range pane.BoardRes.Stones[i] {
			pane.BoardRes.Stones[i][j].Visible = false
			bmap.ClearStone(&bmap.Map, i, j)
		}
	}
	game.End = 0
	pane.EndRes.DrawEnd = false
	bmap.SetPlayerTakenStones(&bmap.Map, true, 0)
	bmap.SetPlayerTakenStones(&bmap.Map, false, 0)
	pane.BoardRes.Wscore = pane.Font.Write(strconv.Itoa(game.Players[0].Points))
	pane.BoardRes.Bscore = pane.Font.Write(strconv.Itoa(game.Players[1].Points))
	pane.GameState = "menu"
	pane.BoardRes.BadX, pane.BoardRes.BadY = 0, 0
	// rand.Seed(time.Now().UTC().UnixNano())
	// game.Turn = rand.Uint32()%2 == 0
	game.Turn = true
	pane.Turn = game.Turn
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

func IsStoneAtPos(Map *[363]uint64, i, j int) bool {
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

func CheckAlignement(Map *[363](uint64), x, y, i, j, lim, ite int, del bool) bool {
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

func TakeTwoStones(Map *[363](uint64), x, y int, color bool) bool {
	hap := false
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if CheckAlignement(Map, x, y, i, j, 2, 0, true) {
				bmap.SetPlayerTakenStones(Map, color, bmap.GetPlayerTakenStones(Map, color)+2)
				hap = true
			}
		}
	}
	return hap
}

func getInfosNbStonesDirection(Map *[363]uint64, x, y int, color bool, i, j int) int {
	cpt := 0
	for IsStoneAtPos(Map, x+i, y+j) &&
		bmap.IsWhite(Map, x+i, y+j) == color {
		cpt += 1
		i, j = augmentPos(i), augmentPos(j)
	}
	return cpt
}

func ResetTeamInfos(Map *[363]uint64, color bool) {
	for x := 0; x <= 18; x++ {
		for y := 0; y <= 18; y++ {
			if bmap.IsWhite(Map, x, y) == color {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						bmap.SetNbTeamAt(Map, x, y, i+1, j+1, 0)
						bmap.SetNbOppoAt(Map, x, y, i+1, j+1, 0)
					}
				}
				bmap.SetBreakable(Map, x, y, false)
			}
		}
	}
}

func UpdateInfos(Map *[363](uint64), color bool) {
	//ResetTeamInfos(dat, color)
	for x := 0; x <= 18; x++ {
		for y := 0; y <= 18; y++ {
			//UpdateThreeGroups(Map, x, y, color)
			totOpp, totTeam := 0, 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) {
						if bmap.IsWhite(Map, x, y) == color {
							bmap.SetNbTeamAt(Map, x, y, 1+j, 1+i, uint64(getInfosNbStonesDirection(Map, x, y,
								bmap.IsWhite(Map, x, y), i, j)))
							totTeam += bmap.GetNbT(Map, x, y, 1+j, 1+i)
						} else {
							bmap.SetNbOppoAt(Map, x, y, 1+j, 1+i, uint64(getInfosNbStonesDirection(Map, x, y,
								!bmap.IsWhite(Map, x, y), i, j)))
							totOpp += bmap.GetNbO(Map, x, y, 1+j, 1+i)
						}
					}
				}
			}
			bmap.SetNbTeamAt(Map, x, y, 1, 1, uint64(totTeam))
			bmap.SetNbOppoAt(Map, x, y, 1, 1, uint64(totOpp))
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

func HasTakenEnoughStones(Map *[363]uint64) (bool, bool) {
	if bmap.GetPlayerTakenStones(Map, true) >= 10 {
		return true, true
	} else if bmap.GetPlayerTakenStones(Map, false) >= 10 {
		return true, false
	}
	return false, false
}

func ThreeBlockNear(Map *[363]uint64, x, y int, color bool) bool {
	if UpdateThreeGroups(Map, x, y, color) {
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
						if UpdateThreeGroups(Map, x+a, y+b, color) {
							bmap.ResetStone(Map, x, y)
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

func updateThreeGroupLoop(Map *[363]uint64, color bool, x, y, dirI, dirJ, cptHowManyThreeGroups, cptFourGroups int) (int, int) {
	cpt := 0
	if !(dirI == 0 && dirJ == 0) {
		i, j := dirI, dirJ
		end := 2
		if (i <= 0 && j <= 0) || (i == -1 && j == 1) {
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

func UpdateThreeGroups(Map *[363]uint64, x, y int, color bool) bool {
	cptHowManyThreeGroups, cptFourGroups := 0, 0
	for dirI := 0; dirI <= 1; dirI++ {
		for dirJ := 0; dirJ <= 1; dirJ++ {
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
	return cptHowManyThreeGroups >= 2
}

func IsDraw(pane *window.Drawer, game *GomokuGame) {
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

func fillAlignedArray(FiveAligned *[][5]Coor, x, y, j, i int) {
	if x+4*i >= 0 && x+4*i <= 18 && y+4*j >= 0 && y+4*j <= 18 {
		StonesTab := [5]Coor{
			Coor{x, y},
			Coor{x + i, y + j},
			Coor{x + i + i, y + j + j},
			Coor{x + i + i + i, y + j + j + j},
			Coor{x + i + i + i + i, y + j + j + j + j}}
		*FiveAligned = append(*FiveAligned, StonesTab)
	}
}

func CheckWinAl(Map *[363](uint64), color bool, FiveAligned *[][5]Coor) {
	for x := 0; x <= 18; x++ {
		for y := 0; y <= 18; y++ {
			if bmap.IsVisible(Map, x, y) && bmap.IsWhite(Map, x, y) == color && bmap.GetValStones(Map, x, y, bmap.MT) >= 4 {
				if bmap.GetNbT(Map, x, y, 0, 2) >= 4 {
					fillAlignedArray(FiveAligned, x, y, 1, -1)
				}
				for j := 1; j <= 2; j++ {
					for i := 1; i <= 2; i++ {
						if !(i == 1 && j == 1) && bmap.GetNbT(Map, x, y, i, j) >= 4 {
							fillAlignedArray(FiveAligned, x, y, i-1, j-1)
						}
					}
				}
			}
		}
	}
}

func Break_cases(Map *[363]uint64, Ipos, Jpos, i, j int) bool {
	return bmap.GetNbT(Map, Ipos, Jpos, i+1, j+1) == 1 &&
		((!IsStoneAtPos(Map, Ipos-i-i, Jpos-j-j) && bmap.GetNbO(Map, Ipos, Jpos, 1+(-1*i), 1+(-1*j)) >= 1) ||
			(bmap.GetNbO(Map, Ipos+j, Jpos+i, 1+i, 1+j) >= 1 && !IsStoneAtPos(Map, Ipos+(-1*j), Jpos+(-1*i))))
}

// func Break_cases(Map *[363]uint64, st *window.Stone, i, j int) bool {
// 	return bmap.GetNbT(Map, st.Infos.Ipos, st.Infos.Jpos, i+1, j+1) == 1 &&
// 		((!IsStoneAtPos(Map, st.Infos.Ipos-i-i, st.Infos.Jpos-j-j) && bmap.GetNbO(Map, st.Infos.Ipos, st.Infos.Jpos, 1+(-1*i), 1+(-1*j)) >= 1) ||
// 			(bmap.GetNbO(Map, st.Infos.Ipos+j, st.Infos.Jpos+i, 1+i, 1+j) >= 1 && !IsStoneAtPos(Map, st.Infos.Ipos+(-1*j), st.Infos.Jpos+(-1*i))))
// }

func CheckBreakableAlign(Map *[363]uint64, fl [][5]Coor, color bool) bool {
	tot := 0
	for _, line := range fl {
		cpt := 0
		for _, st := range line {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && IsStoneAtPos(Map, st.X, st.Y) &&
						Break_cases(Map, st.X, st.Y, i, j) {
						bmap.SetBreakable(Map, st.X, st.Y, true)
						cpt = 1
					} else if IsStoneAtPos(Map, st.X, st.Y) {
						bmap.SetBreakable(Map, st.X, st.Y, false)
					}
				}
			}
		}
		tot += cpt
	}
	if tot < len(fl) {
		return true
	}
	return false
}

func ApplyRules(Map *[363](uint64), i, j int, color bool, rule1, rule2 bool) bool {
	bmap.SetColor(Map, i, j, color)
	if rule2 && ThreeBlockNear(Map, i, j, color) == true {
		return false
	}
	bmap.SetVisibility(Map, i, j, true)
	TakeTwoStones(Map, i, j, color)
	UpdateInfos(Map, color)
	return true
}
