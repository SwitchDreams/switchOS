package algorithms

import (
	"reflect"
	"testing"

	m "github.com/switchdreams/switchOS/memory"
)

func TestAleteia(t *testing.T) {
	memory := m.Memory{
		Size:     6,
		Sequence: []int{1, 3, 5, 9, 3, 9, 5, 7, 8, 1, 10, 1, 2, 1, 3, 9, 10, 7},
	}

	got := SecondChance(memory)
	want := 11

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
