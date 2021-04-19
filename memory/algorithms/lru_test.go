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
