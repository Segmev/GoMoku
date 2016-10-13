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
		!exists("ressources/wstone.png") || !exists("ressources/LiberationSans-Bold.ttf") {
		return false
	}
	return true
}

func addInput(pane *window.Drawer) {
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
	})
	input.AddMousePressFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Press!")
		fmt.Println(e.X, e.Y, e.Button)
		if e.Button == 1 || e.Button == 3 {
			st := arbitre.IsStoneHere(pane, e.X, e.Y, gfx.DisplayWidth() / 55)
			if st  != nil {
				if e.Button == 1 {
					st.White = true
				} else {
					st.White = false
				}
				arbitre.AppearStone(pane, e.X, e.Y, gfx.DisplayWidth() / 55)
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
	addInput(&pane)
	return true
}

func main() {
	if checkFiles() && launchWindow(1400, 900) {
		gfx.Main()
	} else {
		os.Stderr.WriteString("Couldn't launch the game\n")
	}
	
}
