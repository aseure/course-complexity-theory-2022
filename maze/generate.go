package maze

func (m *Maze) Generate() {
	generateMaze(make(map[*Cell]bool), m.cells[0][0])
}

func generateMaze(visited map[*Cell]bool, c *Cell) {
	visited[c] = true

	for d, neighbor := range c.GetRandomizedNeighbors() {
		if neighbor != nil && !visited[neighbor] {
			c.removeWall(d)
			generateMaze(visited, neighbor)
		}
	}
}
