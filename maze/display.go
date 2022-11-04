package maze

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"os/exec"
)

var (
	black = color.RGBA{A: 0xff}
	white = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	green = color.RGBA{R: 0xa2, G: 0xff, B: 0x9d, A: 0xff}
	blue  = color.RGBA{R: 0x64, G: 0x95, B: 0xed, A: 0xff}
	delay = 2
)

func (m *Maze) getColorForPixel(x, y, imgWidth, imgHeight int) color.RGBA {
	wallCellThickness := m.wallThickness + m.cellThickness

	xMod, yMod := x%wallCellThickness, y%wallCellThickness

	isBorder := x < m.wallThickness || imgWidth-m.wallThickness <= x || y < m.wallThickness || imgHeight-m.wallThickness <= y
	isCorner := 0 <= xMod && xMod < m.wallThickness && 0 <= yMod && yMod < m.wallThickness

	column, row := int(math.Floor(float64(x)/float64(wallCellThickness))), int(math.Floor(float64(y)/float64(wallCellThickness)))
	c := m.Cell(column, row)

	isWall := (xMod < m.wallThickness && c.IsWall(Left)) ||
		(xMod > wallCellThickness && c.IsWall(Right)) ||
		(yMod < m.wallThickness && c.IsWall(Top)) ||
		(yMod > wallCellThickness && c.IsWall(Bottom))

	if isBorder || isCorner || isWall {
		return black
	}

	mark := c.GetMark()
	isVisiting := mark == Visiting
	isPath := mark == Path

	if isVisiting {
		return blue
	} else if isPath {
		return green
	}

	return white
}

func (m *Maze) takeScreenshot() {
	imgWidth := (m.wallThickness+m.cellThickness)*m.width + 1*m.wallThickness
	imgHeight := (m.wallThickness+m.cellThickness)*m.height + 1*m.wallThickness

	img := image.NewPaletted(
		image.Rect(0, 0, imgWidth, imgHeight),
		color.Palette{black, green, white, blue},
	)

	for y := 0; y < img.Rect.Max.Y; y++ {
		for x := 0; x < img.Rect.Max.X; x++ {
			img.Set(x, y, m.getColorForPixel(x, y, imgWidth, imgHeight))
		}
	}

	m.img.Image = append(m.img.Image, img)
	m.img.Delay = append(m.img.Delay, delay)
}

func (m *Maze) Display() error {
	if len(m.img.Image) == 0 {
		m.takeScreenshot()
	}

	f, err := os.CreateTemp("", "*.gif")
	if err != nil {
		return fmt.Errorf("cannot create temporary image file: %w", err)
	}

	errEncode := gif.EncodeAll(f, m.img)
	errClose := f.Close()

	if errEncode != nil {
		return fmt.Errorf("could not encode GIF file: %w", errEncode)
	}

	if errClose != nil {
		return fmt.Errorf("could not close GIF file: %w", errClose)
	}

	fmt.Printf("Image file created at %s\n", f.Name())

	if err = exec.Command("open", f.Name()).Run(); err != nil {
		return fmt.Errorf("could not open final GIF file: %w", err)
	}

	return nil
}
