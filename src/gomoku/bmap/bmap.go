// Bmap is a set of functions for getting and setting game information
package bmap

const Map_size = 19
const Nb_Players = 2

var Map [Map_size*Map_size + Nb_Players](uint64)

const (
	VISIBLE = 0
	COLOR   = 1

	INTWOGROUP    = 2
	BREAKABLE     = 3
	INTHREEGROUP  = 4
	INDOUBLETHREE = 5
	ISCHECKED     = 6

	ULT, UT, URT = 6 + (3 * 1), 6 + (3 * 2), 6 + (3 * 3)
	MLT, MT, MRT = 6 + (3 * 4), 6 + (3 * 5), 6 + (3 * 6)
	DLT, DT, DRT = 6 + (3 * 7), 6 + (3 * 8), 6 + (3 * 9)

	ULO, UO, URO = 6 + (3 * 10), 6 + (3 * 11), 6 + (3 * 12)
	MLO, MO, MRO = 6 + (3 * 13), 6 + (3 * 14), 6 + (3 * 15)
	DLO, DO, DRO = 6 + (3 * 16), 6 + (3 * 17), 6 + (3 * 18)
)

var TabTeam = [][]int{{ULT, UT, URT}, {MLT, MT, MRT}, {DLT, DT, DRT}}
var TabOppo = [][]int{{ULO, UO, URO}, {MLO, MO, MRO}, {DLO, DO, DRO}}

func ResetStone(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) {
	MMap[(i*Map_size)+j] = 0
}

func getValAt(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, info uint) int {
	if MMap[(i*Map_size)+j]&(1<<info) == 0 {
		return 0
	}
	return 1
}

func GetValStones(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, info uint) int {
	res := 0
	res = (res << 1) | getValAt(MMap, i, j, info)
	res = (res << 1) | getValAt(MMap, i, j, info+1)
	res = (res << 1) | getValAt(MMap, i, j, info+2)
	return res
}

func GetNbO(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, posx, posy int) int {
	return GetValStones(MMap, i, j, uint(TabOppo[posx][posy]))
}

func GetNbT(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, posx, posy int) int {
	return GetValStones(MMap, i, j, uint(TabTeam[posx][posy]))
}

func setAtPos(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, infos uint, val uint64) {
	if val != MMap[(i*Map_size)+j]&(1<<infos) {
		if val == 1 {
			MMap[(i*Map_size)+j] |= (1 << infos)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << infos)
		}
	}
}

func SetNbOppoAt(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, posx, posy int, nb uint64) {
	var info uint = uint(TabOppo[posx][posy])
	if nb > 7 {
		nb = 7
	}
	setAtPos(MMap, i, j, info+2, (nb>>0)&(1))
	setAtPos(MMap, i, j, info+1, (nb>>1)&(1))
	setAtPos(MMap, i, j, info+0, (nb>>2)&(1))
}

func SetNbTeamAt(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, posx, posy int, nb uint64) {
	var info uint = uint(TabTeam[posx][posy])
	if nb > 7 {
		nb = 7
	}
	setAtPos(MMap, i, j, info+2, (nb>>0)&(1))
	setAtPos(MMap, i, j, info+1, (nb>>1)&(1))
	setAtPos(MMap, i, j, info+0, (nb>>2)&(1))
}

func IsChecked(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<ISCHECKED) != 0
}

func ResetCheck(MMap *[Map_size*Map_size + Nb_Players](uint64)) {
	for i := 0; i <= Map_size*Map_size; i++ {
		if MMap[i]&(1<<ISCHECKED) != 0 {
			MMap[i] &^= (1 << ISCHECKED)
		}
	}
}

func SetChecked(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, val bool) {
	if val != IsChecked(MMap, i, j) {
		if val {
			MMap[(i*Map_size)+j] |= (1 << ISCHECKED)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << ISCHECKED)
		}
	}
}

func IsInThreeGroup(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<INTHREEGROUP) != 0
}

func SetInThreeGroup(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, val bool) {
	if val != IsInThreeGroup(MMap, i, j) {
		if val {
			MMap[(i*Map_size)+j] |= (1 << INTHREEGROUP)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << INTHREEGROUP)
		}
	}
}

func IsBreakable(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<BREAKABLE) != 0
}

func SetBreakable(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, val bool) {
	if val != IsBreakable(MMap, i, j) {
		if val {
			MMap[(i*Map_size)+j] |= (1 << BREAKABLE)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << BREAKABLE)
		}
	}
}

func IsInDoubleThree(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<INDOUBLETHREE) != 0
}

func SetInDoubleThree(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, val bool) {
	if val != IsInDoubleThree(MMap, i, j) {
		if val {
			MMap[(i*Map_size)+j] |= (1 << INDOUBLETHREE)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << INDOUBLETHREE)
		}
	}
}

func IsInTwoGroup(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<INTWOGROUP) != 0
}

func SetInTwoGroup(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, val bool) {
	if val != IsInTwoGroup(MMap, i, j) {
		if val {
			MMap[(i*Map_size)+j] |= (1 << INTWOGROUP)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << INTWOGROUP)
		}
	}
}

func IsWhite(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<COLOR) != 0
}

func SetColor(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, val bool) {
	if val != IsWhite(MMap, i, j) {
		if val {
			MMap[(i*Map_size)+j] |= (1 << COLOR)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << COLOR)
		}
	}
}

func IsVisible(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) bool {
	return MMap[(i*Map_size)+j]&(1<<VISIBLE) != 0
}

func SetVisibility(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int, vis bool) {
	if vis != IsVisible(MMap, i, j) {
		if vis {
			MMap[(i*Map_size)+j] |= (1 << VISIBLE)
		} else {
			MMap[(i*Map_size)+j] &^= (1 << VISIBLE)
		}
	}
}

func ClearStone(MMap *[Map_size*Map_size + Nb_Players](uint64), i, j int) {
	MMap[(i*Map_size)+j] = 0
}

func AddToPlayerTakenStones(MMap *[Map_size*Map_size + Nb_Players](uint64), player_color bool, nb uint64) {
	player := Map_size*Map_size + func(color bool) int {
		if color == false {
			return 0
		} else {
			return 1
		}
	}(player_color)
	MMap[player] += nb
}

func SetPlayerTakenStones(MMap *[Map_size*Map_size + Nb_Players](uint64), player_color bool, nb uint64) {
	player := Map_size*Map_size + func(color bool) int {
		if color == false {
			return 0
		} else {
			return 1
		}
	}(player_color)
	MMap[player] = nb
}

func GetPlayerTakenStones(MMap *[Map_size*Map_size + Nb_Players](uint64), player_color bool) uint64 {
	player := Map_size*Map_size + func(color bool) int {
		if color == false {
			return 0
		} else {
			return 1
		}
	}(player_color)
	return MMap[player]
}
