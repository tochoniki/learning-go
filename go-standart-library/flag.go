package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("port", 9999, "Server Port")
	flag.Parse()

	fmt.Println("port", *port)
}