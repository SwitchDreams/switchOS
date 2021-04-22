package utils

import (
	m "github.com/switchdreams/switchOS/memory"
)

// Find RAM frame
func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// FindFrameLru Find Frame by Page - LRU
func FindFrameLru(slice []m.FramesList, val int) bool {
	for _, item := range slice {
		if item.Page == val {
			return true
		}
	}
	return false
}

// RemoveFrame removes a speficic process based on its id
func RemoveFrame(frames []m.FramesList, val int) []m.FramesList {
	var newFrames []m.FramesList
	for _, frame := range frames {
		if frame.Page != val {
			newFrames = append(newFrames, frame)
		}
	}
	return newFrames
}

// MinFrame return frame with smallest arrived
func MinFrame(frames []m.FramesList) m.FramesList {
	min := frames[0]
	for _, frame := range frames {
		if min.Arrived > frame.Arrived {
			min = frame
		}
	}
	return min
}

// SwapFrame removes the smallest frame and add the new
func SwapFrame(frames []m.FramesList, newFrame m.FramesList) []m.FramesList {
	min := MinFrame(frames)
	frames = RemoveFrame(frames, min.Page)
	return append(frames, newFrame)
}

// UpdateUsedFrame when the frame is in use, we update the arrived attribute
func UpdateUsedFrame(frames []m.FramesList, page int, cont int) []m.FramesList {
	for idx, frame := range frames {
		if frame.Page == page {
			frames[idx].Arrived = cont
		}
	}
	return frames
}
