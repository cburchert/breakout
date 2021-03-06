package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ballImage *ebiten.Image
)

const (
	ballRadius               = 10
	ballSpeed        float64 = 500
	ballAcceleration         = 1.001
)

func init() {
	ballImage = ebiten.NewImage(ballRadius*2, ballRadius*2)
	ballImage.Fill(color.RGBA{100, 100, 100, 255})
}

type Ball struct {
	x  float64
	y  float64
	dx float64
	dy float64
}

func (b *Ball) Update() {
	b.x += b.dx / fps
	b.y += b.dy / fps
	b.dx *= ballAcceleration
	b.dy *= ballAcceleration
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(ballImage, &op)
}

func (b *Ball) BounceFromCollision(t CollisionType) {
	switch t {
	case CollidingFromLeft:
		b.dx *= -1
	case CollidingFromRight:
		b.dx *= -1
	case CollidingFromTop:
		b.dy *= -1
	case CollidingFromBottom:
		b.dy *= -1
	}
}

func (b *Ball) BoundingBox() Rectangle {
	return Rectangle{b.x, b.y, ballRadius * 2, ballRadius * 2}
}

// SetAngle makes the ball go into a specific direction while keeping the same speed.
// The angle is in radians, zero angle points along the positive X axis.
func (b *Ball) SetAngle(angle float64) {
	speed := math.Sqrt(b.dx*b.dx + b.dy*b.dy)
	b.dx = speed * math.Cos(angle)
	b.dy = speed * math.Sin(angle)
}
