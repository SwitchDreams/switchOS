package algorithms

import (
	"sort"

	io "github.com/switchdreams/switchOS/io"
	"github.com/switchdreams/switchOS/utils"
)

func SSF(disk io.Disk) int {

	aux := disk.Sequence
	head := disk.Init
	seek := 0
	for range disk.Sequence {
		closest := getClosest(aux, head)

		aux = utils.Remove(aux, closest)

		seek += abs(closest - head)
		head = closest
	}

	return seek
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

func getClosest(list []int, target int) int {
	distance := abs(list[0] - target)
	idx := 0
	for i := 1; i < len(list); i++ {
		dst := abs(list[i] - target)
		if dst < distance {
			idx = i
			distance = dst
		}
	}

	return list[idx]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
