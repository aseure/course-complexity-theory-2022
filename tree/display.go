package tree

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func display(start func(w func(format string, a ...interface{}))) error {
	var b strings.Builder

	w := func(format string, a ...interface{}) {
		b.WriteString(fmt.Sprintf(format, a...))
	}

	w("digraph bst{\n")
	start(w)
	w("}\n")

	f, err := os.CreateTemp("", "*.svg")
	if err != nil {
		return fmt.Errorf("cannot create temporary dot file %q for writing: %w", f.Name(), err)
	}

	cmd := exec.Command("dot", "-Tsvg")
	cmd.Stdin = strings.NewReader(b.String())
	cmd.Stdout = f

	errCmd := cmd.Run()
	errClose := f.Close()

	if errCmd != nil {
		return fmt.Errorf("could not generate dot file %q correctly: %w", f.Name(), errCmd)
	}

	if errClose != nil {
		return fmt.Errorf("could not close dot file %q correctly: %w", f.Name(), errClose)
	}

	errCmd = exec.Command("open", f.Name()).Run()
	if errCmd != nil {
		return fmt.Errorf("could not open dot file %q correctly: %w", f.Name(), errCmd)
	}

	return nil
}
