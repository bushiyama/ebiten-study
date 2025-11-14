package main

import (
	"log"

	"github.com/bushiyama/ebiten-study/internal/scene"
	"github.com/bushiyama/ebiten-study/internal/scene/rotate"
	"github.com/bushiyama/ebiten-study/internal/scene/shootingstar"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	scenes       []scene.Scene
	currentScene int
	lastPressed  bool
}

func NewGame() *Game {
	return &Game{
		scenes: []scene.Scene{
			rotate.New(),
			shootingstar.New(),
		},
		currentScene: 0,
	}
}

func (g *Game) Update() error {
	// 'A'キーが押されたら次のシーンに切り替え
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.currentScene = (g.currentScene + 1) % len(g.scenes)
	}

	// 現在のシーンを更新
	return g.scenes[g.currentScene].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 現在のシーンを描画
	g.scenes[g.currentScene].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.scenes[g.currentScene].Layout(outsideWidth, outsideHeight)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Scene Switcher - Press 'A' to switch")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
