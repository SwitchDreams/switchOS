package algorithms

import (
	"reflect"
	"testing"

	m "github.com/switchdreams/switchOS/memory"
)

func TestFifo(t *testing.T) {
	memory := m.Memory{
		Size: 4,
		Sequence: []m.MemorySequence{
			{Page: 1},
			{Page: 2},
			{Page: 3},
			{Page: 4},
			{Page: 1},
			{Page: 2},
			{Page: 5},
			{Page: 1},
			{Page: 2},
			{Page: 3},
			{Page: 4},
			{Page: 5},
		},
	}

	got := Fifo(memory)
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFifoSize3(t *testing.T) {
	memory := m.Memory{
		Size: 3,
		Sequence: []m.MemorySequence{
			{Page: 1},
			{Page: 2},
			{Page: 3},
			{Page: 4},
			{Page: 1},
			{Page: 2},
			{Page: 5},
			{Page: 1},
			{Page: 2},
			{Page: 3},
			{Page: 4},
			{Page: 5},
		},
	}

	got := Fifo(memory)
	want := 9

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
