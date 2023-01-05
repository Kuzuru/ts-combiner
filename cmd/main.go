package main

import (
	"log"

	"github.com/Kuzuru/ts-combiner/internal/app"
)

func main() {
	err := app.Main()
	if err != nil {
		log.Fatalln("[ERR] Main.go: ", err)
	}
}
