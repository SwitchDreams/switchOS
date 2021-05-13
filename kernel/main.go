package kernel

import (
	"fmt"

	"github.com/switchdreams/switchOS/errors"
	io "github.com/switchdreams/switchOS/io"
	ioAlgorithms "github.com/switchdreams/switchOS/io/algorithms"
	m "github.com/switchdreams/switchOS/memory"
	mAlgorithms "github.com/switchdreams/switchOS/memory/algorithms"
	p "github.com/switchdreams/switchOS/processes"
	pAlgorithms "github.com/switchdreams/switchOS/processes/algorithms"
)

const (
	PROCESS = 1
	MEMORY  = 2
	IO      = 3
)

func Exec(module int, filename string) errors.IOSError {

	switch module {
	case PROCESS:
		err := handleProcess(filename)
		if err != nil {
			return errors.WrapError(err, "failed to handle processes")
		}
	case MEMORY:
		err := handleMemory(filename)
		if err != nil {
			return errors.WrapError(err, "failed to handle memory")
		}
	case IO:
		err := handleIO(filename)
		if err != nil {
			return errors.WrapError(err, "failed to handle IO")
		}
	default:
		return errors.NewOSError("Opção não reconhecida. As opções disponíveis são 1, 2 ou 3")
	}

	return nil
}

func handleProcess(filename string) errors.IOSError {
	processList, err := p.Parse(filename)
	if err != nil {
		return errors.WrapError(err, fmt.Sprintf("failed to parse %s", filename))
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
		return errors.WrapError(err, "failed to print fifo.txt")
	}

	err = p.ExecutionToFile(processExecutionSjf, "sjf.txt")
	if err != nil {
		return errors.WrapError(err, "failed to print sjf.txt")
	}

	err = p.ExecutionToFile(processExecutionRoundRobin, "rr.txt")
	if err != nil {
		return errors.WrapError(err, "failed to print rr.txt")
	}

	return nil
}

func handleMemory(filename string) errors.IOSError {
	memory, err := m.Parse(filename)
	if err != nil {
		return errors.WrapError(err, fmt.Sprintf("failed to parse %s", filename))
	}
	// Run algorithms
	fifoFaults := mAlgorithms.Fifo(memory)
	secondChanceFaults := mAlgorithms.SecondChance(memory)
	lruFaults := mAlgorithms.Lru(memory)

	fmt.Printf("FIFO %d\n", fifoFaults)
	fmt.Printf("SC %d\n", secondChanceFaults)
	fmt.Printf("LRU %d\n", lruFaults)

	return nil
}

func handleIO(filename string) errors.IOSError {
	disk, err := io.Parse(filename)
	if err != nil {
		return errors.WrapError(err, fmt.Sprintf("failed to parse %s", filename))
	}

	fcfsSeekTime := ioAlgorithms.Fcfs(disk)
	ssfSeekTime := ioAlgorithms.SSF(disk)
	scanSeekTime := ioAlgorithms.Scan(disk)

	fmt.Printf("FCFS %d\n", fcfsSeekTime)
	fmt.Printf("SSF %d\n", ssfSeekTime)
	fmt.Printf("SCAN %d\n", scanSeekTime)

	return nil
}
