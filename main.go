package main

import (
	"fmt"
	"os"

	p "github.com/switchdreams/switchOS/processes"
	"github.com/switchdreams/switchOS/processes/algorithms"
)

func main() {

	inputFile := os.Args[1]
	processList, _ := p.Parse(inputFile)

	processExecutionFifo, _ := algorithms.Fifo(processList)
	processExecutionSjf, _ := algorithms.Sjf(processList)
	processExecutionRoundRobin, _ := algorithms.Sjf(processList)
	ttTime, responseTime, waitTime := p.AverageTime(processList, processExecutionFifo, false)
	fmt.Printf("FIFO %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)
	ttTime, responseTime, waitTime = p.AverageTime(processList, processExecutionSjf, false)
	fmt.Printf("SJF %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)
	ttTime, responseTime, waitTime = p.AverageTime(processList, processExecutionRoundRobin, true)
	fmt.Printf("RR %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)
	p.ExecutionToFile(processExecutionFifo, "fifo.txt")
	p.ExecutionToFile(processExecutionSjf, "sjf.txt")
	p.ExecutionToFile(processExecutionRoundRobin, "rr.txt")

}
