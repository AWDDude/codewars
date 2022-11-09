package component

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

func New(x, y int, char rune, style tcell.Style, children ...*Component) *Component {
	c := &Component{
		position: Coordinates{X: x, Y: y},
		style:    style,
		Char:     char,
		Children: children,
		ch:       make(chan Coordinates),
	}

	for i := range c.Children {
		c.Children[i].Parent = c
	}

	return c
}

type Component struct {
	mutex    sync.RWMutex
	Char     rune
	position Coordinates
	ch       chan Coordinates
	style    tcell.Style
	Parent   *Component
	Children []*Component
}

func (c *Component) CurrentPosition() Coordinates {
	c.mutex.RLock()
	pos := c.position
	c.mutex.RUnlock()
	if c.Parent != nil {
		pPos := c.Parent.CurrentPosition()
		pos.X += pPos.X
		pos.Y += pPos.Y
	}
	return pos
}

func (c *Component) Move(mov Coordinates) {
	c.ch <- mov
}

func (c *Component) Clear(screen tcell.Screen) {
	pos := c.CurrentPosition()

	if c.Char != 0 {
		screen.SetContent(pos.X, pos.Y, ' ', nil, c.style)
	}

	for _, child := range c.Children {
		child.Clear(screen)
	}
}

func (c *Component) Write(screen tcell.Screen) {
	pos := c.CurrentPosition()

	if c.Char != 0 {
		screen.SetContent(pos.X, pos.Y, c.Char, nil, c.style)
	}

	for _, child := range c.Children {
		child.Write(screen)
	}
}

func (c *Component) Render(screen tcell.Screen) {
	select {
	case newPos := <-c.ch:
		// clear old
		c.Clear(screen)
		// update position
		c.mutex.Lock()
		c.position.X += newPos.X
		c.position.Y += newPos.Y
		c.mutex.Unlock()
		// write new
		c.Write(screen)

	default:
		// default prevents the thread from being blocked while waiting for input from the channel
	}
}

type Coordinates struct {
	X, Y int
}
