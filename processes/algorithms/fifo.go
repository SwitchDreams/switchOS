package algorithms

import p "github.com/switchdreams/switchOS/processes"

// Fifo implements the First In First Out scheduling algorithm
func Fifo(processes []p.Process) []p.ProcessExecution {
	var processExecutionList []p.ProcessExecution
	var currentTime int
	var StartTime int
	for _, process := range processes {
		if process.ArrivalTime > currentTime {
			StartTime = process.ArrivalTime
		} else {
			StartTime = currentTime
		}
		processExecutionList = append(processExecutionList, p.ProcessExecution{
			Pid:        process.ID,
			StartTime:  StartTime,
			FinishTime: StartTime + process.Duration,
		})
		currentTime = process.Duration + StartTime
	}
	return processExecutionList
}
