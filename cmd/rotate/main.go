package main

import (
	"bytes"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

type Game struct {
	angle float64
}

func init() {
	// GoフォントのTrueTypeデータを読み込み
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	faceSource = s
}

func (g *Game) Update() error {
	g.angle += 0.02
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 画面をクリア
	screen.Fill(color.RGBA{0x20, 0x20, 0x40, 0xff})

	// 球体の点を描画
	radius := 100.0
	centerX := float64(screenWidth) / 2
	centerY := float64(screenHeight) / 2

	// "Hello, World!" のテキストを球体表面に配置
	textStr := "Hello, World!"
	numPoints := len(textStr)

	for i := range numPoints {
		// 球面上の点を計算（緯度・経度ベース）
		theta := float64(i) * 2 * math.Pi / float64(numPoints)
		phi := math.Pi / 3 // 赤道付近に配置

		// 3D座標
		x := radius * math.Sin(phi) * math.Cos(theta+g.angle)
		y := radius * math.Cos(phi)
		z := radius * math.Sin(phi) * math.Sin(theta+g.angle)

		// Z座標による遠近感（奥にあるものは小さく）
		scale := 1.0 / (1.0 + z/300.0)
		if z < 0 { // 裏側は暗く
			scale *= 0.3
		}

		// 2D投影
		screenX := centerX + x*scale
		screenY := centerY - y*scale

		// 文字を描画
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

	// 追加の円周も描画して球体感を出す
	for lat := 0; lat < 6; lat++ {
		phi := math.Pi * float64(lat) / 5
		for lng := 0; lng < 50; lng++ {
			theta := 2 * math.Pi * float64(lng) / 50

			x := radius * math.Sin(phi) * math.Cos(theta+g.angle)
			y := radius * math.Cos(phi)
			z := radius * math.Sin(phi) * math.Sin(theta+g.angle)

			scale := 1.0 / (1.0 + z/300.0)
			if z > 0 { // 手前側のみ描画
				screenX := centerX + x*scale
				screenY := centerY - y*scale
				col := color.RGBA{
					uint8(100 * scale),
					uint8(150 * scale),
					uint8(200 * scale),
					0x80,
				}
				ebitenutil.DrawRect(screen, screenX-1, screenY-1, 2, 2, col)
			}
		}
	}

	ebitenutil.DebugPrint(screen, "3D Rotating Sphere with Text")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World! - 3D Rotating")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
