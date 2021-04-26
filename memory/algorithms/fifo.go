package algorithms

import (
	m "github.com/switchdreams/switchOS/memory"
	"github.com/switchdreams/switchOS/utils"
)

// Fifo implements the First In First Out algorithm
func Fifo(memory m.Memory) int {
	ram := make([]int, memory.Size)
	var faults int
	var cont int
	for _, page := range memory.Sequence {
		if !utils.Find(ram, page.Page) {
			faults++
			ram[cont%memory.Size] = page.Page
			cont++
		}
	}
	return faults
}
