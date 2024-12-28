package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
}