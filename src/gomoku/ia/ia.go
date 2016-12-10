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
		son.val = tree.val + Factor(son.carte, x, y, color, 1)
	} else {
		son.val = tree.val - Factor(son.carte, x, y, color, 0)
	}
	tree.sons = append(tree.sons, son)
	return false
}

func Factor(carte [363]uint64, x int, y int, color bool, player int) int {

	// AFFICHAGE DE DEBUG //
	s := "string"
	if player == 0 {
		s = "Blanc joue : "
	} else {
		s = "Noir joue : "
	}
	// FIN //

	if bmap.IsInFourGroup(&carte, x, y) == true {
		print(s)
		print("peut créer un groupe de quatre en  ")
		println(x, y)

		return 1000
	}
	if bmap.IsInThreeGroup(&carte, x, y) == true {
		print(s)

		print("peut créer un groupe de trois en  ")
		println(x, y)

		return 3
	}
	if bmap.IsInTwoGroup(&carte, x, y) == true {
		print(s)
		print("peut créer un groupe de quatre en  ")
		println(x, y)

		return 2
	}
	return 1
}

func sonSeek(father *NodeTree, t_color bool, deep int, rule1, rule2 bool, ckey chan<- *NodeTree) {
	var curr *NodeTree
	var result *NodeTree

	print("SUCKTHISSHIT")
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
			result.father.val = result.val
			result = result.father
		}
	}
	ckey <- result
}

func SeekWithRoutine(carte [363]uint64, t_color bool, deep int, rule1, rule2 bool) (int, int) {
	var racine NodeTree
	var result *NodeTree
	var tmp *NodeTree
	var y int

	cpt := 0
	color := t_color
	racine.father = nil
	racine.carte = carte
	racine.x = 17
	racine.y = 17
	result = nil
	ckey := make(chan *NodeTree)
	for x := 0; x < 19; x++ {
		for y = 0; y < 19; y++ {
			if !bmap.IsVisible(&racine.carte, x, y) {
				println(x, y)
			}
			if !bmap.IsVisible(&racine.carte, x, y) &&
				bmap.GetValStones(&racine.carte, x, y, bmap.MO)+bmap.GetValStones(&racine.carte, x, y, bmap.MT) != 0 {
				if !AddSon(&racine, x, y, t_color, color, rule1, rule2) {
					go sonSeek(racine.sons[len(racine.sons)-1], t_color, deep-1, rule1, rule2, ckey)
					cpt = cpt + 1
					break
				}
			}
		}
		if y != 19 {
			break
		}
	}
	println("OBAMA OUT")
	if cpt == 0 {
		return 9, 9
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
	result = nil
	check := true
	for curr != nil {
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
