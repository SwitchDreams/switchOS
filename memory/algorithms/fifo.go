package algorithms

import (
	m "github.com/switchdreams/switchOS/memory"
)

// Fifo implements the First In First Out algorithm
func Fifo(memory m.Memory) int {
	var ram = make([]int, memory.Size)
	var faults int
	var cont int
	for _, page := range memory.Sequence {
		if !find(ram, page) {
			faults += 1
			ram[cont%memory.Size] = page
			cont += 1
		}
	}
	return faults
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
