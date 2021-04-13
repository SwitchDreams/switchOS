package processes

import (
	"sort"
)

func Sjf(processes []Process) ([]ProcessExecution, error) {

	// Ordenando com base na menor duração, levando em consideração a ordem
	sort.Slice(processes, func(i, j int) bool {
		if processes[i].arrivalTime != processes[j].arrivalTime {
			return processes[i].arrivalTime < processes[j].arrivalTime
		}

		return processes[i].duration < processes[j].duration
	})

	return Fifo(processes)
}
