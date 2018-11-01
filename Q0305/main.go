package main

import (
	"io"
)

func copyN(dest io.Writer, src io.Reader, length int) (written int64, err error) {
	reader := io.LimitReader(src, int64(length))
	written, err = io.Copy(dest, reader)
	if written < int64(length) && err == nil {
		err = io.EOF
	}
	return
}

func main() {

}
