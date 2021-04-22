package algorithms

import (
	"reflect"
	"testing"

	m "github.com/switchdreams/switchOS/memory"
)

func TestFifo(t *testing.T) {
	memory := m.Memory{
		Size:     4,
		Sequence: []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5},
	}
	got := Fifo(memory)
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFifoSize3(t *testing.T) {
	memory := m.Memory{
		Size:     3,
		Sequence: []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5},
	}
	got := Fifo(memory)
	want := 9

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
