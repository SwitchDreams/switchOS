package processes

func Fifo(processes []Process) ([]ProcessExecution, error) {
	var processExecutionList []ProcessExecution
	var currentTime = 0
	for i, process := range processes {
		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        i,
			startTime:  currentTime,
			finishTime: currentTime + process.duration,
		})
		currentTime += process.duration
	}
	return processExecutionList, nil
}
