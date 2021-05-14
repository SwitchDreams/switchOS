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
		Sequence: []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 4, 0, 3, 2},
	}
	got := SecondChance(memory)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
