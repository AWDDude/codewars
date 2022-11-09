package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"pong/component"

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

	maxWidth, maxHeight := screen.Size()
	screen.SetStyle(style)
	leftScore := component.New(-2, 0, '0', style)
	rightScore := component.New(2, 0, '0', style)
	scoreBoard := component.New(maxWidth/2, 2, 0, style,
		rightScore,
		leftScore,
	)
	ball := component.New(maxWidth/2, maxHeight/2, '*', style)
	leftPaddle := component.New(0, maxHeight/2, 0, style,
		component.New(0, -2, '|', style),
		component.New(0, -1, '|', style),
		component.New(0, 0, '|', style),
		component.New(0, 1, '|', style),
		component.New(0, 2, '|', style),
	)
	rightPaddle := component.New(0, maxHeight/2, 0, style,
		component.New(maxWidth-1, -2, '|', style),
		component.New(maxWidth-1, -1, '|', style),
		component.New(maxWidth-1, 0, '|', style),
		component.New(maxWidth-1, 1, '|', style),
		component.New(maxWidth-1, 2, '|', style),
	)

	// goroutine for rendering screen
	go RenderScreen(screen, style, ball, leftPaddle, rightPaddle)

	ballTimer := time.NewTicker(time.Millisecond * 100)
	go func() {
		var ballPos component.Coordinates
		slope := component.Coordinates{X: 1, Y: RandNum(2) + 1}
		if CoinFlip() {
			slope.X *= -1
		}
		if CoinFlip() {
			slope.Y *= -1
		}
		for {
			<-ballTimer.C

			maxWidth, maxHeight = screen.Size()
			ballPos = ball.CurrentPosition()

			// right
			if ballPos.X > maxWidth-4 {
				slope.X *= -1
				switch ballPos.X - rightPaddle.CurrentPosition().X {
				case -2:
					slope.Y += 2
				case -1:
					slope.Y += 1
				case 0:
					// don't change the slope
				case 1:
					slope.Y += -1
				case -2:
					slope.Y += -1
				default:

				}

			}
			// left
			if ballPos.X < 3 {
				slope.X *= -1
			}

			// bottom
			if ballPos.Y > maxHeight-2 {
				slope.Y *= -1
			}
			// top
			if ballPos.Y < 1 {
				slope.Y *= -1
			}

			ball.Move(slope)
		}
	}()

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
			case tcell.KeyRune:
				switch event.Rune() {
				case 'w':
					leftPaddle.Move(component.Coordinates{X: 0, Y: -1})
				case 's':
					leftPaddle.Move(component.Coordinates{X: 0, Y: 1})
				case 'i':
					rightPaddle.Move(component.Coordinates{X: 0, Y: -1})
				case 'k':
					rightPaddle.Move(component.Coordinates{X: 0, Y: 1})
				}
			}
		}
	}
}

func RandNum(n int) int {
	rand.Seed(time.Now().UnixMicro())
	return rand.Intn(n + 1)
}
func CoinFlip() bool {
	return RandNum(1) == 1
}

func RenderScreen(screen tcell.Screen, style tcell.Style, objects ...*component.Component) {
	for _, v := range objects {
		v.Write(screen)
	}
	for {
		for _, v := range objects {
			v.Render(screen)
		}
		screen.Show()
	}
}
