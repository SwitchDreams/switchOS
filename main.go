package main

import (
	"fmt"

	"github.com/switchdreams/switchOS/processes"
)

func main() {
	processList, _ := processes.Parse("./file.txt")
	fmt.Println(processList)
}
