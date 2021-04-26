package main

import (
	"os"
	"strconv"

	"github.com/switchdreams/switchOS/kernel"
)

func main() {
	module := os.Args[1]
	inputFile := os.Args[2]
	moduleInt, _ := strconv.Atoi(module)
	kernel.Exec(moduleInt, inputFile)
}
