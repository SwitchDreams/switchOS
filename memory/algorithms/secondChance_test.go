package algorithms

import (
	"reflect"
	"testing"

	m "github.com/switchdreams/switchOS/memory"
)

func TestSecondChance(t *testing.T) {
	memory := m.Memory{
		Size:     4,
		Sequence: []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5},
	}
	got := SecondChance(memory)
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSecondChanceGeek(t *testing.T) {
	memory := m.Memory{
		Size:     3,
		Sequence: []int{2, 5, 10, 1, 2, 2, 6, 9, 1, 2, 10, 2, 6, 1, 2, 1, 6, 9, 5, 1, 3},
	}
	got := SecondChance(memory)
	want := 13

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
