package window

import (
//	"os"
	"fmt"
	"github.com/gtalent/starfish/gfx"
)

type Stone struct {
	x int
	y int
	white bool
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
	stone.x, stone.y, stone.white = x, y, white
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
		c.DrawImage(me.black_stone, gfx.DisplayHeight() / 88, gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 1 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 2 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 3 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 4 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 5 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 6 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 7 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 8 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 9 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 10 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 11 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 12 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 13 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 14 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 15 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 16 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 17 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		c.DrawImage(me.white_stone, gfx.DisplayHeight() / 88 + 18 * (gfx.DisplayHeight() / 19),
			gfx.DisplayHeight() / 88)
		// c.DrawImage(me.white_stone, 5 * gfx.DisplayHeight() / 50 , gfx.DisplayHeight() / 50)
		// c.DrawImage(me.white_stone, 9 * gfx.DisplayHeight() / 50 , gfx.DisplayHeight() / 50)
		// c.DrawImage(me.white_stone, 13 * gfx.DisplayHeight() / 50 , gfx.DisplayHeight() / 50)
		// c.DrawImage(me.white_stone, 20 * gfx.DisplayHeight() / 50 , gfx.DisplayHeight() / 50)

		for i := range(me.Stones) {
			if me.Stones[i].white == false {
				c.DrawImage(me.white_stone, me.Stones[i].x, me.Stones[i].y)
			} else {
				c.DrawImage(me.black_stone, me.Stones[i].x, me.Stones[i].y)
			}
		}
	}
	c.PopViewport()
}
