package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	defaultBarWidth         = 400
	defaultBarSpeed         = 1920
	barHeight               = 40
	barY            float64 = 1000
)

var (
	barImage *ebiten.Image
)

func init() {
	barImage = ebiten.NewImage(defaultBarWidth, barHeight)
	barImage.Fill(color.RGBA{180, 80, 80, 255})
}

type Bar struct {
	x     float64
	w     float64
	speed float64
}

func NewBar() *Bar {
	return &Bar{screenW/2 - defaultBarWidth/2, defaultBarWidth, defaultBarSpeed}
}

func (b *Bar) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		b.x -= b.speed / fps
		if b.x < 0 {
			b.x = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		b.x += b.speed / fps
		if b.x+b.w > screenW {
			b.x = screenW - b.w
		}
	}
}

func (b *Bar) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	barImageW, barImageH := barImage.Size()
	op.GeoM.Scale(b.w/float64(barImageW), barHeight/float64(barImageH))
	op.GeoM.Translate(b.x, barY)
	screen.DrawImage(barImage, &op)
}

func (b *Bar) BoundingBox() Rectangle {
	return Rectangle{b.x, barY, b.w, barHeight}
}
