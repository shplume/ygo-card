package main

import (
	"fmt"

	"github.com/shplume/ygo-cards/util/config"
)

func main() {
	name := config.Getstring("name")
	fmt.Println("name:", name)
}
