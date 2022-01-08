package main

import (
	"fmt"
	"io"
	"os"
)

type logFile struct{}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lw := logFile{}
	io.Copy(lw, file)
}

func (logFile) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
