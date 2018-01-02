package ui

import "github.com/gotk3/gotk3/gtk"

type SplashPanel struct {
	CommonPanel
	Label *gtk.Label
}

func NewSplashPanel(ui *UI) *SplashPanel {
	m := &SplashPanel{CommonPanel: NewCommonPanel(ui)}
	m.initialize()
	return m
}

func (m *SplashPanel) initialize() {
	logo := MustImageFromFile("octoprint-logo.png")
	m.Label = MustLabel("Connecting to OctoPrint...")

	box := MustBox(gtk.ORIENTATION_VERTICAL, 15)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.SetHExpand(true)

	box.Add(logo)
	box.Add(m.Label)

	m.grid.Attach(box, 1, 0, 1, 1)
}
