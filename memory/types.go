package memory

// Memory represents a memory
type Memory struct {
	Size     int
	Sequence []int
}

type FramesList struct {
	Page    int
	Arrived int
}

type FramesListSecondChance struct {
	Page  int
	R     bool
	Count int
}
