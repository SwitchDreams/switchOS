package algorithms

import (
	m "github.com/switchdreams/switchOS/memory"
	"github.com/switchdreams/switchOS/utils"
)

// Lru implements the Least Recently Used algorithm
func Lru(memory m.Memory) int {
	var framesList []m.FramesList
	var faults int
	var cont int
	for _, page := range memory.Sequence {
		if !utils.FindFrameLru(framesList, page.Page) {
			faults++
			if cont < memory.Size {
				framesList = append(framesList, m.FramesList{Page: page.Page, Arrived: cont})
			} else {
				framesList = utils.SwapFrame(framesList, m.FramesList{Page: page.Page, Arrived: cont})
			}
			cont++
		} else {
			framesList = utils.UpdateUsedFrame(framesList, page.Page, cont)
		}
	}
	return faults
}
