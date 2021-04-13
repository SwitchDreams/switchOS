package processes

func Fifo(processes []Process) ([]ProcessExecution, error) {
	var processExecutionList []ProcessExecution
	var currentTime = 0
	var startTime = 0
	for i, process := range processes {
		if process.arrivalTime > currentTime {
			startTime = process.arrivalTime
		} else {
			startTime = currentTime
		}
		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        i,
			startTime:  startTime,
			finishTime: startTime + process.duration,
			duration:   process.duration,
		})
		currentTime += process.duration
	}
	return processExecutionList, nil
}
