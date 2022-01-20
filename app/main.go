// build js, window

package main

import (
	log "github.com/go-pkgz/lgr"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/proj2048/app/resources"
	"github.com/pencroff/proj2048/app/scene"
	"github.com/sedyh/mizu/pkg/engine"
)

func main() {
	initResources()
	runGame()
}

func initResources() {
	resources.InitColors()
	resources.InitFonts()
	resources.InitTiles()
}

func runGame() {
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(500, 800)
	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Printf("[Fatal] Can't run game: %v", err)
	}
}
