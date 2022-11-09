package main

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

func NewAssembly(x, y int, parts ...*Part) *Assembly {
	return &Assembly{
		Loc: Coordinates{
			X: x,
			Y: y,
		},
		Parts: parts,
		ch:    make(chan Coordinates),
	}
}

type Assembly struct {
	Loc   Coordinates
	Parts []*Part
	ch    chan Coordinates
}

func (a *Assembly) Render(screen tcell.Screen, style tcell.Style) {
	select {
	case newLoc := <-a.ch:
		// remove old
		for i := range a.Parts {
			screen.SetContent(a.Loc.X+a.Parts[i].Loc.X,
				a.Loc.Y+a.Parts[i].Loc.Y, ' ', nil, style)
		}
		a.Loc.X += newLoc.X
		a.Loc.Y += newLoc.Y
		// add new
		for i := range a.Parts {
			screen.SetContent(a.Loc.X+a.Parts[i].Loc.X,
				a.Loc.Y+a.Parts[i].Loc.Y, a.Parts[i].Char, nil, style)
		}
	default:
		// default prevents the thread from being blocked while waiting for input from the channel
	}
}

func (a *Assembly) Move(x, y int) {
	a.ch <- Coordinates{X: x, Y: y}
}

func NewPart(x, y int, char rune) *Part {
	return &Part{
		Loc:  Coordinates{X: x, Y: y},
		Char: char,
	}
}

type Part struct {
	Loc  Coordinates
	Char rune
}

type Coordinates struct {
	X, Y int
}

type Component struct {
	mutex    sync.RWMutex
	position Coordinates
	char     rune
	style    tcell.Style
	Children []*Component
}

func (c *Component) CurrentPosition() Coordinates {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.position
}

func (c *Component) Move(x, y int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.position.X += x
	c.position.Y += y
}

func (c *Component) Render(screen tcell.Screen) {
	c.mutex.Lock()
	case newLoc := <-a.ch:
		// remove old
		for i := range a.Parts {
			screen.SetContent(a.Loc.X+a.Parts[i].Loc.X,
				a.Loc.Y+a.Parts[i].Loc.Y, ' ', nil, style)
		}
		a.Loc.X += newLoc.X
		a.Loc.Y += newLoc.Y
		// add new
		for i := range a.Parts {
			screen.SetContent(a.Loc.X+a.Parts[i].Loc.X,
				a.Loc.Y+a.Parts[i].Loc.Y, a.Parts[i].Char, nil, style)
		}
	default:
		// default prevents the thread from being blocked while waiting for input from the channel
	}
}
