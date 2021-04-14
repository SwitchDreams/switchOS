package processes

import "sort"

func Sjf(processes []Process) ([]ProcessExecution, error) {

	var currentTime int
	var currentTimeProcess []Process
	var processExecution []ProcessExecution
	tamanho := len(processes)

	for i := 0; i < tamanho; i++ {
		currentTimeProcess = getCurrentTimeProcess(processes, currentTime)
		sort.Slice(currentTimeProcess, func(i, j int) bool { return currentTimeProcess[i].duration < currentTimeProcess[j].duration })
		processExecution = append(processExecution, ProcessExecution{
			pid:        currentTimeProcess[0].id,
			startTime:  currentTime,
			finishTime: currentTime + currentTimeProcess[0].duration,
		})
		currentTime += currentTimeProcess[0].duration
		processes = removeCurrentTimeProcess(processes, currentTimeProcess[0].id)
	}

	return processExecution, nil
}

func removeCurrentTimeProcess(processes []Process, id int) []Process {
	var newProcesses []Process
	for _, process := range processes {
		if process.id != id {
			newProcesses = append(newProcesses, process)
		}
	}
	return newProcesses
}

func getCurrentTimeProcess(processes []Process, currentTime int) []Process {
	var currentTimeProcess []Process
	for _, process := range processes {
		if process.arrivalTime <= currentTime {
			currentTimeProcess = append(currentTimeProcess, process)
		}
	}
	return currentTimeProcess
}
