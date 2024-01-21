package objects

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	img     *ebiten.Image
	h       int
	w       int
	x       int
	y       int
	offsetX float64
	offsetY float64
}

func NewPlayer() Object {
	img, _, err := ebitenutil.NewImageFromFile("./assets/png/player.png")

	if err != nil {
		log.Fatal(err)
	}

	w, h := img.Size()

	return &Player{
		img:     img,
		w:       w,
		h:       h,
		x:       320 - 32,
		y:       320 - 32,
		offsetX: float64(w),
		offsetY: float64(h),
	}
}

func (p *Player) Update(_ *ebiten.Image) {

}

func (p *Player) Draw(target *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.x), float64(p.y))

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.x -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.x += 5
	}

	target.DrawImage(p.img, op)
	return nil
}
