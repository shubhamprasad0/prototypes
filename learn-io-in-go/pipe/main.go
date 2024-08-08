package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprintf(w, "writing to the pipe\n")
		w.Close()
	}()

	io.Copy(os.Stdout, r)
}
