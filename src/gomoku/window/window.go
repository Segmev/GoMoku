package window

import (
	//	"os"
	//	"fmt"
	"gomoku/bmap"
	"strconv"

	"github.com/gtalent/starfish/gfx"
)

type InfosStone struct {
	Ipos, Jpos int
	OppoSt     [3][3]int
	TeamSt     [3][3]int
	Breakable  bool
}

type Stone struct {
	X, Y    int
	Visible bool
	Color   bool
	Infos   InfosStone
}

type MenuRes struct {
	solo *gfx.Text
	duo  *gfx.Text
}

type EndRes struct {
	end         *gfx.Text
	drawEndText *gfx.Text
	DrawEnd     bool
}

type BoardRes struct {
	badmove    *gfx.Animation
	BadX, BadY int
	board      *gfx.Image
	Wscore     *gfx.Text
	Bscore     *gfx.Text
	Restart    *gfx.Text
	St         bool
	Stones     [][]*Stone
}

type Drawer struct {
	backgrnd    *gfx.Image
	title       *gfx.Text
	quitGame    *gfx.Text
	score       *gfx.Text
	mturn       *gfx.Text
	black_stone *gfx.Image
	white_stone *gfx.Image
	anim        *gfx.Animation
	Font        *gfx.Font
	Turn        bool
	GameState   string
	WinnerColor bool
	Board_res   BoardRes
	Menu_res    MenuRes
	End_res     EndRes
}

func (me *Drawer) initMenu() bool {
	me.Menu_res.solo = me.Font.Write("Play against Computer")
	me.Menu_res.duo = me.Font.Write("Play against Human")
	return true
}

func (me *Drawer) initGame() bool {
	me.Board_res.badmove = gfx.NewAnimation(500)
	me.Board_res.badmove.LoadImageSize("ressources/Red.png", gfx.DisplayHeight()/20, gfx.DisplayHeight()/20)
	me.Board_res.badmove.LoadImageSize("ressources/Red.png", gfx.DisplayHeight()/20, gfx.DisplayHeight()/20)
	me.Board_res.badmove.LoadImageSize("ressources/Empty.png", gfx.DisplayHeight()/20, gfx.DisplayHeight()/20)
	me.Board_res.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight()/20,
		gfx.DisplayHeight()/20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight()/20,
		gfx.DisplayHeight()/20)
	me.Board_res.Stones = nil
	for i := 0; i <= 18; i++ {
		row := []*Stone{}
		for j := 0; j <= 18; j++ {
			stone := new(Stone)
			stone.Infos.Ipos, stone.Infos.Jpos = i, j
			stone.X = gfx.DisplayHeight()/88 + i*(gfx.DisplayHeight()/19)
			stone.Y = gfx.DisplayHeight()/88 + j*(gfx.DisplayHeight()/19)
			stone.Color, stone.Visible = true, false
			row = append(row, stone)
		}
		me.Board_res.Stones = append(me.Board_res.Stones, row)
	}
	me.Board_res.Wscore = me.Font.Write(strconv.Itoa(0))
	me.Board_res.Bscore = me.Font.Write(strconv.Itoa(0))
	me.Board_res.Restart = me.Font.Write("Restart game")
	me.End_res.end = me.Font.Write("Winner")
	me.End_res.drawEndText = me.Font.Write("Draw")
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
	if me.initGame() && me.initMenu() {
		return true
	}
	return false
}

func (me *Drawer) drawGameBoard(c *gfx.Canvas) {
	c.DrawText(me.title, gfx.DisplayWidth()*4/5, 0)
	c.DrawText(me.mturn, gfx.DisplayWidth()*3/4, gfx.DisplayHeight()/4)
	if me.Turn {
		c.DrawImage(me.white_stone, gfx.DisplayWidth()*6/7, gfx.DisplayHeight()/4+10)
	} else {
		c.DrawImage(me.black_stone, gfx.DisplayWidth()*6/7, gfx.DisplayHeight()/4+10)
	}
	c.DrawText(me.score, gfx.DisplayWidth()*3/4, 4*gfx.DisplayHeight()/10)
	c.DrawImage(me.black_stone, gfx.DisplayWidth()*10/13, 6*gfx.DisplayHeight()/12+8)
	c.DrawText(me.Board_res.Wscore, gfx.DisplayWidth()*5/6, 6*gfx.DisplayHeight()/12)
	c.DrawImage(me.white_stone, gfx.DisplayWidth()*10/13, 7*gfx.DisplayHeight()/12+8)
	c.DrawText(me.Board_res.Bscore, gfx.DisplayWidth()*5/6, 7*gfx.DisplayHeight()/12)

	c.PushViewport(0, 0, gfx.DisplayWidth(), gfx.DisplayWidth())
	{
		c.DrawImage(me.Board_res.board, 0, 0)

		for i := 0; i < bmap.Map_size; i++ {
			for j := 0; j < bmap.Map_size; j++ {
				if bmap.IsVisible(i, j) {
					if bmap.IsWhite(i, j) {
						c.DrawImage(me.white_stone, me.Board_res.Stones[i][j].X,
							me.Board_res.Stones[i][j].Y)
					} else {
						c.DrawImage(me.black_stone, me.Board_res.Stones[i][j].X,
							me.Board_res.Stones[i][j].Y)
					}
				}
			}
		}

		// for i := range me.Board_res.Stones {
		// 	for j := range me.Board_res.Stones[i] {
		// 		if me.Board_res.Stones[i][j].Visible {
		// 			if me.Board_res.Stones[i][j].Color == true {
		// 				c.DrawImage(me.white_stone, me.Board_res.Stones[i][j].X,
		// 					me.Board_res.Stones[i][j].Y)
		// 			} else {
		// 				c.DrawImage(me.black_stone, me.Board_res.Stones[i][j].X,
		// 					me.Board_res.Stones[i][j].Y)
		// 			}
		// 		}
		// 	}
		//}
	}
	c.PopViewport()
	if me.Board_res.BadX > 0 {
		c.DrawAnimation(me.Board_res.badmove, me.Board_res.BadX, me.Board_res.BadY)
	}
	c.SetRGBA(133, 94, 66, 140)
	c.FillRect(10*gfx.DisplayWidth()/14, gfx.DisplayHeight()*10/11,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11)
	c.DrawText(me.Board_res.Restart, 14*gfx.DisplayWidth()/19, gfx.DisplayHeight()*11/12)
}

func (me *Drawer) drawMenu(c *gfx.Canvas) {
	c.DrawText(me.title, 5*gfx.DisplayWidth()/12, gfx.DisplayHeight()*1/10)
	c.SetRGBA(133, 94, 66, 150)
	c.FillRoundedRect(4*gfx.DisplayWidth()/14, gfx.DisplayHeight()*4/10,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11, 10)
	c.FillRoundedRect(4*gfx.DisplayWidth()/14, gfx.DisplayHeight()*5/10,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11, 10)
	c.DrawText(me.Menu_res.solo, 6*gfx.DisplayWidth()/20, gfx.DisplayHeight()*4/10)
	c.DrawText(me.Menu_res.duo, 6*gfx.DisplayWidth()/19, gfx.DisplayHeight()*5/10)
	c.DrawText(me.quitGame, 6*gfx.DisplayWidth()/22, gfx.DisplayHeight()*9/10)
}

func (me *Drawer) drawEnd(c *gfx.Canvas) {
	c.SetRGBA(255, 255, 255, 60)
	c.FillRect(0, gfx.DisplayHeight()*20/50, gfx.DisplayWidth(), gfx.DisplayHeight()*10/50)
	if !me.End_res.DrawEnd {
		c.DrawText(me.End_res.end, 91*gfx.DisplayWidth()/200, 8*gfx.DisplayHeight()/20)
		if me.WinnerColor {
			c.DrawImage(me.white_stone, gfx.DisplayWidth()/2, gfx.DisplayHeight()/2)
		} else {
			c.DrawImage(me.black_stone, gfx.DisplayWidth()/2, gfx.DisplayHeight()/2)
		}
	} else {
		c.DrawText(me.End_res.drawEndText, 91*gfx.DisplayWidth()/200, 8*gfx.DisplayHeight()/20)
	}
	c.FillRect(10*gfx.DisplayWidth()/14, gfx.DisplayHeight()*10/11,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11)
	c.DrawText(me.Board_res.Restart, 14*gfx.DisplayWidth()/19, gfx.DisplayHeight()*11/12)
	c.DrawText(me.quitGame, 4*gfx.DisplayWidth()/22, gfx.DisplayHeight()*11/12)
	c.SetRGBA(255, 255, 255, 255)
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	c.SetRGB(255, 255, 255)
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())
	c.DrawImage(me.backgrnd, gfx.DisplayWidth()*0, 0)

	if me.GameState == "gameOn" {
		me.drawGameBoard(c)
	} else if me.GameState == "menu" {
		me.drawMenu(c)
	} else if me.GameState == "end" {
		me.drawEnd(c)
	}
}
