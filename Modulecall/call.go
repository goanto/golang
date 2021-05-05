package main

import (
	"fmt"

	"example.com/module"
)

func main() {
	message := module.Hello("Geo")
	fmt.Println(message)
}
