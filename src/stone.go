package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	stoneImage *ebiten.Image
)

const ()

func init() {
	stoneImage = ebiten.NewImage(200, 100)
	stoneImage.Fill(color.RGBA{50, 50, 50, 255})
}

type Stone struct {
	x float64
	y float64
	w float64
	h float64
}

func (s *Stone) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	imgW, imgH := stoneImage.Size()
	op.GeoM.Scale(s.w/float64(imgW), s.h/float64(imgH))
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(stoneImage, &op)
}

func (s *Stone) BoundingBox() Rectangle {
	return Rectangle{s.x, s.y, s.w, s.h}
}
