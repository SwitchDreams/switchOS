package algorithms

import (
	"fmt"
	m "github.com/switchdreams/switchOS/memory"
	"github.com/switchdreams/switchOS/utils"
)

// SecondChance implements the FIFO variation
func SecondChance(memory m.Memory) int {
	var framesList []m.FramesList
	var faults int
	for i := 0; i < len(memory.Sequence); i++ {
		page := memory.Sequence[i]
		fmt.Println(framesList)
		if !utils.FindFrameLru(framesList, page) {
			// If memory is not full
			faults += 1
			if i < memory.Size {
				framesList = append(framesList, m.FramesList{Page: page, Arrived: 0})
			} else {
				for _, frame := range framesList {
					if frame.Arrived == 1 {
						framesList = utils.SwapFirstFrame(framesList, m.FramesList{Page: page, Arrived: 0})
					} else {
						framesList[0].Arrived += 1

						// Swaps the first and last elements of a slice
						first := framesList[0]
						framesList = framesList[1:]
						framesList = append(framesList, first)
					}
				}
			}
		} else {
			framesList = utils.UpdateUsedFrame(framesList, page, 0)
		}
	}
	return faults
}
