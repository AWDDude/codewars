package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	foregroundColor := tcell.ColorReset
	backgroundColor := tcell.ColorReset

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := screen.Init(); err != nil {
		log.Fatal(err.Error())
	}
	style := tcell.StyleDefault.Background(backgroundColor).Foreground(foregroundColor)
	screen.SetStyle(style)

	ball := Assembly{
		X:     1,
		Y:     1,
		Parts: []Part{{X: 0, Y: 0, Char: '*'}},
	}

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEsc:
				screen.Fini()
				os.Exit(0)
			case tcell.KeyCtrlC:
				screen.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				ball.X++
			case tcell.KeyLeft:
				ball.X--
			case tcell.KeyDown:
				ball.Y++
			case tcell.KeyUp:
				ball.Y--
			}
		}

		ball.Render(screen, style)

		screen.Show()
	}
}

type Assembly struct {
	X, Y, oldX, oldY int
	Parts            []Part
}

func (a *Assembly) Render(screen tcell.Screen, style tcell.Style) {
	// remove old
	for i := range a.Parts {
		screen.SetContent(a.oldX+a.Parts[i].X, a.oldY+a.Parts[i].Y, ' ', nil, style)
	}
	// add new
	for i := range a.Parts {
		screen.SetContent(a.X+a.Parts[i].X, a.Y+a.Parts[i].Y, a.Parts[i].Char, nil, style)
		a.oldX, a.oldY = a.X, a.Y
	}
}

type Part struct {
	X, Y int
	Char rune
}
