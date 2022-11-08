package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := screen.Init(); err != nil {
		log.Fatal(err.Error())
	}
	screen.SetStyle(style)
	ball := NewAssembly(1, 1, NewPart(0, 0, '*'))

	// goroutine for rendering screen
	go RenderScreen(screen, style, ball)

	// main thread sits around waiting for input
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
				ball.Move(1, 0)
			case tcell.KeyLeft:
				ball.Move(-1, 0)
			case tcell.KeyDown:
				ball.Move(0, 1)
			case tcell.KeyUp:
				ball.Move(0, -1)
			}
		}
	}
}

func RenderScreen(screen tcell.Screen, style tcell.Style, objects ...*Assembly) {
	for {
		for i := range objects {
			objects[i].Render(screen, style)
		}
		screen.Show()
	}
}
