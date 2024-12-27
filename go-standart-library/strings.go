package main

import (
	"fmt"
	"strings"
)

func main() {
	contains := strings.Contains("Hello World", "ello")
	split := strings.Split("Hello World", " ")
	upperCase := strings.ToUpper("hello world")
	lowercase := strings.ToLower("HELLO WORLD")
	trim := strings.Trim("   Hello World   ", " ")
	replaceAll := strings.ReplaceAll("Hello World", "Hello", "Hi")

	fmt.Println(contains)
	fmt.Println(split)
	fmt.Println(split[0])
	fmt.Println(upperCase)
	fmt.Println(lowercase)
	fmt.Println(trim)
	fmt.Println(replaceAll)
}