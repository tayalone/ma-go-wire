package main

import (
	"fmt"

	Config "github.com/tayalone/ma-go-wire/config"
)

func main() {
	fmt.Println("Start ma go wire")
	config := Config.Initialize()
	fmt.Println("config.Port", config.Port)
}
