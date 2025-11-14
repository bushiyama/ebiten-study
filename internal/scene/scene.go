package scene

import "github.com/hajimehoshi/ebiten/v2"

// Scene は各シーンが実装すべきインターフェース
type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}
