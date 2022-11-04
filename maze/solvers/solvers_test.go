package solvers

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github/aseure/course-complexity-theory-2022/maze"
)

func TestMaze(t *testing.T) {
	width, height := 100, 100

	m := maze.NewMaze(width, height, false)
	m.Generate()
	begin := m.Cell(0, 0)
	end := m.Cell(m.Width()-1, m.Height()-1)

	for _, c := range []struct {
		name   string
		solver Solver
	}{
		{"Backtracking solver", BacktrackingSolver},
	} {
		m.Reset()
		start := time.Now()
		c.solver(begin, end)
		fmt.Printf("%s took %s\n", c.name, time.Since(start))
		require.NoError(t, m.Display())
	}
}
