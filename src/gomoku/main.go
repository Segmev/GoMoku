package main

import (
	"fmt"
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
	if !exists("ressources/board.png") || !exists("ressources/bstone.png") ||
		!exists("ressources/woodback.jpg") ||
		!exists("ressources/wstone.png") || !exists("ressources/MotionControl-Bold.otf") {
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
		fmt.Println("Key Press! ", e.Key)
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
			arbitre.GamePlay(pane, game, e.X, e.Y, gfx.DisplayWidth() / 55)
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
	game.End = 0
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
