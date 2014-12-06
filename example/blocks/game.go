package blocks

import (
	"github.com/hajimehoshi/ebiten/graphics"
	"github.com/hajimehoshi/ebiten/ui"
	_ "image/png"
)

type Size struct {
	Width  int
	Height int
}

// TODO: Should they be global??
var texturePaths = map[string]string{}
var renderTargetSizes = map[string]Size{}

const ScreenWidth = 256
const ScreenHeight = 240

type GameState struct {
	SceneManager *SceneManager
	Input        *Input
}

type Game struct {
	sceneManager *SceneManager
	input        *Input
	textures     *Textures
}

func NewGame() *Game {
	game := &Game{
		sceneManager: NewSceneManager(NewTitleScene()),
		input:        NewInput(),
	}
	return game
}

func (game *Game) SetTextureFactory(textureFactory graphics.TextureFactory) {
	game.textures = NewTextures(textureFactory)
	for name, path := range texturePaths {
		game.textures.RequestTexture(name, path)
	}
	for name, size := range renderTargetSizes {
		game.textures.RequestRenderTarget(name, size)
	}
}

func (game *Game) isInitialized() bool {
	for name := range texturePaths {
		if !game.textures.Has(name) {
			return false
		}
	}
	for name := range renderTargetSizes {
		if !game.textures.Has(name) {
			return false
		}
	}
	return true
}

func (game *Game) Update(state ui.InputState) {
	if !game.isInitialized() {
		return
	}
	game.input.Update(state)
	game.sceneManager.Update(&GameState{
		SceneManager: game.sceneManager,
		Input:        game.input,
	})
}

func (game *Game) Draw(context graphics.Context) {
	if !game.isInitialized() {
		return
	}
	game.sceneManager.Draw(context, game.textures)
}
