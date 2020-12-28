package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	stoneImageW          = 120
	stoneImageH          = 50
	stoneBorderThickness = 2
)

var (
	stoneColor       = color.RGBA{100, 100, 100, 255}
	stoneBorderColor = color.RGBA{0, 0, 0, 255}
	stoneImage       *ebiten.Image
)

func init() {
	MakeStoneImage()
}

func MakeStoneImage() {
	stoneImage = ebiten.NewImage(stoneImageW, stoneImageH)
	stoneImage.Fill(stoneColor)
	for x := 0; x < stoneBorderThickness; x++ {
		for y := 0; y < stoneImageH; y++ {
			stoneImage.Set(x, y, stoneBorderColor)
			stoneImage.Set(stoneImageW-x-1, y, stoneBorderColor)
		}
	}
	for x := 0; x < stoneImageW; x++ {
		for y := 0; y < stoneBorderThickness; y++ {
			stoneImage.Set(x, y, stoneBorderColor)
			stoneImage.Set(x, stoneImageH-y-1, stoneBorderColor)
		}
	}
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
