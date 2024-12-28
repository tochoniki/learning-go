package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("h([a-z])o")
	
	fmt.Println(regex.MatchString("heo"))
	fmt.Println(regex.MatchString("hello"))	
}
