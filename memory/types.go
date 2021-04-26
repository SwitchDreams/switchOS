package memory

type MemorySequence struct {
	Page int
	R    bool
}

// Memory represents a memory
type Memory struct {
	Size     int
	Sequence []MemorySequence
}

type FramesList struct {
	Page    int
	Arrived int
}
