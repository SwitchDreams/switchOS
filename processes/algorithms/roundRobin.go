package algorithms

import (
	p "github.com/switchdreams/switchOS/processes"
	"github.com/switchdreams/switchOS/utils"
)

// RoundRobin implements scheduling algorithm that uses the quantum time
// to organize processes
func RoundRobin(processes []p.Process) []p.ProcessExecution {
	var processExecutionList []p.ProcessExecution
	var arrivedProcesses, registeredProcesses, nextArrivedProcesses []p.Process
	var FinishTime, decreaseTime int
	for currentTime := 0; ; {

		arrivedProcesses = utils.GetAtMomentProcesses(processes, currentTime)
		registeredProcesses = append(registeredProcesses, arrivedProcesses...)

		// nextArrived adiciona na lista final de execução (registeredProcess) os processos do próximo ciclo
		// para ordenar corretamente o Slice
		nextArrivedProcesses = utils.GetAtMomentProcesses(processes, currentTime+p.Quantum)
		if len(nextArrivedProcesses) > 0 {
			registeredProcesses = append(registeredProcesses, nextArrivedProcesses...)
		}

		// Ao adicionar os processos no slice de execução (registered), remove da lista inicial de processos
		for _, p := range arrivedProcesses {
			processes = utils.RemoveProcesses(processes, p.ID)
		}

		for _, p := range nextArrivedProcesses {
			processes = utils.RemoveProcesses(processes, p.ID)
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

		registeredProcesses = utils.DecreaseDurationProcess(registeredProcesses, currentProcess.ID, decreaseTime)

		if currentProcess.Duration <= p.Quantum {
			registeredProcesses = utils.RemoveProcesses(registeredProcesses, currentProcess.ID)
		} else {
			registeredProcesses = utils.FirstToLast(registeredProcesses)
		}

		if len(registeredProcesses) == 0 && len(processes) == 0 {
			break
		}

		currentTime = FinishTime
	}
	return processExecutionList
}
