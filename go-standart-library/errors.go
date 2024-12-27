package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "1" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("1")

	if err == nil {
		fmt.Println("Success")
	}

	switch {
	case errors.Is(err, ValidationError):
		fmt.Println("Validation Error")
	case errors.Is(err, NotFoundError):
		fmt.Println("Not Found")
	default:
		if err != nil {
			fmt.Println("Unknown Error")
		}
	}
}
