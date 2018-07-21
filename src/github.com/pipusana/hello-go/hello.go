package main

import (
	"fmt"

	"github.com/pipusana/gotools"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(gotools.Add(10, 2))
}
