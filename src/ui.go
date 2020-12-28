package main

import (
	"image/color"
	"log"

	data "github.com/cburchert/breakout/src/data/generated"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	bottomBarBorderColor  = color.RGBA{150, 150, 150, 255}
	bottomBarBorderHeight = 2.
	bottomBarHeight       = 30.

	defaultFontFace24    font.Face
	defaultFontFace18    font.Face
	startHint            ScreenText
	bottomBarBorderImage *ebiten.Image
)

func init() {
	initFonts()
	initStartHint()
	initBottomBarBorder()
}

func initFonts() {
	mPlus1Font, err := opentype.Parse(data.MPlus1PRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	faceOption := opentype.FaceOptions{Size: 24, DPI: 72, Hinting: font.HintingFull}
	defaultFontFace24, err = opentype.NewFace(mPlus1Font, &faceOption)
	if err != nil {
		log.Fatal(err)
	}
	faceOption.Size = 18
	defaultFontFace18, err = opentype.NewFace(mPlus1Font, &faceOption)
	if err != nil {
		log.Fatal(err)
	}
}

func initStartHint() {
	startHint.text = "Press Space to release a ball."
	startHint.fontface = &defaultFontFace24
	startHint.color = color.White
	startHint.SetCenterPosition(screenW/2, 500)
}

func initBottomBarBorder() {
	bottomBarBorderImage = ebiten.NewImage(1, 1)
	bottomBarBorderImage.Fill(bottomBarBorderColor)
}

type ScreenText struct {
	x, y     float64
	text     string
	fontface *font.Face
	color    color.Color
}

func (t *ScreenText) GetSize() (x, y float64) {
	boundRect := text.BoundString(*t.fontface, t.text)
	return float64(boundRect.Max.X - boundRect.Min.X),
		float64(boundRect.Max.Y - boundRect.Min.Y)
}

func (t *ScreenText) SetCenterPosition(x, y float64) {
	textW, textH := t.GetSize()
	t.x = x - textW/2
	t.y = y - textH/2
}

func (t *ScreenText) Draw(screen *ebiten.Image) {
	// text.Draw renders to a shifted position, we can correct this by
	// inspecting the boundRect.
	boundRect := text.BoundString(*t.fontface, t.text)
	fixedX := int(t.x) - boundRect.Min.X
	fixedY := int(t.y) - boundRect.Min.Y

	text.Draw(screen, t.text, *t.fontface, fixedX, fixedY, t.color)
}

func DrawStartHint(screen *ebiten.Image) {
	startHint.Draw(screen)
}

func DrawHelpText(screen *ebiten.Image) {
	txt := ScreenText{
		x:        10,
		y:        screenH - bottomBarHeight + 8,
		text:     "A: Move Left  D: Move Right",
		fontface: &defaultFontFace18,
		color:    color.White}
	txt.Draw(screen)
}

func DrawBottomBar(screen *ebiten.Image) {
	// Border
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(screenW, bottomBarBorderHeight)
	op.GeoM.Translate(0, screenH-bottomBarHeight-bottomBarBorderHeight)
	screen.DrawImage(bottomBarBorderImage, &op)

	// Background
	op = ebiten.DrawImageOptions{}
	op.ColorM.Scale(0.2, 0.2, 0.2, 1.)
	op.GeoM.Scale(screenW, bottomBarHeight)
	op.GeoM.Translate(0, screenH-bottomBarHeight)
	screen.DrawImage(bottomBarBorderImage, &op)

	// Help Text
	DrawHelpText(screen)
}
