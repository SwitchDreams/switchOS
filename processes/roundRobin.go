package processes

func RoundRobin(processes []Process) ([]ProcessExecution, error) {
	var processExecutionList []ProcessExecution
	var arrivedProcesses, registeredProcesses []Process
	for currentTime := 0; ; {
		arrivedProcesses = getArrivedProcesses(processes, currentTime)
		registeredProcesses = append(registeredProcesses, arrivedProcesses...)

		for _, p := range arrivedProcesses {
			processes = removeProcesses(processes, p.id)
		}

		currentProcess := registeredProcesses[0]

		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        currentProcess.id,
			startTime:  currentTime,
			finishTime: currentTime + 2,
		})

		registeredProcesses = decreaseDurationProcess(registeredProcesses, currentProcess.id)

		if currentProcess.duration <= 2 {
			registeredProcesses = removeProcesses(registeredProcesses, currentProcess.id)
		}

		if len(registeredProcesses) == 0 {
			break
		}

		registeredProcesses = firstToLast(registeredProcesses)

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

func firstToLast(arr []Process) []Process {
	curr := arr[0]
	arr = arr[1:]
	return append(arr, curr)
}
