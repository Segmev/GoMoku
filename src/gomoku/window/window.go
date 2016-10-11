package window

import (
//	"os"
	"fmt"
	"github.com/gtalent/starfish/gfx"
)

type Drawer struct {
	box  *gfx.Image
	text *gfx.Text
	anim *gfx.Animation
	St   bool
}

// func exists(path string) (bool) {
//     _, err := os.Stat(path)
//     if err == nil { return true }
//     if os.IsNotExist(err) { return false }
//     return true
// }

func (me *Drawer) Init() bool {
	// if !exists("box.png") || !exists("dots.png") {
	// 	return false
	// }
	me.anim = gfx.NewAnimation(1000)
	me.box = gfx.LoadImageSize("ressources/box.png", 100, 100)
	font := gfx.LoadFont("ressources/LiberationSans-Bold.ttf", 32)
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
	me.anim.LoadImageSize("ressources/box.png", 70, 70)
	me.anim.LoadImageSize("ressources/dots.png", 70, 70)
	return true
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	//clear screen
	c.SetRGB(0, 0, 0)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())

	if !me.St {
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
