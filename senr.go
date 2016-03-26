package main

import (
	"fmt"
	"os"
)

const USAGE = "senr <config file path>"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("No argument.\n")
		fmt.Printf(USAGE)
		os.Exit(1)
	}

	fmt.Println("hello")
}
