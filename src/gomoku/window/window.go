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

type MenuRes struct {
	solo	*gfx.Text
	duo	*gfx.Text
}

type BoardRes struct {
	board		*gfx.Image
	Wscore		*gfx.Text
	Bscore		*gfx.Text
	St		bool
	Stones		[][]*Stone	
}

type Drawer struct {
	backgrnd	*gfx.Image
	title		*gfx.Text
	quitGame	*gfx.Text
	score		*gfx.Text
	mturn		*gfx.Text
	black_stone	*gfx.Image
	white_stone	*gfx.Image
	anim		*gfx.Animation
	Font		*gfx.Font
	Turn		bool
	GameState	string
	Board_res	BoardRes
	Menu_res	MenuRes
}

func (me *Drawer) initMenu() bool {
	me.Menu_res.solo = me.Font.Write("Play against Computer")
	me.Menu_res.duo = me.Font.Write("Play against Human")
	return true
}

func (me *Drawer) initGame() bool {
	me.Board_res.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight() / 20,
		gfx.DisplayHeight() / 20)
	me.Board_res.Stones = nil
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
		me.Board_res.Stones = append(me.Board_res.Stones, row)
	}
	me.Board_res.Wscore = me.Font.Write(strconv.Itoa(0))
	me.Board_res.Bscore = me.Font.Write(strconv.Itoa(0))
	return true
}

func (me *Drawer) Init() bool {
	me.backgrnd = gfx.LoadImageSize("ressources/woodback.jpg", gfx.DisplayWidth(), gfx.DisplayHeight())
		me.Font = gfx.LoadFont("ressources/MotionControl-Bold.otf", 46)
	me.Font.SetRGB(25, 25, 25)
	me.title = me.Font.Write("GoMoku")
	me.score = me.Font.Write("Taken")
	me.mturn = me.Font.Write("Turn")
	me.quitGame = me.Font.Write("Press escape key to quit")
	if (me.initGame() && me.initMenu()) {
		return true
	}
	return false
}

func (me *Drawer) drawGameBoard(c *gfx.Canvas) {
	c.DrawText(me.title, gfx.DisplayWidth() * 4 / 5, 0)
	c.DrawText(me.mturn, gfx.DisplayWidth() * 3 / 4, gfx.DisplayHeight() / 4)
	if me.Turn {
		c.DrawImage(me.white_stone, gfx.DisplayWidth() * 5 / 6, gfx.DisplayHeight() / 4 + 10)
	} else {
		c.DrawImage(me.black_stone, gfx.DisplayWidth() * 5 / 6, gfx.DisplayHeight() / 4 + 10)
	}
	c.DrawText(me.score, gfx.DisplayWidth() * 3 / 4, 4 * gfx.DisplayHeight() / 10)
	c.DrawImage(me.black_stone, gfx.DisplayWidth() * 10 / 13, 5 * gfx.DisplayHeight() / 11 + 8)
	c.DrawText(me.Board_res.Wscore, gfx.DisplayWidth() * 5 / 6, 5 * gfx.DisplayHeight() / 11)
	c.DrawImage(me.white_stone, gfx.DisplayWidth() * 10 / 13, 6 * gfx.DisplayHeight() / 11 + 8)
	c.DrawText(me.Board_res.Bscore, gfx.DisplayWidth() * 5 / 6, 6 * gfx.DisplayHeight() / 11)

	c.PushViewport(0, 0, gfx.DisplayWidth(), gfx.DisplayWidth())
	{
		c.DrawImage(me.Board_res.board, 0, 0)
		
		for i := range(me.Board_res.Stones) {
			for j := range(me.Board_res.Stones[i]) {
				if me.Board_res.Stones[i][j].Visible {
					if me.Board_res.Stones[i][j].White == true {
						c.DrawImage(me.white_stone, me.Board_res.Stones[i][j].X,
							me.Board_res.Stones[i][j].Y)
					} else {
						c.DrawImage(me.black_stone, me.Board_res.Stones[i][j].X,
							me.Board_res.Stones[i][j].Y)
					}
				}
			}
		}
	}
	c.PopViewport()
}

func (me *Drawer) drawMenu(c *gfx.Canvas) {
	c.DrawText(me.title, 5 * gfx.DisplayWidth() / 12, gfx.DisplayHeight() * 1 / 10)
	c.SetRGB(200,200,200)
	c.FillRect(4 * gfx.DisplayWidth() / 14, gfx.DisplayHeight() * 4 / 10,
		8 * gfx.DisplayWidth() / 18, gfx.DisplayHeight() / 11)
	c.SetRGB(200,200,200)
	c.FillRect(4 * gfx.DisplayWidth() / 14, gfx.DisplayHeight() * 5 / 10,
		8 * gfx.DisplayWidth() / 18, gfx.DisplayHeight() / 11)
	c.DrawText(me.Menu_res.solo, 6 * gfx.DisplayWidth() / 20, gfx.DisplayHeight() * 4 / 10)
	c.DrawText(me.Menu_res.duo, 6 * gfx.DisplayWidth() / 19, gfx.DisplayHeight() * 5 / 10)
	c.DrawText(me.quitGame, 6 * gfx.DisplayWidth() / 22, gfx.DisplayHeight() * 9 / 10)
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	c.SetRGB(255,255,255)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())
	c.DrawImage(me.backgrnd, gfx.DisplayWidth() * 0, 0)
	
	if me.GameState == "gameOn" {
		me.drawGameBoard(c)
	} else if me.GameState == "menu" {
		me.drawMenu(c)
	}
}
