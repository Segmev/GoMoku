package main

import (
	"fmt"
	"github.com/gtalent/starfish/gfx"
	"github.com/gtalent/starfish/input"
	"os"
	// "runtime"
	// "runtime/pprof"
)

type Drawer struct {
	box  *gfx.Image
	text *gfx.Text
	anim *gfx.Animation
	st   bool
}

func (me *Drawer) init() bool {
	me.anim = gfx.NewAnimation(1000)
	me.box = gfx.LoadImageSize("box.png", 100, 100)
	font := gfx.LoadFont("LiberationSans-Bold.ttf", 32)
	if font != nil {
		font.SetRGB(0, 0, 255)
		me.text = font.Write("starfish 0.12!")
		font.Free()
	} else {
		fmt.Println("Could not load LiberationSans-Bold.ttf.")
		return false
	}
	if me.box == nil {
		fmt.Println("Could not load box.png.")
		return false
	}
	me.anim.LoadImageSize("box.png", 70, 70)
	me.anim.LoadImageSize("dots.png", 70, 70)
	return true
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	//clear screen
	c.SetRGB(0, 0, 0)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())

	if !me.st {
		c.SetRGBA(0, 0, 255, 255)
		c.FillRect(42, 42, 100, 100)
	} else {
		c.SetRGBA(255, 0, 0, 255)
		c.FillRect(42, 42, 100, 100)
	}

	//draw box if it's not nil
	if me.box != nil {
		c.DrawImage(me.box, 200, 200)
		c.SetRGBA(0, 0, 0, 100)
		c.FillRect(200, 200, 100, 100)
	}
	c.DrawText(me.text, 400, 400)

	//Note: viewports may be nested
	c.PushViewport(100, 42, 5000, 5000)
	{
		//draw a green rect in a viewport
		c.SetRGBA(0, 255, 100, 127)
		c.FillRect(42, 42, 100, 100)
		c.SetRGB(0, 0, 0)
		c.FillRect(350, 200, 70, 70)
		c.DrawAnimation(me.anim, 350, 200)
	}
	c.PopViewport()
}

func main() {
	if !gfx.OpenDisplay(800, 600, false) {
		return
	}

	gfx.SetDisplayTitle("Gomoku")

	var pane Drawer
	if pane.init() {
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
		pane.st = true
	})

	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Release!")
		fmt.Println(e.X, e.Y)
		pane.st = false

	})

	gfx.Main()
}
