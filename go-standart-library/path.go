package main

import (
	"fmt"
	"path"
)

func main() {
	pathname := "/go/main.go"
	fmt.Println(path.Dir(pathname))
	fmt.Println(path.Base(pathname))
}