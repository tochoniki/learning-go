package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Hello World"))

	fmt.Println(encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)

	fmt.Println(string(decoded))
}
