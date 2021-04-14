package processes

import "sort"

func Sjf(processes []Process) ([]ProcessExecution, error) {

	var currentTime int
	var arrivedProcesses []Process
	var processExecutionList []ProcessExecution
	lenProcesses := len(processes)

	for i := 0; i < lenProcesses; i++ {
		arrivedProcesses = getArrivedProcesses(processes, currentTime)
		// Handles idle time
		if len(arrivedProcesses) == 0 {
			sort.Slice(processes, func(i, j int) bool { return processes[0].arrivalTime < processes[0].arrivalTime })
			currentTime = processes[0].arrivalTime
			arrivedProcesses = getArrivedProcesses(processes, currentTime)
		}
		sort.Slice(arrivedProcesses, func(i, j int) bool { return arrivedProcesses[i].duration < arrivedProcesses[j].duration })
		processExecutionList = append(processExecutionList, ProcessExecution{
			pid:        arrivedProcesses[0].id,
			startTime:  currentTime,
			finishTime: currentTime + arrivedProcesses[0].duration,
		})
		currentTime += arrivedProcesses[0].duration
		processes = removeProcesses(processes, arrivedProcesses[0].id)
	}

	return processExecutionList, nil
}

func removeProcesses(processes []Process, id int) []Process {
	var newProcesses []Process
	for _, process := range processes {
		if process.id != id {
			newProcesses = append(newProcesses, process)
		}
	}
	return newProcesses
}

func getArrivedProcesses(processes []Process, currentTime int) []Process {
	var arrivedProcesses []Process
	for _, process := range processes {
		if process.arrivalTime <= currentTime {
			arrivedProcesses = append(arrivedProcesses, process)
		}
	}
	return arrivedProcesses
}

func getSortedArrivedProcesses(processes []Process, currentTime int) []Process {
	arrivedProcesses := getArrivedProcesses(processes, currentTime)
	sort.Slice(arrivedProcesses, func(i, j int) bool { return arrivedProcesses[i].duration < arrivedProcesses[j].duration })
	return arrivedProcesses
}
