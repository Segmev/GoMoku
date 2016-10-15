package window

import (
//	"os"
//	"fmt"
	"github.com/gtalent/starfish/gfx"
	"strconv"
)


type Stone struct {
	X	int
	Y	int
	Ipos	int
	Jpos	int
	Visible	bool
	White	bool
}

type Drawer struct {
	board		*gfx.Image
	text		*gfx.Text
	Wscore		*gfx.Text
	Bscore		*gfx.Text
	black_stone	*gfx.Image
	white_stone	*gfx.Image
	anim		*gfx.Animation
	St		bool
	Stones		[][]*Stone
	Font		*gfx.Font
}

func (me *Drawer) Init() bool {
	me.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)

	me.Font = gfx.LoadFont("ressources/LiberationSans-Bold.ttf", 26)
	me.Font.SetRGB(100, 100, 255)
	me.text = me.Font.Write("GoMoku")
	for i := 0; i <= 18; i++ {
		row := []*Stone{}
		for j := 0; j <= 18; j++ {
			stone := new(Stone)
			stone.Ipos, stone.Jpos = i, j
			stone.X = gfx.DisplayHeight() / 88 + i * (gfx.DisplayHeight() / 19)
			stone.Y = gfx.DisplayHeight() / 88 + j * (gfx.DisplayHeight() / 19)
			stone.White, stone.Visible = true, false
			row = append(row, stone)
		}
		me.Stones = append(me.Stones, row)
	}
	me.Wscore = me.Font.Write(strconv.Itoa(0))
	me.Bscore = me.Font.Write(strconv.Itoa(0))
	return true
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	c.SetRGB(55, 55, 55)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())
	me.Font.SetRGB(100, 100, 255)
	c.DrawText(me.text, gfx.DisplayWidth() * 3 / 4, 0)
	c.DrawText(me.Wscore, gfx.DisplayWidth() * 3 / 4, 20)
	c.DrawText(me.Bscore, gfx.DisplayWidth() * 4 / 5, 20)
	c.PushViewport(0, 0, gfx.DisplayWidth(), gfx.DisplayWidth())
	{
		c.DrawImage(me.board, 0, 0)
		
		for i := range(me.Stones) {
			for j := range(me.Stones[i]) {
				if me.Stones[i][j].Visible {
					if me.Stones[i][j].White == true {
						c.DrawImage(me.white_stone, me.Stones[i][j].X, me.Stones[i][j].Y)
					} else {
						c.DrawImage(me.black_stone, me.Stones[i][j].X, me.Stones[i][j].Y)
					}
				}
			}
		}
	}
	c.PopViewport()
}
