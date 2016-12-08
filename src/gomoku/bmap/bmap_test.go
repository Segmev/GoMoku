package bmap

import "testing"

func TestGetValStones(t *testing.T) {
	var MMap [Map_size*Map_size + Nb_Players](uint64)
	SetNbTeamAt(&MMap, 0, 0, 0, 0, 1)
	if GetNbT(&MMap, 0, 0, 0, 0) != 1 {
		t.Error("1 Got", GetNbT(&MMap, 0, 0, 0, 0))
	}
	SetNbTeamAt(&MMap, 0, 0, 0, 0, 2)
	if GetNbT(&MMap, 0, 0, 0, 0) != 2 {
		t.Error("2 Got", GetNbT(&MMap, 0, 0, 0, 0))
	}
	SetNbTeamAt(&MMap, 0, 0, 0, 0, 3)
	if GetNbT(&MMap, 0, 0, 0, 0) != 3 {
		t.Error("3 Got", GetNbT(&MMap, 0, 0, 0, 0))
	}
	SetNbTeamAt(&MMap, 0, 0, 0, 0, 4)
	if GetNbT(&MMap, 0, 0, 0, 0) != 4 {
		t.Error("4 Got", GetNbT(&MMap, 0, 0, 0, 0))
	}
	SetNbTeamAt(&MMap, 0, 0, 0, 0, 5)
	if GetNbT(&MMap, 0, 0, 0, 0) != 5 {
		t.Error("5 Got", GetNbT(&MMap, 0, 0, 0, 0))
	}
}
