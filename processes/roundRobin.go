package processes

func RoundRobin(processes []Process) ([]ProcessExecution, error) {
	var processExecutionList []ProcessExecution
	var arrivedProcesses, registeredProcesses, nextArrivedProcesses []Process
	for currentTime := 0; ; {
		arrivedProcesses = getAtMomentProcesses(processes, currentTime)
		registeredProcesses = append(registeredProcesses, arrivedProcesses...)

		nextArrivedProcesses = getAtMomentProcesses(processes, (currentTime + 2))
		if len(nextArrivedProcesses) > 0 {
			registeredProcesses = append(registeredProcesses, nextArrivedProcesses...)
		}

		for _, p := range arrivedProcesses {
			processes = removeProcesses(processes, p.id)
		}

		for _, p := range nextArrivedProcesses {
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
		} else {
			registeredProcesses = firstToLast(registeredProcesses)
		}

		if len(registeredProcesses) == 0 {
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

func firstToLast(arr []Process) []Process {
	curr := arr[0]
	arr = arr[1:]
	return append(arr, curr)
}

func getAtMomentProcesses(processes []Process, currentTime int) []Process {
	var arrivedProcesses []Process
	for _, process := range processes {
		if process.arrivalTime == currentTime {
			arrivedProcesses = append(arrivedProcesses, process)
		}
	}
	return arrivedProcesses
}
