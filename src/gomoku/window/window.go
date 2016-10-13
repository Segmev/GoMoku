package window

import (
//	"os"
	"fmt"
	"github.com/gtalent/starfish/gfx"
)


type Stone struct {
	X int
	Y int
	Visible bool
	White bool
}

type Drawer struct {
	board		*gfx.Image
	text		*gfx.Text
	black_stone	*gfx.Image
	white_stone	*gfx.Image
	anim		*gfx.Animation
	St		bool
	Stones		[]*Stone
}

func (me *Drawer) AddStone(x, y int, white bool) bool {
	stone := new(Stone)
	stone.X, stone.Y, stone.White = x, y, white
	me.Stones = append(me.Stones, stone)
	fmt.Println("Added!")
	return true
}

func (me *Drawer) Init() bool {
	me.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	font := gfx.LoadFont("ressources/LiberationSans-Bold.ttf", 26)
	defer font.Free()
	if font != nil {
		font.SetRGB(100, 100, 255)
		me.text = font.Write("GoMoku")
	} else {
		fmt.Println("Could not load LiberationSans-Bold.ttf.")
		return false
	}
	me.Stones = []*Stone{}
	for j := 0; j <= 18; j++ {
		for i := 0; i <= 18; i++ {
			// c.DrawImage(me.white_stone,
			// 	gfx.DisplayHeight() / 88 + i * (gfx.DisplayHeight() / 19),
			// 	gfx.DisplayHeight() / 88 + j * (gfx.DisplayHeight() / 19))
			stone := new(Stone)
			stone.X = gfx.DisplayHeight() / 88 + i * (gfx.DisplayHeight() / 19)
			stone.Y = gfx.DisplayHeight() / 88 + j * (gfx.DisplayHeight() / 19)
			stone.White, stone.Visible = true, false
			me.Stones = append(me.Stones, stone)
		}
	}
	return true
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	//clear screen
	c.SetRGB(55, 55, 55)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())
	c.DrawText(me.text, gfx.DisplayWidth() * 3 / 4, 0)

	c.PushViewport(0, 0, gfx.DisplayWidth(), gfx.DisplayWidth())
	{
		c.DrawImage(me.board, 0, 0)

		for i := range(me.Stones) {
			if me.Stones[i].Visible {
				if me.Stones[i].White == true {
					c.DrawImage(me.white_stone, me.Stones[i].X, me.Stones[i].Y)
				} else {
					c.DrawImage(me.black_stone, me.Stones[i].X, me.Stones[i].Y)
				}
			}
		}
	}
	c.PopViewport()
}
