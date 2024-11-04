package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"io/fs"
	"log"
)

// App struct
type App struct {
	Width  int
	Height int
	Page   string
	ctx    context.Context
}

func (a *App) SwitchPage(page string, width, height int) error {
	a.Page = page
	a.Width = width
	a.Height = height
	return a.restartApp()
}

// 重启应用程序以加载新页面和窗口大小
func (a *App) restartApp() error {
	if a.Page == "" {
		return a.defaultPage()
	}
	fmt.Println(a.Page)
	fmt.Println(a.Width, a.Height)

	return nil
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) GetPageContent(name string) (string, error) {
	var gamesPage = "/frontend/src/games/" + name + "/index.html"
	content, err := fs.ReadFile(assets, gamesPage)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (a *App) defaultPage() error {
	// Create application with options
	err := wails.Run(&options.App{
		Title:             "pf_game_set",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         a.startup,
		OnDomReady:        a.domReady,
		OnBeforeClose:     a.beforeClose,
		OnShutdown:        a.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			a,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "pf_game_set",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
