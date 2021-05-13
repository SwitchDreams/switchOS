package algorithms

import (
	m "github.com/switchdreams/switchOS/memory"
	"github.com/switchdreams/switchOS/utils"
)

// SecondChance implements the FIFO variation
func SecondChance(memory m.Memory) int {
	ram := make([]int, memory.Size)
	var faults, cont int

	for _, page := range memory.Sequence {
		if !utils.Find(ram, page.Page) {
			if !page.R {
				page.R = true
			} else {
				faults++
				ram[cont%memory.Size] = page.Page
				cont++
			}
		}
	}
	return faults
}
