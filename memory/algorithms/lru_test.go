package algorithms

import (
	"reflect"
	"testing"

	m "github.com/switchdreams/switchOS/memory"
)

func TestLru(t *testing.T) {
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

	got := Lru(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test2Lru(t *testing.T) {
	memory := m.Memory{
		Size: 4,
		Sequence: []m.MemorySequence{
			{Page: 2},
			{Page: 3},
			{Page: 1},
			{Page: 5},
			{Page: 7},
			{Page: 3},
			{Page: 1},
			{Page: 1},
			{Page: 8},
			{Page: 9},
			{Page: 1},
			{Page: 5},
		},
	}

	got := Lru(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test3Lru(t *testing.T) {
	memory := m.Memory{
		Size: 3,
		Sequence: []m.MemorySequence{
			{Page: 1},
			{Page: 3},
			{Page: 5},
			{Page: 2},
			{Page: 3},
			{Page: 7},
			{Page: 1},
			{Page: 3},
			{Page: 8},
			{Page: 4},
		},
	}

	got := Lru(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test4Lru(t *testing.T) {
	memory := m.Memory{
		Size: 3,
		Sequence: []m.MemorySequence{
			{Page: 7},
			{Page: 0},
			{Page: 1},
			{Page: 2},
			{Page: 0},
			{Page: 3},
			{Page: 0},
			{Page: 4},
			{Page: 2},
			{Page: 3},
			{Page: 0},
			{Page: 3},
			{Page: 2},
			{Page: 1},
			{Page: 2},
		},
	}

	got := Lru(memory)
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
