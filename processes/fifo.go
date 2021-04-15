package processes

func Fifo(processes []Process) []ProcessExecution {
	var processExecutionList []ProcessExecution
	var currentTime int
	var startTime int
	for _, process := range processes {
		if process.arrivalTime > currentTime {
			startTime = process.arrivalTime
		} else {
			startTime = currentTime
		}
		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        process.id,
			startTime:  startTime,
			finishTime: startTime + process.duration,
		})
		currentTime = process.duration + startTime
	}
	return processExecutionList
}
