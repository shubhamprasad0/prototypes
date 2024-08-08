package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello world\n")
	io.Copy(os.Stdout, r)
}
