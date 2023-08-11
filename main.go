package main

import (
	"flag"
	"fmt"
	"os"
)

var test = flag.Bool("test", true, "test flag")

func main() {
	flag.Parse()
	fmt.Println("Hello! This is myjenkins!")
	if !*test {
		fmt.Println("test failed!")
		os.Exit(1)
	}
}
