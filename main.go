package main

import (
	"fmt"

	"github.com/switchdreams/switchOS/processes"
)

func main() {
	processList, _ := processes.Parse("./file.txt")
	processExecutionFifo, _ := processes.Fifo(processList)
	processExecutionSjf, _ := processes.Sjf(processList)
	processes.ExecutionToFile(processExecutionFifo, "fifo.txt")
	processes.ExecutionToFile(processExecutionFifo, "sjf.txt")
	fmt.Println(processList)
	fmt.Println(processExecutionSjf)
	fmt.Println(processExecutionFifo)
}
