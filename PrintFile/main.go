package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fmt.Println(os.Args)
	f, _ := os.Open(os.Args[1])
	b := make([]byte,10)
	_, err := f.Read(b)
	if err == nil {
		fmt.Println(string(b))
	}

	io.Copy(os.Stdout, f)
}
