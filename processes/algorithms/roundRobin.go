package algorithms

import p "github.com/switchdreams/switchOS/processes"

// RoundRobin implements scheduling algorithm that uses the quantum time
// to organize processes
func RoundRobin(processes []p.Process) ([]p.ProcessExecution, error) {
	var processExecutionList []p.ProcessExecution
	var arrivedProcesses, registeredProcesses, nextArrivedProcesses []p.Process
	var FinishTime, decreaseTime int
	for currentTime := 0; ; {

		arrivedProcesses = getAtMomentProcesses(processes, currentTime)
		registeredProcesses = append(registeredProcesses, arrivedProcesses...)

		// nextArrived adiciona na lista final de execução (registeredProcess) os processos do próximo ciclo
		// para ordenar corretamente o Slice
		nextArrivedProcesses = getAtMomentProcesses(processes, currentTime+p.Quantum)
		if len(nextArrivedProcesses) > 0 {
			registeredProcesses = append(registeredProcesses, nextArrivedProcesses...)
		}

		// Ao adicionar os processos no slice de execução (registered), remove da lista inicial de processos
		for _, p := range arrivedProcesses {
			processes = removeProcesses(processes, p.ID)
		}

		for _, p := range nextArrivedProcesses {
			processes = removeProcesses(processes, p.ID)
		}
		// ------------------------------------------------------------------------------------------------

		currentProcess := registeredProcesses[0]

		if currentProcess.ArrivalTime > currentTime {
			// Se o processo atual não está em seu momento de execução -> Idle
			currentTime++
			continue
		}

		if currentProcess.Duration < p.Quantum {
			FinishTime = currentTime + currentProcess.Duration
			decreaseTime = currentProcess.Duration
		} else {
			FinishTime = currentTime + p.Quantum
			decreaseTime = p.Quantum
		}

		processExecutionList = append(processExecutionList, p.ProcessExecution{
			Pid:        currentProcess.ID,
			StartTime:  currentTime,
			FinishTime: FinishTime,
		})

		registeredProcesses = decreaseDurationProcess(registeredProcesses, currentProcess.ID, decreaseTime)

		if currentProcess.Duration <= p.Quantum {
			registeredProcesses = removeProcesses(registeredProcesses, currentProcess.ID)
		} else {
			registeredProcesses = firstToLast(registeredProcesses)
		}

		if len(registeredProcesses) == 0 && len(processes) == 0 {
			break
		}

		currentTime = FinishTime
	}
	return processExecutionList, nil
}

func decreaseDurationProcess(processes []p.Process, ID int, decreaseTime int) []p.Process {
	for index, process := range processes {
		if process.ID == ID {
			processes[index].Duration -= decreaseTime
		}
	}
	return processes
}

func firstToLast(arr []p.Process) []p.Process {
	curr := arr[0]
	arr = arr[1:]
	return append(arr, curr)
}

func getAtMomentProcesses(processes []p.Process, currentTime int) []p.Process {
	var arrivedProcesses []p.Process
	for _, process := range processes {
		if process.ArrivalTime == currentTime {
			arrivedProcesses = append(arrivedProcesses, process)
		}
	}
	return arrivedProcesses
}
