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
	backgrnd	*gfx.Image
	title		*gfx.Text
	score		*gfx.Text
	Wscore		*gfx.Text
	Bscore		*gfx.Text
	mturn		*gfx.Text
	black_stone	*gfx.Image
	white_stone	*gfx.Image
	anim		*gfx.Animation
	St		bool
	Stones		[][]*Stone
	Font		*gfx.Font
	Turn		bool
	GameState	string
}

func (me *Drawer) Init() bool {
	me.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.backgrnd = gfx.LoadImageSize("ressources/woodback.jpg", gfx.DisplayWidth(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.Font = gfx.LoadFont("ressources/MotionControl-Bold.otf", 46)
	me.Font.SetRGB(25, 25, 25)
	me.title = me.Font.Write("GoMoku")
	me.score = me.Font.Write("Taken")
	me.mturn = me.Font.Write("Turn")
	me.Stones = nil
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

func (me *Drawer) drawGameBoard(c *gfx.Canvas) {
	c.DrawText(me.mturn, gfx.DisplayWidth() * 3 / 4, gfx.DisplayHeight() / 4)
	if me.Turn {
		c.DrawImage(me.white_stone, gfx.DisplayWidth() * 5 / 6, gfx.DisplayHeight() / 4 + 10)
	} else {
		c.DrawImage(me.black_stone, gfx.DisplayWidth() * 5 / 6, gfx.DisplayHeight() / 4 + 10)
	}
	c.DrawText(me.score, gfx.DisplayWidth() * 3 / 4, 4 * gfx.DisplayHeight() / 10)
	c.DrawImage(me.black_stone, gfx.DisplayWidth() * 10 / 13, 5 * gfx.DisplayHeight() / 11 + 8)
	c.DrawText(me.Wscore, gfx.DisplayWidth() * 5 / 6, 5 * gfx.DisplayHeight() / 11)
	c.DrawImage(me.white_stone, gfx.DisplayWidth() * 10 / 13, 6 * gfx.DisplayHeight() / 11 + 8)
	c.DrawText(me.Bscore, gfx.DisplayWidth() * 5 / 6, 6 * gfx.DisplayHeight() / 11)

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

func (me *Drawer) Draw(c *gfx.Canvas) {
	c.SetRGB(255,255,255)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())
	c.DrawImage(me.backgrnd, gfx.DisplayWidth() * 0, 0)
	
	c.DrawText(me.title, gfx.DisplayWidth() * 4 / 5, 0)
	// if me.GameState == "menu" {
	// 	me.drawMenu(c)
	// }
	if me.GameState == "gameOn" {
		me.drawGameBoard(c)
	}
}
