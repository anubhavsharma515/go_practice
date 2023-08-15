package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string
	for _, data := range os.Args[1:] {
		s += sep + data
		sep = " "
	}
	fmt.Println(s)
}
