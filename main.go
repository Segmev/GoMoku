package main

import(
	"os"
	"github.com/gtalent/starfish/gfx"
	"github.com/gtalent/starfish/input"
	"fmt"
	"runtime"
	// "runtime/pprof"

)

func main() {
	if !gfx.OpenDisplay(800, 600, false) {
		return
	}
	
	gfx.SetDisplayTitle("Gomoku")

	quit := func () {
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
	})
	
	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Release!")
		fmt.Println(e.X, e.Y)
		
	})
	
	gfx.Main()
}
