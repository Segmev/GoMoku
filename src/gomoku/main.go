package main

import (
	"fmt"
	"os"
	"gomoku/window"
	"github.com/gtalent/starfish/input"
	"github.com/gtalent/starfish/gfx"
)

func main() {
	if !gfx.OpenDisplay(800, 600, false) {
		return
	}

	gfx.SetDisplayTitle("Gomoku")

	var pane window.Drawer
	if pane.Init() {
		gfx.AddDrawer(&pane)
	}

	quit := func() {
		gfx.CloseDisplay()
		os.Exit(0)
	}
	input.AddQuitFunc(quit)

	input.AddKeyPressFunc(func(e input.KeyEvent) {
		fmt.Println("Key Press!")
		if e.Key == input.Key_Escape {
			quit()
		}
	})

	input.AddMousePressFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Press!")
		fmt.Println(e.X, e.Y)
		pane.St = true
	})

	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Release!")
		fmt.Println(e.X, e.Y)
		pane.St = false

	})
	gfx.Main()
}
