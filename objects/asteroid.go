package objects

import (
	"log"
	"math/rand"

	"github.com/ferhatyegin/goAsteroids/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Asteroid struct {
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

func NewAsteroid() Object {
	img, _, err := ebitenutil.NewImageFromFile("./assets/png/asteroid_lg.png")

	if err != nil {
		log.Fatal(err)
	}

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	//Create random moving direction
	randomDx := rand.Intn(3) - 1
	randomDy := rand.Intn(3) - 1

	//If both numbers are 0 then create either -1 or 1
	if randomDx == 0 && randomDy == 0 {
		randomDx = rand.Intn(2)*2 - 1
		randomDy = rand.Intn(2)*2 - 1
	}

	randomX := rand.Intn(641)
	randomY := rand.Intn(641)

	randRot := (rand.Float64() * 0.04) - 0.02

	return &Asteroid{
		img:           img,
		w:             w,
		h:             h,
		x:             float64(randomX),
		y:             float64(randomY),
		dx:            float64(randomDx),
		dy:            float64(randomDy),
		offsetX:       float64(w) / 2,
		offsetY:       float64(h) / 2,
		rotationAngle: randRot,
	}
}

func (a *Asteroid) Update() error {

	utils.WrapCoordinates(a.x, a.y, &a.x, &a.y)

	if a.rotationAngle < 0 {
		a.rotationAngle -= 0.02
	} else {
		a.rotationAngle += 0.02
	}

	a.x += a.dx
	a.y += a.dy

	// fmt.Printf("Dx: %.2f - Dy: %.2f - Rot: %.2f\n", a.dx, a.dy, a.rotationAngle)

	return nil
}

func (a *Asteroid) Draw(target *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	op.GeoM.Translate(-float64(a.offsetX), -float64(a.offsetY))
	op.GeoM.Rotate(a.rotationAngle)
	op.GeoM.Translate(float64(a.x), float64(a.y))

	utils.DrawWrapped(target, a.img, a.x, a.y, a.rotationAngle)
	target.DrawImage(a.img, op)
	return nil
}
