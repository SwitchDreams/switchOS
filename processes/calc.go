package processes

func Calc(processes []Process, processExecutions []ProcessExecution) (float32, float32, float32) {
	var turnAroundTime int
	var waitTime int
	var finishedPids []int
	lenProcesses := len(processes)
	for i := len(processExecutions) - 1; i >= 0; i-- {
		// If process has finished
		if !Find(finishedPids, processExecutions[i].pid) {
			var turnAroundProcessTime int
			// TurnAround
			turnAroundProcessTime += processExecutions[i].finishTime
			turnAroundProcessTime -= processes[processExecutions[i].pid].arrivalTime

			turnAroundTime += turnAroundProcessTime
			// WaitTime
			waitTime += turnAroundProcessTime - processes[processExecutions[i].pid].duration
			finishedPids = append(finishedPids, processExecutions[i].pid)
		}
	}

	return float32(turnAroundTime) / float32(lenProcesses),
		float32(waitTime) / float32(lenProcesses),
		float32(waitTime) / float32(lenProcesses)
}

func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
