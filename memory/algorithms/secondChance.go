package algorithms

import (
	m "github.com/switchdreams/switchOS/memory"
	"github.com/switchdreams/switchOS/utils"
)

// SecondChance implements the FIFO variation
func SecondChance(memory m.Memory) int {
	var storage []m.FramesListSecondChance
	faults := 0
	victimPage := 0
	for i, page := range memory.Sequence {
		if index, ok := utils.FindFrame2chance(storage, page); !ok {
			if i < memory.Size {
				storage = append(storage, m.FramesListSecondChance{
					Page: page,
					R:    1,
				})
				faults++
			} else {
				allOnes := true
				for idx := victimPage; idx < len(storage); idx++ {
					if storage[idx].R == 1 {
						storage[idx].R = 0
					} else {
						storage[idx].Page = page
						storage[idx].R = 1

						allOnes = false
						victimPage = idx
						faults++
						break
					}
				}
				if allOnes {
					storage[0] = m.FramesListSecondChance{
						Page: page,
						R:    1,
					}

					faults++
				}
			}
		} else {
			storage[index].R = 1
		}
	}
	return faults
}
