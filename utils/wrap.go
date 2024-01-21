package utils

import "github.com/hajimehoshi/ebiten/v2"

const (
	SCREEN_W = 640
	SCREEN_H = 640
)

// WrapCoordinates wraps coordinates that are going to be drawn on the screen.
func WrapCoordinates(ix float64, iy float64, ox *float64, oy *float64) {
	*ox = ix
	*oy = iy

	if ix < 0 {
		*ox = ix + float64(SCREEN_W)
	}
	if ix > float64(SCREEN_W) {
		*ox = ix - float64(SCREEN_W)
	}

	if iy < 0 {
		*oy = iy + float64(SCREEN_H)
	}
	if iy > float64(SCREEN_H) {
		*oy = iy - float64(SCREEN_H)
	}
}

// DrawWrapped actually draws the coordinates that are calculated with (WrappedCoordinates)
func DrawWrapped(screen *ebiten.Image, img *ebiten.Image, x, y, rotationAngle float64) {
	// Check if the player is going off the right edge
	imgW := img.Bounds().Dx()
	imgH := img.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}
	if x+float64(imgW)*2 > SCREEN_W {
		op.GeoM.Translate(-float64(imgW)/2, -float64(imgH)/2)
		op.GeoM.Rotate(rotationAngle)
		op.GeoM.Translate(x-SCREEN_W, y)
		screen.DrawImage(img, op)

	} else if x-float64(imgW)*2 < 0 {
		op.GeoM.Translate(-float64(imgW)/2, -float64(imgH)/2)
		op.GeoM.Rotate(rotationAngle)
		op.GeoM.Translate(x+SCREEN_W, y)
		screen.DrawImage(img, op)

	} else if y+float64(imgH)*2 > SCREEN_H {
		op.GeoM.Translate(-float64(imgW)/2, -float64(imgH)/2)
		op.GeoM.Rotate(rotationAngle)
		op.GeoM.Translate(x, y-SCREEN_H)
		screen.DrawImage(img, op)

	} else if y-float64(imgH)*2 < 0 {
		op.GeoM.Translate(-float64(imgW)/2, -float64(imgH)/2)
		op.GeoM.Rotate(rotationAngle)
		op.GeoM.Translate(x, y+SCREEN_H)
		screen.DrawImage(img, op)

	}
}
