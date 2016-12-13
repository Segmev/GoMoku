package ia_monte_carlo

import (
	"gomoku/arbitre"
	"math/rand"
	"time"
)

var myColor bool
var hisColor bool
var _board [363]uint64
var resTab [361]int
var tmpTab [361]int
var tmpTab2 [361]int

var xMin, xMax, yMin, yMax int

func Start(color bool) {
	arbitre.XMin = 9
	arbitre.XMax = 9
	arbitre.YMin = 9
	arbitre.YMax = 9
	myColor = color
	hisColor = !myColor
}

func initResTab() {
	for a := 0; a < 361; a++ {
		resTab[a] = 0
	}
}

func myMemset() {
	for a := 0; a < 361; a++ {
		tmpTab[a] = 0
		tmpTab2[a] = 0
	}
}

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
	if !rule || arbitre.CheckBreakableAlign(&_board, res, color) {
		return true
	}
	return false
}

func refreshTab(value int) {
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

func MonteCarlo(board *[363]uint64, rule1, rule2 bool) (int, int) {
	var cpt, a, b, break_cpt, i int

	win := 0
	loose := 0
	iCheck := 0
	rand.Seed(time.Now().Unix())
	initResTab()
	findRange()
	for cpt = 0; cpt != 3000; cpt++ {
		myMemset()
		_board = *board
		for i = 0; i < 10; i++ {
			break_cpt = 0
			a = rand.Int() % (xMax - xMin)
			b = rand.Int() % (xMax - xMin)
			for !arbitre.ApplyRules(&_board, a+xMin, b+xMin, myColor, rule1, rule2) && break_cpt < 9 {
				a = rand.Int() % (xMax - xMin)
				b = rand.Int() % (xMax - xMin)
				break_cpt++
			}
			if break_cpt == 9 {
				break
			}
			tmpTab[(b+xMin)*19+(a+xMin)] = 1
			if CheckWin(rule1, myColor) {
				refreshTab(1)
				win += 1
				i = 9
				break
			}
			break_cpt = 0
			a = rand.Int() % (xMax - xMin)
			b = rand.Int() % (xMax - xMin)
			for !arbitre.ApplyRules(&_board, a+xMin, b+xMin, hisColor, rule1, rule2) && break_cpt < 9 {
				a = rand.Int() % (xMax - xMin)
				b = rand.Int() % (xMax - xMin)
				break_cpt++
			}
			if break_cpt == 9 {
				break
			}
			tmpTab2[(b+xMin)*19+(a+xMin)] = 1
			if CheckWin(rule1, hisColor) {
				refreshTab(-1)
				loose += 1
				i = 9
				break
			}
		}
		if i == 10 {
			iCheck++
		}
	}
	if iCheck == cpt {
		a = rand.Int() % (xMax - xMin)
		b = rand.Int() % (xMax - xMin)
		for !(arbitre.ApplyRules(board, a+xMin, b+xMin, myColor, rule1, rule2)) {
			a = rand.Int() % (xMax - xMin)
			b = rand.Int() % (xMax - xMin)
		}
		println("Return Par Défaut")
		return a, b
	} else {
		return findAndApply(board, rule1, rule2)
	}
}

func findAndApply(board *[363]uint64, rule1, rule2 bool) (int, int) {
	var a, b, save, saveA, saveB int

	for true {
		save = resTab[xMin+xMin*19]
		for a = xMin; a <= xMax; a++ {
			for b = xMin; b <= xMax; b++ {
				if resTab[a+b*19] > save {
					save = resTab[a+b*19]
					saveA = a
					saveB = b
				}
			}
		}
		if arbitre.ApplyRules(board, saveA, saveB, myColor, rule1, rule2) {
			break
		} else {
			resTab[saveA+saveB*19] = -9999
		}
	}
	println("Return Parfait")
	return saveA, saveB
}
