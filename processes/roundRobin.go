package processes

func RoundRobin(processes []Process) ([]ProcessExecution, error) {
	var processExecutionList []ProcessExecution
	for currentTime := 0; ; {
		arrivedProcesses := getSortedArrivedProcesses(processes, currentTime)
		currentProcess := arrivedProcesses[0]

		if currentProcess.duration <= 2 {
			removeProcesses(processes, currentProcess.id)
		}

		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        arrivedProcesses[0].id,
			startTime:  currentTime,
			finishTime: currentTime + 2,
		})

		decreaseDurationProcess(processes, currentProcess.id)
		if len(processes) == 0 {
			break
		}

		currentTime += 2

	}
	return processExecutionList, nil
}

func decreaseDurationProcess(processes []Process, id int) []Process {
	for index, process := range processes {
		if process.id == id {
			processes[index].duration -= 2
		}
	}
	return processes
}
