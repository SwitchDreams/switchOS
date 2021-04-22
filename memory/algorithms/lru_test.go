package algorithms

import (
	"reflect"
	"testing"

	m "github.com/switchdreams/switchOS/memory"
)

func TestLru(t *testing.T) {
	memory := m.Memory{
		Size:     4,
		Sequence: []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5},
	}
	got := Lru(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test2Lru(t *testing.T) {
	memory := m.Memory{
		Size:     4,
		Sequence: []int{2, 3, 1, 5, 7, 3, 1, 1, 8, 9, 1, 5},
	}
	got := Lru(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test3Lru(t *testing.T) {
	memory := m.Memory{
		Size:     3,
		Sequence: []int{1, 3, 5, 2, 3, 7, 1, 3, 8, 4},
	}
	got := Lru(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test4Lru(t *testing.T) {
	memory := m.Memory{
		Size:     3,
		Sequence: []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2},
	}
	got := Lru(memory)
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
