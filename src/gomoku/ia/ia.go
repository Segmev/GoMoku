package ia

import "gomoku/bmap"

type NodeTree struct {
	val    int
	carte  [19 * 19]uint64
	sons   []*NodeTree
	father *NodeTree
}

func AddSon(tree *NodeTree, x int, y int, carte[361]uint64) {
	var son *Nodetree = new NodeTree()

	son.father = tree;
}

func seek(carte [19 * 19]uint64, color bool) {
	var racine NodeTree

	racine.father = nil
	racine.val = 0
	racine.carte = carte
	for x := 0; x < 19; x++ {
		for y := 0; x < 19; y++ {
			if bmap.GetValStones(&carte, x, y, bmap.MO)+bmap.GetValStones(&carte, x, y, bmap.MT) != 0 {

			}
		}
	}
}
