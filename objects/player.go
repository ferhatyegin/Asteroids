package objects

import (
	"log"
	"math"

	"github.com/ferhatyegin/goAsteroids/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	img           *ebiten.Image
	h             int
	w             int
	x             float64
	y             float64
	dx            float64
	dy            float64
	offsetX       float64
	offsetY       float64
	rotationAngle float64
}

func NewPlayer() Object {
	img, _, err := ebitenutil.NewImageFromFile("./assets/png/player.png")

	if err != nil {
		log.Fatal(err)
	}

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	return &Player{
		img:     img,
		w:       w,
		h:       h,
		x:       320,
		y:       320,
		offsetX: float64(w) / 2,
		offsetY: float64(h) / 2,
	}
}

func (p *Player) Update() error {

	utils.WrapCoordinates(p.x, p.y, &p.x, &p.y)

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.rotationAngle -= 0.07
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.rotationAngle += 0.07
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.dx += math.Sin(p.rotationAngle) * 0.2
		p.dy += -math.Cos(p.rotationAngle) * 0.2
	}

	p.x += p.dx
	p.y += p.dy

	return nil
}

func (p *Player) Draw(target *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	op.GeoM.Translate(-float64(p.offsetX), -float64(p.offsetY))
	op.GeoM.Rotate(p.rotationAngle)
	op.GeoM.Translate(float64(p.x), float64(p.y))

	utils.DrawWrapped(target, p.img, p.x, p.y, p.rotationAngle)
	target.DrawImage(p.img, op)
	return nil
}
