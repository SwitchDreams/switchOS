package processes

import (
	"sort"
)

func Sjf(processes []Process) ([]ProcessExecution, error) {

	// Ordenando com base na menor duração
	sort.Slice(processes, func(i, j int) bool { return processes[i].duration < processes[j].duration })

	return Fifo(processes)
}
