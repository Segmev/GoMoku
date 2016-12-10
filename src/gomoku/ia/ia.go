package ia

import (
	"fmt"
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
		son.val = tree.val + 1
	} else {
		son.val = tree.val - 1
	}
	tree.sons = append(tree.sons, son)
	return false
}

func sonSeek(father *NodeTree, t_color bool, deep int, rule1, rule2 bool, ckey chan<- *NodeTree) {
	var curr *NodeTree
	var result *NodeTree

	fmt.Print("Halp")
	color := !t_color
	cpt := 0
	best_cpt := deep
	curr = father
	x := 0
	y := 0
	result = nil
	check := true
	for curr != father.father {
		check = true
		if !bmap.IsVisible(&curr.carte, x, y) &&
			bmap.GetValStones(&curr.carte, x, y, bmap.MO)+bmap.GetValStones(&curr.carte, x, y, bmap.MT) != 0 &&
			cpt < deep &&
			curr.val < 1000 && curr.val > -1000 {
			check = AddSon(curr, x, y, t_color, color, rule1, rule2)
			if !check {
				curr = curr.sons[len(curr.sons)-1]
				cpt = cpt + 1
				color = !color
			}
		}
		if (result == nil) || (result.val < curr.val && cpt < best_cpt) {
			if curr != father {
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
	if result != nil {
		for result.father.father != nil {
			result = result.father
		}
	}
	ckey <- result
}

func SeekWithRoutine(carte [363]uint64, t_color bool, deep int, rule1, rule2 bool) (int, int) {
	var racine NodeTree
	var curr *NodeTree
	var result *NodeTree
	var tmp *NodeTree

	cpt := 0
	color := t_color
	racine.father = nil
	racine.carte = carte
	result = nil
	ckey := make(chan *NodeTree)
	for x := 0; x < 19; x++ {
		for y := 0; y < 19; y++ {
			if !bmap.IsVisible(&racine.carte, x, y) &&
				bmap.GetValStones(&racine.carte, x, y, bmap.MO)+bmap.GetValStones(&racine.carte, x, y, bmap.MT) != 0 {
				if AddSon(&racine, x, y, t_color, color, rule1, rule2) {
					go sonSeek(racine.sons[len(curr.sons)-1], t_color, deep-1, rule1, rule2, ckey)
					cpt = cpt + 1
					break
				}
			}
			if y != 19 {
				break
			}
		}
	}
	for cpt >= 0 {
		tmp = <-ckey
		if (tmp != nil) && (result == nil || result.val > tmp.val) {
			result = tmp
		}
	}
	if result == nil {
		return 9, 9
	}
	return result.x, result.y
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
	for curr != nil {
		result = nil
		check := true
		check = true
		if !bmap.IsVisible(&curr.carte, x, y) &&
			bmap.GetValStones(&curr.carte, x, y, bmap.MO)+bmap.GetValStones(&curr.carte, x, y, bmap.MT) != 0 &&
			cpt < deep &&
			curr.val < 1000 && curr.val > -1000 {
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
