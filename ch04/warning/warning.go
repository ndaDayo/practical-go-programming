package warning

import (
	"fmt"
	"os"
)

type Warning interface {
	Show(message string)
}

type ConsoleWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s", os.Args[0], message)
}
