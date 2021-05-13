package algorithms

import (
	"sort"

	io "github.com/switchdreams/switchOS/io"
)

func SSF(disk io.Disk) int {

	separatedSequence := separateList(disk.Sequence, disk.Init)

	seekTime := abs(separatedSequence[0] - disk.Init)
	for i := 0; i < len(separatedSequence)-1; i++ {
		seekTime += abs(separatedSequence[i] - separatedSequence[i+1])
	}

	return seekTime
}

func separateList(list []int, value int) []int {
	var lessThan []int
	var greaterThan []int

	for _, elem := range list {
		if elem > value {
			greaterThan = append(greaterThan, elem)
		} else if elem < value {
			lessThan = append(lessThan, elem)
		}
	}

	sort.Slice(lessThan, func(i, j int) bool {
		return lessThan[i] < lessThan[j]
	})
	sort.Ints(greaterThan)

	return append(lessThan, greaterThan...)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
