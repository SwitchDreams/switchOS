package algorithms

import (
	"sort"

	p "github.com/switchdreams/switchOS/processes"
	"github.com/switchdreams/switchOS/utils"
)

// Sjf implements the Shortest Job First scheduling algorithm
func Sjf(processes []p.Process) []p.ProcessExecution {

	var currentTime int
	var arrivedProcesses []p.Process
	var processExecutionList []p.ProcessExecution
	lenProcesses := len(processes)

	for i := 0; i < lenProcesses; i++ {
		arrivedProcesses = utils.GetArrivedProcesses(processes, currentTime)
		// Handles idle time
		if len(arrivedProcesses) == 0 {
			sort.Slice(processes, func(i, j int) bool { return processes[0].ArrivalTime < processes[0].ArrivalTime })
			currentTime = processes[0].ArrivalTime
			arrivedProcesses = utils.GetArrivedProcesses(processes, currentTime)
		}
		sort.Slice(arrivedProcesses, func(i, j int) bool { return arrivedProcesses[i].Duration < arrivedProcesses[j].Duration })
		processExecutionList = append(processExecutionList, p.ProcessExecution{
			Pid:        arrivedProcesses[0].ID,
			StartTime:  currentTime,
			FinishTime: currentTime + arrivedProcesses[0].Duration,
		})
		currentTime += arrivedProcesses[0].Duration
		processes = utils.RemoveProcesses(processes, arrivedProcesses[0].ID)
	}

	return processExecutionList
}
