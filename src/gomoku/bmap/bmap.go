package bmap

const Map_size = 19

var Map [Map_size * Map_size]uint8

const (
	VISIBLE      = 0
	COLOR        = 1
	SOLO         = 2
	INTWOGROUP   = 3
	INTHREEGROUP = 4
	BREAKABLE    = 5
)

func IsBreakable(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<BREAKABLE) != 0
}

func SetBreakable(i, j int, val bool) {
	if val != IsBreakable(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << BREAKABLE)
		} else {
			Map[(i*Map_size)+j] ^= (1 << BREAKABLE)
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
			Map[(i*Map_size)+j] ^= (1 << INTWOGROUP)
		}
	}
}

func IsInThreeGroup(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<INTHREEGROUP) != 0
}

func SetInThreeGroup(i, j int, val bool) {
	if val != IsInThreeGroup(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << INTHREEGROUP)
		} else {
			Map[(i*Map_size)+j] ^= (1 << INTHREEGROUP)
		}
	}
}

func IsSolo(i, j int) bool {
	return Map[(i*Map_size)+j]&(1<<SOLO) != 0
}

func SetSolo(i, j int, val bool) {
	if val != IsSolo(i, j) {
		if val {
			Map[(i*Map_size)+j] |= (1 << SOLO)
		} else {
			Map[(i*Map_size)+j] ^= (1 << SOLO)
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
			Map[(i*Map_size)+j] ^= (1 << COLOR)
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
			Map[(i*Map_size)+j] ^= (1 << VISIBLE)
		}
	}
}
