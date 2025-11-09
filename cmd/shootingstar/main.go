package main

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

type Game struct {
	textX  float64
	textY  float64
	textVX float64
	textVY float64
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
	// 初期値なら右上から左下へ速度をセット
	if g.textVX == 0 && g.textVY == 0 {
		g.textX = screenWidth - 20
		g.textY = 20
		g.textVX = -2.0
		g.textVY = 2.0
	}
	g.textX += g.textVX
	g.textY += g.textVY
	// 画面外に出たらリセット
	if g.textX < -100 || g.textY > screenHeight+50 {
		g.textX = screenWidth - 20
		g.textY = 20
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 画面をクリア
	screen.Fill(color.RGBA{0x20, 0x20, 0x40, 0xff})

	// テキストを描画
	face := &text.GoTextFace{
		Source: faceSource,
		Size:   32,
	}

	op := &text.DrawOptions{}
	op.GeoM.Rotate(-0.785398) // 移動方向に合わせて斜めに傾ける
	op.GeoM.Translate(g.textX, g.textY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "Hello, World!", face, op)
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
