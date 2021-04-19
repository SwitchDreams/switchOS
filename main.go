package main

import (
	"fmt"
	"os"

	p "github.com/switchdreams/switchOS/processes"
	"github.com/switchdreams/switchOS/processes/algorithms"
)

func main() {

	// Read input file
	inputFile := os.Args[1]
	processList, err := p.Parse(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run algorithms
	processExecutionFifo := algorithms.Fifo(processList)
	processExecutionSjf := algorithms.Sjf(processList)
	processExecutionRoundRobin := algorithms.RoundRobin(processList)

	// Get average time
	ttTime, responseTime, waitTime := p.AverageTime(processList, processExecutionFifo, false)
	fmt.Printf("FIFO %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)

	ttTime, responseTime, waitTime = p.AverageTime(processList, processExecutionSjf, false)
	fmt.Printf("SJF %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)

	ttTime, responseTime, waitTime = p.AverageTime(processList, processExecutionRoundRobin, true)
	fmt.Printf("RR %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)

	err = p.ExecutionToFile(processExecutionFifo, "fifo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = p.ExecutionToFile(processExecutionSjf, "sjf.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = p.ExecutionToFile(processExecutionRoundRobin, "rr.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

}
