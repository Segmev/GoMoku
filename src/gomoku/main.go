package main

import (
	"fmt"
	"os"
	"gomoku/window"
	"gomoku/arbitre"
	"github.com/gtalent/starfish/input"
	"github.com/gtalent/starfish/gfx"
)

// func exists(path string) (bool) {
//     _, err := os.Stat(path)
//     if err == nil { return true }
//     if os.IsNotExist(err) { return false }
//     return true
// }


// func putPion(

func main() {
	if !gfx.OpenDisplay(800, 600, false) {
		return
	}

	gfx.SetDisplayTitle("GoMoku")

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
		fmt.Println("Key Press! ", e.Key)
		if e.Key == input.Key_Escape {
			quit()
		}
	})

	input.AddMousePressFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Press!")
		fmt.Println(e.X, e.Y, e.Button)
		if e.Button == 1 || e.Button == 3 {
			st := arbitre.IsStoneHere(&pane, e.X, e.Y, gfx.DisplayWidth() / 55)
			if st  != nil {
				if e.Button == 1 {
					st.White = true
				} else {
					st.White = false
				}
				arbitre.AppearStone(&pane, e.X, e.Y, gfx.DisplayWidth() / 55)
			}
		}
		// if e.Button == 1 {
		// 	pane.AddStone(e.X, e.Y, true)
		// } else {
		// 	pane.AddStone(e.X, e.Y, false)
		// }
	})

	// input.AddMouseReleaseFunc(func(e input.MouseEvent) {
	// 	fmt.Println("Mouse Release!")
	// 	fmt.Println(e.X, e.Y)
	// })
	gfx.Main()
}
