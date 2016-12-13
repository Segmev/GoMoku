package main

import (
	"gomoku/arbitre"
	"gomoku/bmap"
	"gomoku/ia_monte_carlo"
	"gomoku/window"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/gtalent/starfish/gfx"
	"github.com/gtalent/starfish/input"
)

type Coor struct {
	X, Y, Val int
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func checkFiles() bool {
	if !exists("ressources/board.png") || !exists("ressources/bstone.png") ||
		!exists("ressources/woodback.jpg") || !exists("ressources/MotionControl-Bold.otf") ||
		!exists("ressources/wstone.png") || !exists("ressources/Red.png") ||
		!exists("ressources/Empty.png") || !exists("ressources/Cross.png") {
		return false
	}
	return true
}

func addInput(pane *window.Drawer, game *arbitre.GomokuGame) {
	quit := func() {
		gfx.CloseDisplay()
		os.Exit(0)
	}
	input.AddQuitFunc(quit)
	input.AddKeyPressFunc(func(e input.KeyEvent) {

		if e.Key == input.Key_Escape {
			quit()
		}
		if e.Key == input.Key_r {
			game.Restart(pane)
		}
	})
	input.AddMousePressFunc(func(e input.MouseEvent) {
		if e.Button == 1 {
			if pane.GameState == "gameOn" {
				GamePlay(pane, game, e.X, e.Y, gfx.DisplayWidth()/55)
			}
		}
	})
	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		if pane.GameState == "IA_Turn" {
			return
		}
		if pane.GameState == "menu" {
			if e.X >= 4*gfx.DisplayWidth()/14 && e.X <= 4*gfx.DisplayWidth()/14+8*gfx.DisplayWidth()/18 {
				if gfx.DisplayHeight()*4/10 <= e.Y && e.Y <= gfx.DisplayHeight()*4/10+gfx.DisplayHeight()/11 {
					pane.GameState = "gameOn"
					pane.GameType = "IA"
					if rand.Int()%2 == 0 {
						bmap.SetVisibility(&bmap.Map, 9, 9, true)
						game.Turn = true
					}
				} else if gfx.DisplayHeight()*5/10 <= e.Y && e.Y <= gfx.DisplayHeight()*5/10+gfx.DisplayHeight()/11 {
					pane.GameState = "gameOn"
					pane.GameType = "Human"
				} else if gfx.DisplayHeight()*6/10 <= e.Y && e.Y <= gfx.DisplayHeight()*6/10+gfx.DisplayHeight()/11 {
					pane.GameState = "options"
				}
			}
		} else if pane.GameState == "gameOn" || pane.GameState == "end" {
			if 10*gfx.DisplayWidth()/14 <= e.X && gfx.DisplayHeight()*10/11 <= e.Y {
				game.Restart(pane)
				pane.GameState = "menu"
			} else if 10*gfx.DisplayWidth()/14 <= e.X && gfx.DisplayHeight()*9/11 <= e.Y &&
				e.Y <= gfx.DisplayHeight()*10/11 {
				pane.GameState = "options"
			}
		} else if pane.GameState == "options" {
			if e.X >= gfx.DisplayHeight()*22/50 && e.X <= gfx.DisplayHeight()*22/50+50 {
				if e.Y >= gfx.DisplayHeight()/4 && e.Y <= gfx.DisplayHeight()/4+50 {
					pane.OptionsRes.Op1 = !pane.OptionsRes.Op1
				} else if e.Y >= gfx.DisplayHeight()/4+gfx.DisplayHeight()/9 &&
					e.Y <= gfx.DisplayHeight()/4+gfx.DisplayHeight()/9+50 {
					pane.OptionsRes.Op2 = !pane.OptionsRes.Op2
				}
			} else if e.X >= gfx.DisplayHeight()*27/50 && e.X <= gfx.DisplayHeight()*27/50+150 {
				if gfx.DisplayHeight()/4+(2*gfx.DisplayHeight()/9) <= e.Y &&
					e.Y <= gfx.DisplayHeight()/4+(2*gfx.DisplayHeight()/9)+50 {
					game.Restart(pane)
					pane.GameState = "gameOn"
				} else if gfx.DisplayHeight()/7 <= e.Y &&
					e.Y <= gfx.DisplayHeight()/7+50 {
					pane.OptionsRes.Lvl = (pane.OptionsRes.Lvl + 1) % 3
				} else if gfx.DisplayHeight()/4+(3*gfx.DisplayHeight()/9) <= e.Y &&
					e.Y <= gfx.DisplayHeight()/4+(3*gfx.DisplayHeight()/9)+50 {
					quit()
				} else if gfx.DisplayHeight()/4+(4*gfx.DisplayHeight()/9) <= e.Y &&
					e.Y <= gfx.DisplayHeight()/4+(4*gfx.DisplayHeight()/9)+50 {
					if pane.GameType != "" {
						pane.GameState = "gameOn"
					} else {
						pane.GameState = "menu"
					}
				}
			}
		}
	})

}

func launchWindow(h, w int) bool {
	if !gfx.OpenDisplay(h, w, false) {
		return false
	}
	gfx.SetDisplayTitle("GoMoku")

	var pane window.Drawer
	if pane.Init() {
		gfx.AddDrawer(&pane)
	}
	var game arbitre.GomokuGame
	// rand.Seed(time.Now().UTC().UnixNano())
	game.Turn = true
	pane.Turn = game.Turn
	game.End = 0
	pane.GameState = "menu"
	addInput(&pane, &game)

	return true
}

func PlayTurn(pane *window.Drawer, game *arbitre.GomokuGame, Map *[363]uint64, st *window.Stone) bool {
	if !arbitre.ApplyRules(&bmap.Map, st.Infos.Ipos, st.Infos.Jpos, game.Turn, pane.OptionsRes.Op1, pane.OptionsRes.Op2) {
		pane.BoardRes.BadX, pane.BoardRes.BadY = st.X, st.Y
		return false
	}
	var fl [][5]arbitre.Coor
	arbitre.CheckWinAl(&bmap.Map, game.Turn, &fl)
	if len(fl) > 0 {
		if !pane.OptionsRes.Op1 || arbitre.CheckBreakableAlign(&bmap.Map, fl, game.Turn) {
			game.End = 2
			pane.WinnerColor = game.Turn
		}
	}
	game.Turn = !game.Turn
	end, winColor := arbitre.HasTakenEnoughStones(&bmap.Map)
	if end {
		game.End = 2
		pane.WinnerColor = winColor
	}
	arbitre.IsDraw(pane, game)
	if game.End == 2 {
		pane.GameState = "end"
	}
	pane.Turn = game.Turn
	pane.BoardRes.Wscore = pane.Font.Write(strconv.Itoa(int(bmap.GetPlayerTakenStones(&bmap.Map, true))))
	pane.BoardRes.Bscore = pane.Font.Write(strconv.Itoa(int(bmap.GetPlayerTakenStones(&bmap.Map, false))))
	if st.Infos.Ipos < arbitre.XMin {
		arbitre.XMin = st.Infos.Ipos
	}
	if st.Infos.Ipos > arbitre.XMax {
		arbitre.XMax = st.Infos.Ipos
	}
	if st.Infos.Jpos < arbitre.YMin {
		arbitre.YMin = st.Infos.Jpos
	}
	if st.Infos.Jpos > arbitre.YMax {
		arbitre.YMax = st.Infos.Jpos
	}
	return true
}

func GamePlay(pane *window.Drawer, game *arbitre.GomokuGame, x, y, size int) {
	if game.End != 2 {
		st := arbitre.IsStoneHere(pane, x, y, size)
		if st != nil && !bmap.IsVisible(&bmap.Map, st.Infos.Ipos, st.Infos.Jpos) {
			if !PlayTurn(pane, game, &bmap.Map, st) {
				return
			}
		} else {
			return
		}
	}
	pane.BoardRes.BadX, pane.BoardRes.BadY = 0, 0
	if pane.GameType == "IA" && game.End != 2 {
		var iaStone window.Stone
		pane.GameState = "IA_Turn"

		//TON CODE HERE
		iaStone.Infos.Ipos, iaStone.Infos.Jpos = ia_monte_carlo.Play(&bmap.Map, pane.OptionsRes.Op1, pane.OptionsRes.Op2, 1000, bmap.Map)

		pane.GameState = "gameOn"
		PlayTurn(pane, game, &bmap.Map, &iaStone)
	}
}

func main() {
	ia_monte_carlo.Start(false)
	if checkFiles() && launchWindow(900, 640) {
		go http.ListenAndServe(":8080", http.DefaultServeMux)
		gfx.Main()
	} else {
		os.Stderr.WriteString("Couldn't launch the game, missing ressources or window can't be opened.\n")
	}
}
