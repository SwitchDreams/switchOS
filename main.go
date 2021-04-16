package main

import (
	"fmt"
	"github.com/switchdreams/switchOS/processes"
)

func main() {
	processList, _ := processes.Parse("./file.txt")
	processExecutionFifo, _ := processes.Fifo(processList)
	processExecutionSjf, _ := processes.Sjf(processList)
	processExecutionRoundRobin, _ := processes.Sjf(processList)
	ttTime, responseTime, waitTime := processes.AverageTime(processList, processExecutionFifo, false)
	fmt.Printf("FIFO %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)
	ttTime, responseTime, waitTime = processes.AverageTime(processList, processExecutionSjf, false)
	fmt.Printf("SJF %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)
	ttTime, responseTime, waitTime = processes.AverageTime(processList, processExecutionRoundRobin, true)
	fmt.Printf("RR %.1f %.1f %.1f\n", ttTime, responseTime, waitTime)
	processes.ExecutionToFile(processExecutionFifo, "fifo.txt")
	processes.ExecutionToFile(processExecutionSjf, "sjf.txt")
	processes.ExecutionToFile(processExecutionRoundRobin, "rr.txt")

}
