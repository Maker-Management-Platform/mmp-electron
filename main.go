package main

import (
	"embed"

	stlib "github.com/eduardooliveira/stLib/core"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go stlib.Run()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mmp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
