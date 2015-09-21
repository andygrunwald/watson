package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	Name    = "Watson"
	Version = "0.0.1"
)

func main() {
	var (
		version = flag.Bool("version", false, "Prints the version and exits")
	)
	flag.Parse()

	if *version {
		fmt.Printf("%s v%s\n", Name, Version)
		os.Exit(0)
	}

	fmt.Println("Hi, i am Watson. Nice to meet you.")
}
