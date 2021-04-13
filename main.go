package main

import (
	"fmt"

	"github.com/switchdreams/switchOS/processes"
)

func main() {
	processList, _ := processes.Parse("./file.txt")
	processExecutionSjf, _ := processes.Sjf(processList)
	processExecutionFifo, _ := processes.Fifo(processList)
	fmt.Println(processList)
	fmt.Println(processExecutionSjf)
	fmt.Println(processExecutionFifo)
}
