package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Udin", "George", "Bubu"}

	fmt.Println(slices.Contains(names, "Udin"))
	fmt.Println(slices.Min(names))
}