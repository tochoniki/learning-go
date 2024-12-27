package main

import (
	"fmt"
	"strconv"
)

func main() {
	parsedResult, err := strconv.ParseBool("true")

	if (err != nil) {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parsedResult)
	}

	parsedInt, err := strconv.Atoi("10")
	if (err != nil) {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parsedInt)
	}
}