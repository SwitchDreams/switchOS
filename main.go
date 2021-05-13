package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/switchdreams/switchOS/kernel"
)

func main() {
	module := os.Args[1]
	inputFile := os.Args[2]
	moduleInt, _ := strconv.Atoi(module)
	err := kernel.Exec(moduleInt, inputFile)
	if err != nil {
		fmt.Printf("ERRO: %s\n", err.Error())
	}
}
