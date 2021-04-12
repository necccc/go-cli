package main

import (
	"flag"
	"fmt"
)

var output string

func init() {
	flag.StringVar(&output, "output", "json", "Output format, default is 'json'")
	flag.StringVar(&output, "o", "json", "Output format, default is 'json' (shorthand)")
}

func main() {
	flag.Parse()

	fmt.Println("output has value ", output)
}
