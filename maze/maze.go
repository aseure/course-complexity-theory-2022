package maze

import (
	"image/gif"
	"math/rand"
)

type Maze struct {
	width         int
	height        int
	wallThickness int
	cellThickness int
	cells         [][]*Cell
	img           *gif.GIF
}

func NewMaze(width, height int, generateAnimation bool) *Maze {
	rand.Seed(int64(width * height))

	m := &Maze{
		width:         width,
		height:        height,
		wallThickness: 4,
		cellThickness: 40,
		cells:         make([][]*Cell, height),
		img:           new(gif.GIF),
	}

	screenshotFn := m.takeScreenshot
	if !generateAnimation {
		screenshotFn = func() {}
	}

	for row := 0; row < height; row++ {
		var rowCells []*Cell
		for column := 0; column < width; column++ {
			rowCells = append(rowCells, newCell(screenshotFn))
		}
		m.cells[row] = rowCells
	}

	for row := 0; row < m.height; row++ {
		for column := 0; column < m.width; column++ {
			c, right, bottom := m.Cell(column, row), m.Cell(column+1, row), m.Cell(column, row+1)
			c.connect(Right, right)
			c.connect(Bottom, bottom)
		}
	}

	m.Reset()

	return m
}

func (m *Maze) Width() int {
	return m.width
}

func (m *Maze) Height() int {
	return m.height
}

func (m *Maze) Reset() {
	for row := 0; row < m.height; row++ {
		for column := 0; column < m.width; column++ {
			m.Cell(column, row).SetMark(Blank)
		}
	}
	m.img = new(gif.GIF)
}

func (m *Maze) Cell(column, row int) *Cell {
	if column < 0 || m.width <= column ||
		row < 0 || m.height <= row {
		return nil
	}
	return m.cells[row][column]
}
