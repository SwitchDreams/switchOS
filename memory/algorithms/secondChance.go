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
	for _, page := range memory.Sequence {
		if index, ok := utils.FindFrame2chance(storage, page); !ok {
			if len(storage) < memory.Size {
				storage = append(storage, m.FramesListSecondChance{
					Page: page,
					R:    true,
				})
				faults++
			} else {
				allOnes := true
				for idx := victimPage; idx < len(storage); idx++ {
					if storage[idx].R == true {
						if storage[idx].Count == 3 {
							storage[idx].R = false
						} else {
							storage[idx].Count++
						}
					} else {
						storage[idx].Page = page
						storage[idx].R = true

						allOnes = false
						victimPage = idx
						faults++
						break
					}
				}
				if allOnes {
					storage[0] = m.FramesListSecondChance{
						Page: page,
						R:    true,
					}
					faults++
				}
			}
		} else {
			storage[index].R = true
		}
	}
	return faults
}
