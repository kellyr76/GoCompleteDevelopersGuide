package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	//fmt.Println(os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, f)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote bytes: ", len(bs))

	return len(bs), nil
}
