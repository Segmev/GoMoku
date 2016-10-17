package main

import (
	"os"
	"gomoku/window"
	"gomoku/arbitre"
	"github.com/gtalent/starfish/input"
	"github.com/gtalent/starfish/gfx"
)

func exists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return true
}

func checkFiles() bool {
	if !exists("ressources/board.png")         || !exists("ressources/bstone.png") ||
		!exists("ressources/woodback.jpg") || !exists("ressources/MotionControl-Bold.otf") ||
		!exists("ressources/wstone.png") {
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
		// fmt.Println("Mouse Press!")
		// fmt.Println(e.X, e.Y, e.Button)
		if e.Button == 1  {
			if pane.GameState == "gameOn" {
				arbitre.GamePlay(pane, game, e.X, e.Y, gfx.DisplayWidth() / 55)
			}
		}
	})
	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		if pane.GameState == "menu" {
	// 			c.FillRect(4 * gfx.DisplayWidth() / 14, gfx.DisplayHeight() * 4 / 10,
	// 	8 * gfx.DisplayWidth() / 18, gfx.DisplayHeight() / 10)
	// c.SetRGB(200,200,200)
	// c.FillRect(4 * gfx.DisplayWidth() / 14, gfx.DisplayHeight() * 5 / 10,
	// 	8 * gfx.DisplayWidth() / 18, gfx.DisplayHeight() / 10)

			if e.X >= 4 * gfx.DisplayWidth() / 14 && e.X <= 4 * gfx.DisplayWidth() / 14 +
				8 * gfx.DisplayWidth() / 18 {
				if gfx.DisplayHeight() * 4 / 10 <= e.Y && e.Y <= gfx.DisplayHeight() * 4 / 10 +
					gfx.DisplayHeight() / 11 {
					pane.GameState = "gameOn"
				} else if gfx.DisplayHeight() * 5 / 10 <= e.Y && e.Y <= gfx.DisplayHeight() *
					5 / 10 + gfx.DisplayHeight() / 11 {
					pane.GameState = "gameOn"
				}
			}
		} else if pane.GameState == "gameOn" {
			if 10 * gfx.DisplayWidth() / 14 <= e.X && gfx.DisplayHeight() * 10 / 11 <= e.Y {
				game.Restart(pane)
				pane.GameState = "menu"
			}
		}
	})
	// c.FillRect(10 * gfx.DisplayWidth() / 14, gfx.DisplayHeight() * 10 / 11,
	// 	8 * gfx.DisplayWidth() / 18, gfx.DisplayHeight() / 11)

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
	game.End = 0
	pane.GameState = "menu"
	addInput(&pane, &game)
	return true
}

func main() {
	if checkFiles() && launchWindow(900, 640) {
		gfx.Main()
	} else {
		os.Stderr.WriteString("Couldn't launch the game\n")
	}
}
