package rotate

import (
"bytes"
"image/color"
"log"
"math"

"github.com/hajimehoshi/ebiten/v2"
"github.com/hajimehoshi/ebiten/v2/text/v2"
"github.com/hajimehoshi/ebiten/v2/vector"
"golang.org/x/image/font/gofont/goregular"
)

const (
screenWidth  = 640
screenHeight = 480
)

var (
faceSource        *text.GoTextFaceSource
textRotationDir   = 1.0
circleRotationDir = -1.0
)

type Scene struct {
	angle float64
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	faceSource = s
}

func New() *Scene {
	return &Scene{}
}

func (s *Scene) Update() error {
	s.angle += 0.02
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x20, 0x20, 0x40, 0xff})

	radius := 100.0
	centerX := float64(screenWidth) / 2
	centerY := float64(screenHeight) / 2

	textStr := "Hello, World!"
	numPoints := len(textStr)

	for i := range numPoints {
		theta := float64(i) * 2 * math.Pi / float64(numPoints)
		phi := math.Pi / 3

		x := radius * math.Sin(phi) * math.Cos(theta+s.angle*textRotationDir)
		y := radius * math.Cos(phi)
		z := radius * math.Sin(phi) * math.Sin(theta+s.angle*textRotationDir)

		scale := 1.0 / (1.0 + z/300.0)
		if z < 0 {
			scale *= 0.3
		}

		screenX := centerX + x*scale
		screenY := centerY - y*scale

		col := color.RGBA{
			uint8(255 * scale),
			uint8(255 * scale),
			uint8(255 * scale),
			0xff,
		}

		char := string(textStr[i])
		op := &text.DrawOptions{}
		op.GeoM.Translate(screenX-3, screenY-4)
		op.ColorScale.ScaleWithColor(col)

		face := &text.GoTextFace{
			Source: faceSource,
			Size:   13,
		}
		text.Draw(screen, char, face, op)
	}

	for lat := 0; lat < 6; lat++ {
		phi := math.Pi * float64(lat) / 5
		for lng := 0; lng < 50; lng++ {
			theta := 2 * math.Pi * float64(lng) / 50

			x := radius * math.Sin(phi) * math.Cos(theta+s.angle*circleRotationDir)
			y := radius * math.Cos(phi)
			z := radius * math.Sin(phi) * math.Sin(theta+s.angle*circleRotationDir)

			scale := 1.0 / (1.0 + z/300.0)
			if z > 0 {
				screenX := centerX + x*scale
				screenY := centerY - y*scale
				col := color.RGBA{
					uint8(100 * scale),
					uint8(150 * scale),
					uint8(200 * scale),
					0x80,
				}
				vector.FillRect(screen, float32(screenX-1), float32(screenY-1), 2, 2, col, false)
			}
		}
	}
}

func (s *Scene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
