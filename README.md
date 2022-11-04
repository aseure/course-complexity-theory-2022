# Course

Available commands:
- `make test`
- `make bench`

# Maze generation / solving

Implement the `BacktrackingSolver()` function in
`maze/solvers/backtracking_solver.go` with a pathfinding algorithm using a
backtracking algorithm. You **only** need to write code inside that file.

To do that, read the existing code and its available datastructures in the
`maze/` directory. If you think one of the maze cells is part of the final path,
mark it with `cell.SetMark(maze.Path)`, this way the generated GIF image which
represents the maze will mark the path in green. To generate a new image and
test your algorithm, simply run `make maze`.
