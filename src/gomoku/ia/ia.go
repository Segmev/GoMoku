package ia

import (
	"gomoku/arbitre"
	"gomoku/bmap"
)

type NodeTree struct {
	val    int
	x      int
	y      int
	carte  [363]uint64
	sons   []*NodeTree
	father *NodeTree
}

func AddSon(tree *NodeTree, x int, y int, true_color, color, rule1, rule2 bool) bool {
	son := new(NodeTree)

	son.carte = tree.carte
	if !arbitre.ApplyRules(&son.carte, x, y, color, rule1, rule2) {
		return true
	}
	son.x = x
	son.y = y
	son.father = tree
	if color == true_color {
		son.val = tree.val + 2
	} else {
		son.val = tree.val - 1
	}
	tree.sons = append(tree.sons, son)
	return false
}

func Seek(carte [363]uint64, t_color bool, deep int, rule1, rule2 bool) (int, int) {
	var racine NodeTree
	var curr *NodeTree
	var result *NodeTree

	color := t_color
	cpt := 0
	best_cpt := deep
	racine.father = nil
	racine.val = 0
	racine.carte = carte
	racine.x = 17
	racine.y = 17
	curr = &racine
	x := 0
	y := 0
	result = nil
	check := true
	for curr != nil {
		check = true
		if !bmap.IsVisible(&curr.carte, x, y) &&
			bmap.GetValStones(&curr.carte, x, y, bmap.MO)+bmap.GetValStones(&curr.carte, x, y, bmap.MT) != 0 &&
			cpt < deep &&
			(curr.val < 1000 && curr.val > -1000) {
			check = AddSon(curr, x, y, t_color, color, rule1, rule2)
			if !check {
				curr = curr.sons[len(curr.sons)-1]
				cpt = cpt + 1
				color = !color
			}
		}
		if (result == nil) || (result.val < curr.val && cpt < best_cpt) {
			if curr != &racine {
				result = curr
			}
		}
		if check {
			if x == 18 {
				if y == 18 {
					if curr.x == 18 {
						x = curr.x
					} else {
						x = curr.x + 1
					}
					if curr.y == 18 {
						y = curr.y
					} else {
						y = curr.y + 1
					}
					curr = curr.father
					cpt = cpt - 1
					color = !color
				} else {
					x = 0
					y = y + 1
				}
			} else {
				x = x + 1
			}
		}
	}
	if result == nil {
		return 9, 9
	}
	for result.father.father != nil {
		result = result.father
	}
	return result.x, result.y
}
