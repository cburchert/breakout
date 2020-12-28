package main

import (
	"container/list"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenW = 1920
	screenH = 1080

	stoneW       float64 = 120
	stoneH       float64 = 50
	stoneYOffset         = 100

	fps = 60
)

var (
	screenBackgroundColor = color.RGBA{20, 20, 20, 255}
)

type Game struct {
	bar    *Bar
	ball   *Ball
	stones *list.List

	score int
	lives int
}

func NewGame() *Game {
	g := &Game{NewBar(), nil, list.New(), 0, 0}
	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.SpawnStones()
	g.score = 0
	g.lives = 3
}

func (g *Game) SpawnStones() {
	g.stones.Init()
	for x := 0; x < 16; x++ {
		for y := 0; y < 3; y++ {
			g.stones.PushBack(&Stone{float64(x) * stoneW, float64(y)*stoneH + stoneYOffset, stoneW, stoneH})
		}
	}
}

func (g *Game) SpawnBall() {
	g.ball = &Ball{screenW / 2, screenH / 2, ballSpeed, ballSpeed}
}

func (g *Game) CheckCollisions() {
	if g.ball != nil {
		ballBoundingBox := g.ball.BoundingBox()
		// Ball vs bar
		g.ball.BounceFromCollision(g.bar.BoundingBox().CollisionCase(ballBoundingBox))

		// Ball vs stones
		for s := g.stones.Front(); s != nil; s = s.Next() {
			collisionType := s.Value.(*Stone).BoundingBox().CollisionCase(ballBoundingBox)
			g.ball.BounceFromCollision(collisionType)
			if collisionType != NotColliding {
				g.stones.Remove(s)
				g.score++
			}
		}

		// Ball vs screen borders
		if g.ball.x < 0 {
			g.ball.dx *= -1
			g.ball.x = 0
		} else if g.ball.x > screenW-2*ballRadius {
			g.ball.x = screenW - 2*ballRadius
			g.ball.dx *= -1
		} else if g.ball.y < 0 {
			g.ball.y = 0
			g.ball.dy *= -1
		} else if g.ball.y > screenH {
			g.ball = nil
			g.lives -= 1
			if g.lives == 0 {
				g.Restart()
			}
		}
	}
}

// Advance a frame
func (g *Game) Update() error {
	g.bar.Update()

	if g.ball == nil {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.SpawnBall()
		}
	} else {
		g.ball.Update()
	}
	g.CheckCollisions()

	return nil
}

// Render a frame
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(screenBackgroundColor)

	g.bar.Draw(screen)
	if g.ball != nil {
		g.ball.Draw(screen)
	}
	for s := g.stones.Front(); s != nil; s = s.Next() {
		s.Value.(*Stone).Draw(screen)
	}

	if g.ball == nil {
		DrawStartHint(screen)
	}

	DrawBottomBar(screen, g.score, g.lives)
}

// Returns the size of the viewport we would like for any given size of the screen.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Breakout")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
