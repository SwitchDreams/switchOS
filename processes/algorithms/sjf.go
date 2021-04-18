package algorithms

import (
	"sort"

	p "github.com/switchdreams/switchOS/processes"
)

// Sjf implements the Shortest Job First scheduling algorithm
func Sjf(processes []p.Process) ([]p.ProcessExecution, error) {

	var currentTime int
	var arrivedProcesses []p.Process
	var processExecutionList []p.ProcessExecution
	lenProcesses := len(processes)

	for i := 0; i < lenProcesses; i++ {
		arrivedProcesses = getArrivedProcesses(processes, currentTime)
		// Handles idle time
		if len(arrivedProcesses) == 0 {
			sort.Slice(processes, func(i, j int) bool { return processes[0].ArrivalTime < processes[0].ArrivalTime })
			currentTime = processes[0].ArrivalTime
			arrivedProcesses = getArrivedProcesses(processes, currentTime)
		}
		sort.Slice(arrivedProcesses, func(i, j int) bool { return arrivedProcesses[i].Duration < arrivedProcesses[j].Duration })
		processExecutionList = append(processExecutionList, p.ProcessExecution{
			Pid:        arrivedProcesses[0].ID,
			StartTime:  currentTime,
			FinishTime: currentTime + arrivedProcesses[0].Duration,
		})
		currentTime += arrivedProcesses[0].Duration
		processes = removeProcesses(processes, arrivedProcesses[0].ID)
	}

	return processExecutionList, nil
}

func removeProcesses(processes []p.Process, ID int) []p.Process {
	var newProcesses []p.Process
	for _, process := range processes {
		if process.ID != ID {
			newProcesses = append(newProcesses, process)
		}
	}
	return newProcesses
}

func getArrivedProcesses(processes []p.Process, currentTime int) []p.Process {
	var arrivedProcesses []p.Process
	for _, process := range processes {
		if process.ArrivalTime <= currentTime {
			arrivedProcesses = append(arrivedProcesses, process)
		}
	}
	return arrivedProcesses
}

func getSortedArrivedProcesses(processes []p.Process, currentTime int) []p.Process {
	arrivedProcesses := getArrivedProcesses(processes, currentTime)
	sort.Slice(arrivedProcesses, func(i, j int) bool { return arrivedProcesses[i].Duration < arrivedProcesses[j].Duration })
	return arrivedProcesses
}
