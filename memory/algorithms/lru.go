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
		if !utils.FindFrameLru(framesList, page) {
			faults += 1
			if cont < memory.Size {
				framesList = append(framesList, m.FramesList{Page: page, Arrived: cont})
			} else {
				framesList = utils.SwapFrame(framesList, m.FramesList{Page: page, Arrived: cont})
			}
			cont += 1
		} else {
			framesList = utils.UpdateUsedFrame(framesList, page, cont)
		}
	}
	return faults
}
