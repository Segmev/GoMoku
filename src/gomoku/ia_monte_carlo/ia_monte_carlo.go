package ia_monte_carlo

import (
	"crypto/rand"
	"fmt"
	"gomoku/arbitre"
	"gomoku/bmap"
	"math/big"
	"sync"
	"time"
)

var myColor bool
var hisColor bool
var _board [363]uint64
var resTab [361]int
var mute sync.Mutex

var xMin, xMax, yMin, yMax int64

func Start(color bool) {
	arbitre.XMin = 9
	arbitre.XMax = 9
	arbitre.YMin = 9
	arbitre.YMax = 9
	myColor = color
	hisColor = !myColor
}

func ApplyRules(Map *[363](uint64), i, j int, color bool, rule1, rule2, updateSt bool) bool {
	if bmap.IsVisible(Map, i, j) {
		return false
	}
	if updateSt {
		arbitre.UpdateStone(Map, i, j, color)
	}
	return arbitre.ApplyRules(Map, i, j, color, rule1, rule2, false)
}

func initResTab() {
	for a := 0; a < 361; a++ {
		resTab[a] = 0
	}
}

// func myMemset() {
// 	for a := 0; a < 361; a++ {
// 		tmpTab[a] = 0
// 		tmpTab2[a] = 0
// 	}
// }

func ResBoard(board *[363]uint64) {
	for a := 0; a < 363; a++ {
		_board[a] = board[a]
	}
}

func setBoard(board *[363]uint64) {
	_board = *board
}

func CheckWin(rule bool, color bool) bool {
	if _board[361] >= 10 {
		return true
	}
	if _board[362] >= 10 {
		return true
	}

	var res [][5]arbitre.Coor

	arbitre.CheckWinAl(&_board, color, &res)
	if (!rule && len(res) != 0) || (rule && arbitre.CheckBreakableAlign(&_board, res, color)) {
		return true
	}
	return false
}

func refreshTab(value int, tmpTab *[361]int, tmpTab2 *[361]int) {
	mute.Lock()
	if value == 1 {
		for y := 0; y < 361; y++ {
			resTab[y] -= tmpTab2[y]
			resTab[y] += tmpTab[y]
		}
	} else {
		for y := 0; y < 361; y++ {
			resTab[y] += tmpTab2[y]
			resTab[y] -= tmpTab[y]
		}
	}
	mute.Unlock()
}

func findRange() {
	if arbitre.XMin > 1 {
		xMin = arbitre.XMin - 2
	} else if arbitre.XMin > 0 {
		xMin = arbitre.XMin - 1
	}
	if arbitre.YMin > 1 {
		yMin = arbitre.YMin - 2
	} else if arbitre.YMin > 0 {
		yMin = arbitre.YMin - 1
	}
	if arbitre.XMax < 17 {
		xMax = arbitre.XMax + 2
	} else if arbitre.XMax < 18 {
		xMax = arbitre.XMax + 1
	}
	if arbitre.YMax < 17 {
		yMax = arbitre.YMax + 2
	} else if arbitre.YMax < 18 {
		yMax = arbitre.YMax + 1
	}
}

func Play(board *[363]uint64, rule1, rule2 bool, test_nb int64, tmpboard [363]uint64) (int64, int64) {
	initResTab()
	starttime := time.Now()
	ch := make(chan bool, 6)
	findRange()
	go MonteCarlo(board, rule1, rule2, test_nb/4, ch)
	go MonteCarlo(board, rule1, rule2, test_nb/4, ch)
	go MonteCarlo(board, rule1, rule2, test_nb/4, ch)
	go MonteCarlo(board, rule1, rule2, test_nb/4, ch)
	<-ch
	x, y := findAndApply(&tmpboard, rule1, rule2)
	fmt.Println("out:", time.Now().Sub(starttime).Seconds())
	return x, y
}

func MonteCarlo(board *[363]uint64, rule1, rule2 bool, test_nb int64, ch chan bool) {
	var cpt, a, b, break_cpt, i int64
	var large *big.Int
	var tmpTab [361]int
	var tmpTab2 [361]int
	var empty [361]int

	yLim := big.NewInt(yMax - yMin)
	xLim := big.NewInt(xMax - xMin)
	reader := rand.Reader
	iCheck := 0
	for cpt = 0; cpt != test_nb; cpt++ {
		tmpTab = empty
		tmpTab2 = empty
		_board = *board
		for i = 0; i < 10; i++ {
			break_cpt = 0
			large, _ = rand.Int(reader, xLim)
			a = large.Int64()
			large, _ = rand.Int(reader, yLim)
			b = large.Int64()
			for !ApplyRules(&_board, int(a+xMin), int(b+yMin), myColor, rule1, rule2, break_cpt > 6) && break_cpt < 9 {
				large, _ = rand.Int(reader, xLim)
				a = large.Int64()
				large, _ = rand.Int(reader, yLim)
				b = large.Int64()
				break_cpt++
			}
			if break_cpt == 9 {
				break
			}
			tmpTab[(b+yMin)*19+(a+xMin)] = 1
			if CheckWin(rule1, myColor) {
				refreshTab(1, &tmpTab, &tmpTab2)
				i = 9
				break
			}
			break_cpt = 0
			large, _ = rand.Int(reader, xLim)
			a = large.Int64()
			large, _ = rand.Int(reader, yLim)
			b = large.Int64()
			for !ApplyRules(&_board, int(a+xMin), int(b+yMin), hisColor, rule1, rule2, break_cpt > 6) && break_cpt < 9 {
				large, _ = rand.Int(reader, xLim)
				a = large.Int64()
				large, _ = rand.Int(reader, yLim)
				b = large.Int64()
				break_cpt++
			}
			if break_cpt == 9 {
				break
			}
			tmpTab2[(b+yMin)*19+(a+xMin)] = 1
			if CheckWin(rule1, hisColor) {
				refreshTab(-1, &tmpTab, &tmpTab2)
				i = 9
				break
			}
		}
		if i == 10 {
			iCheck++
		}
	}
	ch <- true
}

func findAndApply(board *[363]uint64, rule1, rule2 bool) (int64, int64) {
	var a, b int64
	var save int

	saveA := xMin
	saveB := yMin
	for true {
		save = resTab[xMin+yMin*19]
		for a = xMin; a <= xMax; a++ {
			for b = yMin; b <= yMax; b++ {
				if resTab[a+b*19] > save {
					save = resTab[a+b*19]
					saveA = a
					saveB = b
				}
			}
		}
		if ApplyRules(board, int(saveA), int(saveB), myColor, rule1, rule2, true) {
			break
		} else {
			resTab[saveA+saveB*19] = -9999
		}
	}
	return saveA, saveB
}
