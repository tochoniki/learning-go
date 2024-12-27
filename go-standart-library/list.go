package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Hello World")
	data.PushBack("Hello World 2")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}