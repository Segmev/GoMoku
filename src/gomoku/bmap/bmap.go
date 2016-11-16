package bmap

const Map_size = 19

var Map [Map_size * Map_size](uint64)

const (
	VISIBLE = 0
	COLOR   = 1

	// TODO: Update infos in arbitre, not done for now
	INTWOGROUP    = 2
	BREAKABLE     = 3
	INTHREEGROUP  = 4
	INDOUBLETHREE = 6

	ULT, UT, URT = 6 + (3 * 1), 6 + (3 * 2), 6 + (3 * 3)
	MLT, MT, MRT = 6 + (3 * 4), 6 + (3 * 5), 6 + (3 * 6)
	DLT, DT, DRT = 6 + (3 * 7), 6 + (3 * 8), 6 + (3 * 9)

	ULO, UO, URO = 6 + (3 * 10), 6 + (3 * 11), 6 + (3 * 12)
	MLO, MO, MRO = 6 + (3 * 13), 6 + (3 * 14), 6 + (3 * 15)
	DLO, DO, DRO = 6 + (3 * 16), 6 + (3 * 17), 6 + (3 * 18)
)

var TabTeam = [][]int{{ULT, UT, URT}, {MLT, MT, MRT}, {DLT, DT, DRT}}
var TabOppo = [][]int{{ULO, UO, URO}, {MLO, MO, MRO}, {DLO, DO, DRO}}

func getValAt(i, j int, info uint) int {
	if Map[(i*Map_size)+j]&(1<<info) == 0 {
		return 0
	}
	return 1
}

func GetValStones(i, j int, info uint) int {
	res := 0
	res = (res << 1) | getValAt(i, j, info)
	res = (res << 1) | getValAt(i, j, info+1)
	res = (res << 1) | getValAt(i, j, info+2)
	return res
}

func GetNbO(i, j int, posx, posy int) int {
	return GetValStones(i, j, uint(TabOppo[posx][posy]))
}

func GetNbT(i, j int, posx, posy int) int {
	return GetValStones(i, j, uint(TabTeam[posx][posy]))
}

func setAtPos(i, j int, infos uint, val uint64) {
	if val != Map[(i*Map_size)+j]&(1<<infos) {
		if val == 1 {
			Map[(i*Map_size)+j] |= (val << infos)
		} else {
			Map[(i*Map_size)+j] &^= (val << infos)
		}
	}
}

func SetNbOppoAt(i, j int, posx, posy int, nb uint64) {
	var info uint = uint(TabOppo[posx][posy])
	setAtPos(i, j, info+0, (nb>>2)&(1))
	setAtPos(i, j, info+1, (nb>>1)&(1))
	setAtPos(i, j, info+2, (nb>>0)&(1))
}

func SetNbTeamAt(i, j int, posx, posy int, nb uint64) {
	var info uint = uint(TabTeam[posx][posy])
	setAtPos(i, j, info+0, (nb>>2)&(1))
	setAtPos(i, j, info+1, (nb>>1)&(1))
	setAtPos(i, j, info+2, (nb>>0)&(1))
}

func IsInThreeGroup(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<INTHREEGROUP) != 0
}

func SetInThreeGroup(i, j int, val bool) {
	if val != IsBreakable(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << INTHREEGROUP)
		} else {
			Map[(i*Map_size)+j] &^= (1 << INTHREEGROUP)
		}
	}
}

func IsBreakable(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<BREAKABLE) != 0
}

func SetBreakable(i, j int, val bool) {
	if val != IsBreakable(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << BREAKABLE)
		} else {
			Map[(i*Map_size)+j] &^= (1 << BREAKABLE)
		}
	}
}

func IsInDoubleThree(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<INDOUBLETHREE) != 0
}

func SetInDoubleThree(i, j int, val bool) {
	if val != IsBreakable(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << INDOUBLETHREE)
		} else {
			Map[(i*Map_size)+j] &^= (1 << INDOUBLETHREE)
		}
	}
}

func IsInTwoGroup(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<INTWOGROUP) != 0
}

func SetInTwoGroup(i, j int, val bool) {
	if val != IsInTwoGroup(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << INTWOGROUP)
		} else {
			Map[(i*Map_size)+j] &^= (1 << INTWOGROUP)
		}
	}
}

func IsWhite(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<COLOR) != 0
}

func SetColor(i, j int, val bool) {
	if val != IsWhite(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << COLOR)
		} else {
			Map[(i*Map_size)+j] &^= (1 << COLOR)
		}
	}
}

func IsVisible(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<VISIBLE) != 0
}

func SetVisibility(i, j int, vis bool) {
	if vis != IsVisible(i, j) {
		if vis {
			Map[(i*Map_size)+j] |= (1 << VISIBLE)
		} else {
			Map[(i*Map_size)+j] &^= (1 << VISIBLE)
		}
	}
}

func ClearStone(i, j int) {
	Map[(i*Map_size)+j] = 0
}
