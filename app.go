package main

import (
	"context"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"io/fs"
)

// App struct
type App struct {
	Width  int
	Height int
	Page   string
	ctx    context.Context
}

// 切换到指定页面并调整窗口大小
func (a *App) SwitchPage(page string, width, height int) error {
	a.Page = page
	a.Width = width
	a.Height = height
	return a.restartApp()
}

// 重启应用程序以加载新页面和窗口大小
func (a *App) restartApp() error {
	return wails.Run(&options.App{
		Title:   "My Wails App",
		Width:   a.Width,
		Height:  a.Height,
		Assets:  assets,
		Windows: &windows.Options{},
		Bind:    []interface{}{a},
	})
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
