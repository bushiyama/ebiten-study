package shootingstar

import (
"bytes"
"image/color"
"log"

"github.com/hajimehoshi/ebiten/v2"
"github.com/hajimehoshi/ebiten/v2/text/v2"
"golang.org/x/image/font/gofont/goregular"
)

const (
screenWidth  = 640
screenHeight = 480
)

var (
faceSource *text.GoTextFaceSource
)

type Scene struct {
	textX  float64
	textY  float64
	textVX float64
	textVY float64
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	faceSource = s
}

func New() *Scene {
	return &Scene{
		textX:  screenWidth - 20,
		textY:  20,
		textVX: -2.0,
		textVY: 2.0,
	}
}

func (s *Scene) Update() error {
	s.textX += s.textVX
	s.textY += s.textVY
	if s.textX < -100 || s.textY > screenHeight+50 {
		s.textX = screenWidth - 20
		s.textY = 20
	}
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x20, 0x20, 0x40, 0xff})

	face := &text.GoTextFace{
		Source: faceSource,
		Size:   32,
	}

	op := &text.DrawOptions{}
	op.GeoM.Rotate(-0.785398)
	op.GeoM.Translate(s.textX, s.textY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "Hello, World!", face, op)
}

func (s *Scene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
