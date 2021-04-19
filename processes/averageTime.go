package processes

// AverageTime calculates the average times during executions
func AverageTime(processes []Process, processExecutions []ProcessExecution, isRR bool) (float32, float32, float32) {
	var turnAroundTime int
	var waitTime int
	var finishedPids []int
	lenProcesses := len(processes)
	for i := len(processExecutions) - 1; i >= 0; i-- {
		// If process has finished
		if !find(finishedPids, processExecutions[i].Pid) {
			var turnAroundProcessTime int
			// TurnAround
			turnAroundProcessTime += processExecutions[i].FinishTime
			turnAroundProcessTime -= processes[processExecutions[i].Pid].ArrivalTime

			turnAroundTime += turnAroundProcessTime
			// WaitTime
			waitTime += turnAroundProcessTime - processes[processExecutions[i].Pid].Duration
			finishedPids = append(finishedPids, processExecutions[i].Pid)
		}
	}

	turnAroundAverage := float32(turnAroundTime) / float32(lenProcesses)
	waitTimeAverage := float32(waitTime) / float32(lenProcesses)
	var responseAverage float32
	if isRR {
		responseAverage = float32(Quantum)
	} else {
		responseAverage = waitTimeAverage
	}

	return turnAroundAverage, responseAverage, waitTimeAverage
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
