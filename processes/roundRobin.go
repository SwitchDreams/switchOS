package processes

func RoundRobin(processes []Process) ([]ProcessExecution, error) {
	var processExecutionList []ProcessExecution
	var arrivedProcesses, registeredProcesses, nextArrivedProcesses []Process
	var finishTime, decreaseTime int
	for currentTime := 0; ; {

		arrivedProcesses = getAtMomentProcesses(processes, currentTime)
		registeredProcesses = append(registeredProcesses, arrivedProcesses...)

		// nextArrived adiciona na lista final de execução (registeredProcess) os processos do próximo ciclo
		// para ordenar corretamente o Slice
		nextArrivedProcesses = getAtMomentProcesses(processes, (currentTime + 2))
		if len(nextArrivedProcesses) > 0 {
			registeredProcesses = append(registeredProcesses, nextArrivedProcesses...)
		}

		// Ao adicionar os processos no slice de execução (registered), remove da lista inicial de processos
		for _, p := range arrivedProcesses {
			processes = removeProcesses(processes, p.id)
		}

		for _, p := range nextArrivedProcesses {
			processes = removeProcesses(processes, p.id)
		}
		// ------------------------------------------------------------------------------------------------

		currentProcess := registeredProcesses[0]

		if currentProcess.arrivalTime > currentTime {
			// Se o processo atual não está em seu momento de execução -> Idle
			currentTime += 1
			continue
		}

		if currentProcess.duration < 2 {
			finishTime = currentTime + currentProcess.duration
			decreaseTime = currentProcess.duration
		} else {
			finishTime = currentTime + 2
			decreaseTime = 2
		}

		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        currentProcess.id,
			startTime:  currentTime,
			finishTime: finishTime,
		})

		registeredProcesses = decreaseDurationProcess(registeredProcesses, currentProcess.id, decreaseTime)

		if currentProcess.duration <= 2 {
			registeredProcesses = removeProcesses(registeredProcesses, currentProcess.id)
		} else {
			registeredProcesses = firstToLast(registeredProcesses)
		}

		if len(registeredProcesses) == 0 && len(processes) == 0 {
			break
		}

		currentTime = finishTime
	}
	return processExecutionList, nil
}

func decreaseDurationProcess(processes []Process, id int, decreaseTime int) []Process {
	for index, process := range processes {
		if process.id == id {
			processes[index].duration -= decreaseTime
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
