package main

import (
	"fmt"

	"github.com/switchdreams/switchOS/processes"
)

func main() {
	processList, _ := processes.Parse("./file.txt")
	// processExecutionFifo, _ := processes.Fifo(processList)
	// processExecutionSjf, _ := processes.Sjf(processList)
	// fmt.Println(processList)
	// fmt.Println(processExecutionSjf)
	// fmt.Println(processExecutionFifo)

	ret, _ := processes.RoundRobin(processList)
	fmt.Println(ret)
}
