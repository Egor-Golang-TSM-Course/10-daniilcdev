package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()

	if err != nil {
		log.Fatalln("wrong configuration:", err)
	}

	fmt.Println(config)
}
