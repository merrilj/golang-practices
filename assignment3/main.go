package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	io.Copy(os.Stdout, file)
}