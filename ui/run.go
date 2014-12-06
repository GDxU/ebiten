package ui

import (
	"github.com/hajimehoshi/ebiten/graphics"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Game interface {
	Draw(context graphics.Context)
	Update(inputState InputState)
	SetTextureFactory(textureFactory graphics.TextureFactory)
}

func Run(u UI, game Game, width, height, scale int, title string, fps int) {
	canvas := u.Start(width, height, scale, title)
	game.SetTextureFactory(canvas)

	frameTime := time.Duration(int64(time.Second) / int64(fps))
	tick := time.Tick(frameTime)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	defer u.Terminate()
	for {
		u.DoEvents()
		select {
		default:
			canvas.Draw(game.Draw)
		case <-tick:
			game.Update(canvas.InputState())
			if canvas.IsClosed() {
				return
			}
		case <-sigterm:
			return
		}
	}
}
