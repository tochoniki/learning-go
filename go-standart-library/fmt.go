package main

import "fmt"

func main() {
	firstName := "Cho"
	lastName := "Niki"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello %s %s", firstName, lastName)
}