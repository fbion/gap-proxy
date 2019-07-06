package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fanpei91/gap-proxy"
)

type mainUI struct {
	app        fyne.App
	localProxy *gapproxy.LocalProxy
	button     *widget.Button
}

func newMainUI(app fyne.App) *mainUI {
	return &mainUI{
		app: app,
	}
}

func (m *mainUI) loadUI() fyne.Window {
	window := m.app.NewWindow("E-Wall")

	local := widget.NewEntry()
	local.SetPlaceHolder("Local")

	server := widget.NewEntry()
	server.SetPlaceHolder("Server")

	key := widget.NewEntry()
	key.SetPlaceHolder("Secret")

	m.button = widget.NewButton("START", func() {
		m.localProxy = gapproxy.NewLocalProxy(local.Text, server.Text, key.Text)
		m.localProxy.Listen()
		m.button.Disable()
		m.button.SetText("STARTED")
		m.button.Style = widget.DefaultButton
	})
	m.button.Style = widget.PrimaryButton

	top := fyne.NewContainerWithLayout(layout.NewBorderLayout(local, key, nil, nil), local, server, key)

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(top, m.button, nil, nil), top, m.button)

	window.SetContent(content)
	window.Resize(fyne.NewSize(250, 300))
	window.SetFixedSize(true)
	window.CenterOnScreen()

	return window
}

func main() {
	newMainUI(app.New()).loadUI().ShowAndRun()
}
