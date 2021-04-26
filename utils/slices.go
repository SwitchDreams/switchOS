package utils

// Find returns if a value is inside a slice
func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Remove removes a value from slice if exits
func Remove(slice []int, val int) []int {
	var newSlice []int
	for _, sliceValue := range slice {
		if sliceValue != val {
			newSlice = append(newSlice, sliceValue)
		}
	}
	return newSlice
}
