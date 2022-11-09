package main

import (
	"fmt"
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
	go RenderScreen(screen, style, scoreBoard, ball, leftPaddle, rightPaddle)

	ballTimer := time.NewTicker(time.Millisecond * 100)
	go func() {
		var ballPos component.Coordinates
		slope := component.Coordinates{X: 1, Y: 0}
		if CoinFlip() {
			slope.X *= -1
		}

		var slopeOffsetModifier int
		for {
			<-ballTimer.C

			maxWidth, maxHeight = screen.Size()
			ballPos = ball.CurrentPosition()

			slopeOffsetModifier = -1
			if slope.Y < 0 {
				slopeOffsetModifier = 1
			}

			// right
			if ballPos.X >= maxWidth-2 {
				slope.X *= -1
				switch ballPos.Y - rightPaddle.CurrentPosition().Y {
				case -2:
					slope.Y += 2 * slopeOffsetModifier
				case -1:
					slope.Y += 1 * slopeOffsetModifier
				case 0:
					// don't change the slope
				case 1:
					slope.Y += -1 * slopeOffsetModifier
				case 2:
					slope.Y += -2 * slopeOffsetModifier
				default:
					if incrementScore(leftScore) {
						screen.Fini()
						fmt.Println("Left Player Wins!!!")
						os.Exit(0)
					}
				}

			}
			// left
			if ballPos.X <= 1 {
				slope.X *= -1
				switch ballPos.Y - leftPaddle.CurrentPosition().Y {
				case -2:
					slope.Y += 2 * slopeOffsetModifier
				case -1:
					slope.Y += 1 * slopeOffsetModifier
				case 0:
					// don't change the slope
				case 1:
					slope.Y += -1 * slopeOffsetModifier
				case 2:
					slope.Y += -2 * slopeOffsetModifier
				default:
					if incrementScore(rightScore) {
						screen.Fini()
						fmt.Println("Right Player Wins!!!")
						os.Exit(0)
					}
				}
			}

			// limit how much the slope can be changed
			if slope.Y > 2 {
				slope.Y = 2
			}
			if slope.Y < -2 {
				slope.Y = -2
			}

			// bottom
			if ballPos.Y > maxHeight-2 {
				slope.Y *= -1
			}
			// top
			if ballPos.Y < 6 {
				slope.Y *= -1
			}

			ball.Move(slope)
		}
	}()

	// main thread sits around waiting for input
	var y int
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
					y = -1
					if leftPaddle.CurrentPosition().Y <= 0 {
						y = 0
					}
					leftPaddle.Move(component.Coordinates{X: 0, Y: y})
				case 's':
					y = 1
					if leftPaddle.CurrentPosition().Y >= maxHeight {
						y = 0
					}
					leftPaddle.Move(component.Coordinates{X: 0, Y: y})
				case 'i':
					y = -1
					if rightPaddle.CurrentPosition().Y <= 0 {
						y = 0
					}
					rightPaddle.Move(component.Coordinates{X: 0, Y: y})
				case 'k':
					y = 1
					if rightPaddle.CurrentPosition().Y >= maxHeight {
						y = 0
					}
					rightPaddle.Move(component.Coordinates{X: 0, Y: y})
				}
			}
		}
	}
}

func incrementScore(c *component.Component) bool {
	win := false
	c.Mutex.Lock()
	c.Char++
	if c.Char > '4' {
		win = true
	}
	c.Mutex.Unlock()
	c.Parent.Move(component.Coordinates{X: 0, Y: 0})
	return win
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
