package main

import (
	"fmt"

	"github.com/switchdreams/switchOS/processes"
)

func main() {
	processList, err := processes.Parse("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	processExecutionFifo := processes.Fifo(processList)

	processExecutionSjf := processes.Sjf(processList)

	fmt.Println(processList)
	fmt.Println(processExecutionSjf)
	fmt.Println(processExecutionFifo)
}
