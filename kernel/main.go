package kernel

import (
	"fmt"
	m "github.com/switchdreams/switchOS/memory"
	mAlgorithms "github.com/switchdreams/switchOS/memory/algorithms"
	p "github.com/switchdreams/switchOS/processes"
	pAlgorithms "github.com/switchdreams/switchOS/processes/algorithms"
)

const PROCESS = 1
const MEMORY = 2
const IO = 3

func Exec(module int, fileName string) {

	switch module {
	case PROCESS:
		handleProcess(fileName)
	case MEMORY:
		handleMemory(fileName)
	case IO:
	}

}

func handleProcess(fileName string) {
	processList, err := p.Parse(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run algorithms
	processExecutionFifo := pAlgorithms.Fifo(processList)
	processExecutionSjf := pAlgorithms.Sjf(processList)
	processExecutionRoundRobin := pAlgorithms.RoundRobin(processList)

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

func handleMemory(fileName string) {
	memory, err := m.Parse(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Run algorithms
	fifoFaults := mAlgorithms.Fifo(memory)
	secondChanceFaults := mAlgorithms.SecondChance(memory)
	lruFaults := mAlgorithms.Lru(memory)

	fmt.Printf("FIFO %d\n", fifoFaults)
	fmt.Printf("SC %d\n", secondChanceFaults)
	fmt.Printf("LRU %d\n", lruFaults)
}
