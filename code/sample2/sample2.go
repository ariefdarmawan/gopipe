package main

import (
	"flag"
	"fmt"
)

var (
	fname = flag.String("name", "", "-name=clustername")
)

func main() {
	flag.Parse()

	fmt.Printf("Hello %s", *fname)
}
