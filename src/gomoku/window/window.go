package window

import (
	//	"os"
	//	"fmt"
	"gomoku/bmap"
	"strconv"
	"time"

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

type OptionsRes struct {
	Op1, Op2    bool
	cross       *gfx.Image
	optionRule1 *gfx.Text
	optionRule2 *gfx.Text
	restart     *gfx.Text
	back        *gfx.Text
	exit        *gfx.Text
}

type MenuRes struct {
	solo    *gfx.Text
	duo     *gfx.Text
	options *gfx.Text
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
	BoardRes
	MenuRes
	EndRes
	OptionsRes
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
	GameType    string
	GameState   string
	WinnerColor bool
}

func (me *Drawer) initOptions() bool {
	me.OptionsRes.Op1, me.OptionsRes.Op2 = true, true
	me.OptionsRes.optionRule1 = me.Font.Write("Unbroken row")
	me.OptionsRes.optionRule2 = me.Font.Write("Three and three")
	me.OptionsRes.exit = me.Font.Write("Exit...")
	me.OptionsRes.restart = me.Font.Write("Restart")
	me.OptionsRes.back = me.Font.Write("Back")
	return true
}

func (me *Drawer) initMenu() bool {
	me.MenuRes.solo = me.Font.Write("Play against Computer")
	me.MenuRes.duo = me.Font.Write("Play against Human")
	me.MenuRes.options = me.Font.Write("Options")
	return true
}

func (me *Drawer) initGame() bool {
	me.OptionsRes.cross = gfx.LoadImageSize("ressources/Cross.png", gfx.DisplayHeight()/13, gfx.DisplayHeight()/13)
	me.BoardRes.badmove = gfx.NewAnimation(500)
	me.BoardRes.badmove.LoadImageSize("ressources/Red.png", gfx.DisplayHeight()/20, gfx.DisplayHeight()/20)
	me.BoardRes.badmove.LoadImageSize("ressources/Red.png", gfx.DisplayHeight()/20, gfx.DisplayHeight()/20)
	me.BoardRes.badmove.LoadImageSize("ressources/Empty.png", gfx.DisplayHeight()/20, gfx.DisplayHeight()/20)
	me.BoardRes.board = gfx.LoadImageSize("ressources/board.png", gfx.DisplayHeight(), gfx.DisplayHeight())
	me.black_stone = gfx.LoadImageSize("ressources/bstone.png", gfx.DisplayHeight()/20,
		gfx.DisplayHeight()/20)
	me.white_stone = gfx.LoadImageSize("ressources/wstone.png", gfx.DisplayHeight()/20,
		gfx.DisplayHeight()/20)
	me.BoardRes.Stones = nil
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
		me.BoardRes.Stones = append(me.BoardRes.Stones, row)
	}
	me.BoardRes.Wscore = me.Font.Write(strconv.Itoa(0))
	me.BoardRes.Bscore = me.Font.Write(strconv.Itoa(0))
	me.BoardRes.Restart = me.Font.Write("Restart game")
	me.EndRes.end = me.Font.Write("Winner")
	me.EndRes.drawEndText = me.Font.Write("Draw")
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
	if me.initGame() && me.initMenu() && me.initOptions() {
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
	c.DrawText(me.BoardRes.Wscore, gfx.DisplayWidth()*5/6, 6*gfx.DisplayHeight()/12)
	c.DrawImage(me.white_stone, gfx.DisplayWidth()*10/13, 7*gfx.DisplayHeight()/12+8)
	c.DrawText(me.BoardRes.Bscore, gfx.DisplayWidth()*5/6, 7*gfx.DisplayHeight()/12)

	c.PushViewport(0, 0, gfx.DisplayWidth(), gfx.DisplayWidth())
	{
		c.DrawImage(me.BoardRes.board, 0, 0)

		for i := 0; i < bmap.Map_size; i++ {
			for j := 0; j < bmap.Map_size; j++ {
				if bmap.IsVisible(&bmap.Map, i, j) {
					if bmap.IsWhite(&bmap.Map, i, j) {
						c.DrawImage(me.white_stone, me.BoardRes.Stones[i][j].X,
							me.BoardRes.Stones[i][j].Y)
					} else {
						c.DrawImage(me.black_stone, me.BoardRes.Stones[i][j].X,
							me.BoardRes.Stones[i][j].Y)
					}
				}
			}
		}
	}
	c.PopViewport()
	if me.BoardRes.BadX > 0 {
		timer := time.NewTimer(time.Second * 4)
		go func() {
			<-timer.C
			me.BoardRes.BadX = 0
		}()

		c.DrawAnimation(me.BoardRes.badmove, me.BoardRes.BadX, me.BoardRes.BadY)
	}

	c.SetRGBA(133, 94, 66, 140)

	c.FillRect(10*gfx.DisplayWidth()/14, gfx.DisplayHeight()*9/11,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11)
	c.DrawText(me.MenuRes.options, 14*gfx.DisplayWidth()/19, gfx.DisplayHeight()*10/12)

	c.SetRGBA(133, 120, 76, 140)
	c.FillRect(10*gfx.DisplayWidth()/14, gfx.DisplayHeight()*10/11,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11)
	c.DrawText(me.BoardRes.Restart, 14*gfx.DisplayWidth()/19, gfx.DisplayHeight()*11/12)
}

func (me *Drawer) drawMenu(c *gfx.Canvas) {
	c.DrawText(me.title, 5*gfx.DisplayWidth()/12, gfx.DisplayHeight()*1/10)
	c.SetRGBA(133, 94, 66, 150)
	c.FillRoundedRect(4*gfx.DisplayWidth()/14, gfx.DisplayHeight()*4/10,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11, 10)
	c.FillRoundedRect(4*gfx.DisplayWidth()/14, gfx.DisplayHeight()*5/10,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11, 10)
	c.FillRoundedRect(4*gfx.DisplayWidth()/14, gfx.DisplayHeight()*6/10,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11, 10)
	c.DrawText(me.MenuRes.solo, 6*gfx.DisplayWidth()/20, gfx.DisplayHeight()*4/10)
	c.DrawText(me.MenuRes.duo, 6*gfx.DisplayWidth()/19, gfx.DisplayHeight()*5/10)
	c.DrawText(me.MenuRes.options, 5*gfx.DisplayWidth()/12, gfx.DisplayHeight()*6/10)
	c.DrawText(me.quitGame, 6*gfx.DisplayWidth()/22, gfx.DisplayHeight()*9/10)
}

func (me *Drawer) drawEnd(c *gfx.Canvas) {
	c.SetRGBA(255, 255, 255, 60)
	c.FillRect(0, gfx.DisplayHeight()*20/50, gfx.DisplayWidth(), gfx.DisplayHeight()*10/50)
	if !me.EndRes.DrawEnd {
		c.DrawText(me.EndRes.end, 91*gfx.DisplayWidth()/200, 8*gfx.DisplayHeight()/20)
		if me.WinnerColor {
			c.DrawImage(me.white_stone, gfx.DisplayWidth()/2, gfx.DisplayHeight()/2)
		} else {
			c.DrawImage(me.black_stone, gfx.DisplayWidth()/2, gfx.DisplayHeight()/2)
		}
	} else {
		c.DrawText(me.EndRes.drawEndText, 91*gfx.DisplayWidth()/200, 8*gfx.DisplayHeight()/20)
	}
	c.FillRect(10*gfx.DisplayWidth()/14, gfx.DisplayHeight()*10/11,
		8*gfx.DisplayWidth()/18, gfx.DisplayHeight()/11)
	c.DrawText(me.BoardRes.Restart, 14*gfx.DisplayWidth()/19, gfx.DisplayHeight()*11/12)
	c.DrawText(me.quitGame, 4*gfx.DisplayWidth()/22, gfx.DisplayHeight()*11/12)
	c.SetRGBA(255, 255, 255, 255)
}

func (me *Drawer) drawOptions(c *gfx.Canvas) {
	c.SetRGBA(150, 150, 150, 70)
	c.FillRect(gfx.DisplayHeight()*20/50, 0, gfx.DisplayHeight()*30/50, gfx.DisplayWidth())
	elems := []*gfx.Text{
		me.OptionsRes.optionRule1,
		me.OptionsRes.optionRule2,
		me.OptionsRes.restart,
		me.OptionsRes.exit,
		me.OptionsRes.back,
	}
	c.SetRGBA(150, 150, 150, 70)
	c.SetRGBA(133, 94, 66, 150)

	for i, elem := range elems {
		if i >= 2 {
			c.FillRoundedRect(gfx.DisplayHeight()*27/50, gfx.DisplayHeight()/4+(i*gfx.DisplayHeight()/9), 150, 50, 7)
		}
		c.DrawText(elem, gfx.DisplayHeight()*28/50,
			gfx.DisplayHeight()/4+(i*gfx.DisplayHeight()/9))
	}

	c.SetRGBA(0, 0, 0, 120)
	c.FillRect(gfx.DisplayHeight()*22/50,
		gfx.DisplayHeight()/4+(0*gfx.DisplayHeight()/9), 50, 50)
	c.FillRect(gfx.DisplayHeight()*22/50,
		gfx.DisplayHeight()/4+(1*gfx.DisplayHeight()/9), 50, 50)
	if me.OptionsRes.Op1 {
		c.DrawImage(me.OptionsRes.cross, gfx.DisplayHeight()*22/50,
			gfx.DisplayHeight()/4+(0*gfx.DisplayHeight()/9))
	}
	if me.OptionsRes.Op2 {
		c.DrawImage(me.OptionsRes.cross, gfx.DisplayHeight()*22/50,
			gfx.DisplayHeight()/4+(1*gfx.DisplayHeight()/9))
	}
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
	} else if me.GameState == "options" {
		me.drawOptions(c)
	}
}
