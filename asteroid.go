package asteroid

import (
	"fmt"
	"log"

	"github.com/ferhatyegin/goAsteroids/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	win_W = 640
	win_H = 640
)

type Game struct {
	objects []objects.Object
}

func (g *Game) Update() error {
	for _, o := range g.objects {
		if err := o.Update(); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d\tTPS:%d\nLEFT/RIGHT ARROW TO ROTATE\nUP/DOWN ARROW TO FORWARD/BACKWARD\n", int(ebiten.ActualFPS()), int(ebiten.ActualTPS())))

	for _, o := range g.objects {
		if err := o.Draw(screen); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	ebiten.SetWindowSize(win_W, win_H)
	ebiten.SetWindowTitle("Go Asteroids !")

	g := &Game{}

	g.objects = []objects.Object{
		objects.NewPlayer(),
	}

	return g
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
