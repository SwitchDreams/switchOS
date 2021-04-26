package utils

import (
	"sort"

	m "github.com/switchdreams/switchOS/memory"
	p "github.com/switchdreams/switchOS/processes"
)

// DecreaseDurationProcess decreses the duration of a process
func DecreaseDurationProcess(processes []p.Process, ID int, decreaseTime int) []p.Process {
	for index, process := range processes {
		if process.ID == ID {
			processes[index].Duration -= decreaseTime
		}
	}
	return processes
}

// FirstToLast swaps the first and last elements of a slice, implementing a round queue
func FirstToLast(arr []p.Process) []p.Process {
	curr := arr[0]
	arr = arr[1:]
	return append(arr, curr)
}

// GetAtMomentProcesses gets all processes that arrived in the specified time
func GetAtMomentProcesses(processes []p.Process, currentTime int) []p.Process {
	var arrivedProcesses []p.Process
	for _, process := range processes {
		if process.ArrivalTime == currentTime {
			arrivedProcesses = append(arrivedProcesses, process)
		}
	}
	return arrivedProcesses
}

// RemoveProcesses removes a speficic process based on its id
func RemoveProcesses(processes []p.Process, ID int) []p.Process {
	var newProcesses []p.Process
	for _, process := range processes {
		if process.ID != ID {
			newProcesses = append(newProcesses, process)
		}
	}
	return newProcesses
}

// GetArrivedProcesses gets gets all processes that arrived in the specified time
func GetArrivedProcesses(processes []p.Process, currentTime int) []p.Process {
	var arrivedProcesses []p.Process
	for _, process := range processes {
		if process.ArrivalTime <= currentTime {
			arrivedProcesses = append(arrivedProcesses, process)
		}
	}
	return arrivedProcesses
}

// GetSortedArrivedProcesses gets gets all processes that arrived in the specified time and sorts it
func GetSortedArrivedProcesses(processes []p.Process, currentTime int) []p.Process {
	arrivedProcesses := GetArrivedProcesses(processes, currentTime)
	sort.Slice(arrivedProcesses, func(i, j int) bool { return arrivedProcesses[i].Duration < arrivedProcesses[j].Duration })
	return arrivedProcesses
}

// Find Frame by Page - LRU
func FindFrameLru(slice []m.FramesList, val int) bool {
	for _, item := range slice {
		if item.Page == val {
			return true
		}
	}
	return false
}

// RemoveProcesses removes a speficic process based on its id
func RemoveFrame(frames []m.FramesList, val int) []m.FramesList {
	var newFrames []m.FramesList
	for _, frame := range frames {
		if frame.Page != val {
			newFrames = append(newFrames, frame)
		}
	}
	return newFrames
}

// Return frame with smallest arrived
func MinFrame(frames []m.FramesList) m.FramesList {
	min := frames[0]
	for _, frame := range frames {
		if min.Arrived > frame.Arrived {
			min = frame
		}
	}
	return min
}

// Remove the smallest frame and add the new
func SwapFrame(frames []m.FramesList, newFrame m.FramesList) []m.FramesList {
	min := MinFrame(frames)
	frames = RemoveFrame(frames, min.Page)
	return append(frames, newFrame)
}

// When the frame is in use, we update the arrived attribute
func UpdateUsedFrame(frames []m.FramesList, page int, cont int) []m.FramesList {
	for idx, frame := range frames {
		if frame.Page == page {
			frames[idx].Arrived = cont
		}
	}
	return frames
}
