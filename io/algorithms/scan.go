package algorithms

import (
	io "github.com/switchdreams/switchOS/io"
	"github.com/switchdreams/switchOS/utils"
)

// Scan implements the SCAN (Elevator) Disk Scheduling Algorithms
// The default initialize is from right to left
func Scan(disk io.Disk) int {
	var countCylinders int
	pendingCylinders := disk.Sequence
	finished := false
	// Scans left
	for i := disk.Init; i > 0; i-- {
		pendingCylinders, finished = checkCylinder(pendingCylinders, i)
		if finished {
			return countCylinders
		}
		countCylinders++
	}

	// Scans right
	for i := 1; i < disk.Size; i++ {
		countCylinders++
		pendingCylinders, finished = checkCylinder(pendingCylinders, i)
		if finished {
			break
		}
	}
	return countCylinders
}

// checkCylinder returns a modified slice from pending cylinder and if finished algorithm
func checkCylinder(slice []int, val int) ([]int, bool) {
	if utils.Find(slice, val) {
		slice = utils.Remove(slice, val)
		if len(slice) == 0 {
			return slice, true
		}
	}
	return slice, false
}
