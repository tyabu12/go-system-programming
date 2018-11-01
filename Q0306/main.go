package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	stream = io.MultiReader(
		io.NewSectionReader(programming, 5, 1), // A
		io.LimitReader(system, 1),              // S
		io.LimitReader(computer, 1),            // C
		io.NewSectionReader(programming, 8, 1), // I
		io.NewSectionReader(programming, 8, 1), // I
	)

	io.Copy(os.Stdout, stream)
}
