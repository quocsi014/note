package main

import (
	"fmt"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
