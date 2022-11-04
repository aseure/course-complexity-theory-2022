package maze

import "math/rand"

type Cell struct {
	neighbors      [4]*Cell
	walls          [4]bool
	mark           Mark
	screenshotFunc func()
}

func newCell(screenshotFunc func()) *Cell {
	return &Cell{
		neighbors:      [4]*Cell{},
		walls:          [4]bool{},
		mark:           0,
		screenshotFunc: screenshotFunc,
	}
}

func (c *Cell) connect(d Direction, neighbor *Cell) {
	if neighbor != nil {
		c.neighbors[d] = neighbor
		c.walls[d] = true

		neighbor.neighbors[d^1] = c
		neighbor.walls[d^1] = true
	}
}

func (c *Cell) removeWall(d Direction) {
	if n := c.neighbors[d]; n != nil {
		c.walls[d] = false
		n.walls[d^1] = false
	}
}

func (c *Cell) IsWall(d Direction) bool {
	return c == nil || map[Direction]bool{
		Top:    c.neighbors[0] == nil || c.walls[0],
		Bottom: c.neighbors[1] == nil || c.walls[1],
		Left:   c.neighbors[2] == nil || c.walls[2],
		Right:  c.neighbors[3] == nil || c.walls[3],
	}[d]
}

func (c *Cell) GetRandomizedNeighbors() map[Direction]*Cell {
	neighbors := make(map[Direction]*Cell)

	firstDirection := rand.Intn(4)
	for i := 0; i < 4; i++ {
		d := Direction((firstDirection + i) % 4)
		neighbors[d] = c.neighbors[d]
	}

	return neighbors
}

func (c *Cell) GetMark() Mark {
	return c.mark
}

func (c *Cell) SetMark(m Mark) {
	c.mark = m
	c.screenshotFunc()
}
