package window

import (
//	"os"
//	"fmt"
	"github.com/gtalent/starfish/gfx"
	"strconv"
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
	wscore		*gfx.Text
	bscore		*gfx.Text
	black_stone	*gfx.Image
	white_stone	*gfx.Image
	anim		*gfx.Animation
	St		bool
	Stones		[]*Stone
	font		*gfx.Font
}

func (me *Drawer) Init() bool {
	me.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.font = gfx.LoadFont("ressources/LiberationSans-Bold.ttf", 26)
	me.font.SetRGB(100, 100, 255)
	me.text = me.font.Write("GoMoku")
	me.Stones = []*Stone{}
	for j := 0; j <= 18; j++ {
		for i := 0; i <= 18; i++ {
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
	me.font.SetRGB(100, 100, 255)
	c.DrawText(me.text, gfx.DisplayWidth() * 3 / 4, 0)
	me.font.SetRGB(255, 255, 0)
	me.wscore = me.font.Write(strconv.Itoa(3))
	me.bscore = me.font.Write(strconv.Itoa(2))
	c.DrawText(me.wscore, gfx.DisplayWidth() * 3 / 4, 20)
	c.DrawText(me.bscore, gfx.DisplayWidth() * 4 / 5, 20)

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
