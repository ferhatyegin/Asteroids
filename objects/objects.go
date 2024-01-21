package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update(*ebiten.Image)
	Draw(*ebiten.Image) error
}
