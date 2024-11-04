package main

import (
	"embed"
)

//go:embed frontend/src
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()
	app.SwitchPage("", 1024, 768)

}
