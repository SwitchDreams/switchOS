package algorithms

import (
	"reflect"
	"testing"

	io "github.com/switchdreams/switchOS/io"
)

func TestSSF(t *testing.T) {
	disk := io.Disk{
		Size:     199,
		Init:     53,
		Sequence: []int{98, 183, 37, 122, 14, 124, 65, 67},
	}
	got := SSF(disk)
	want := 236

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSSFGeek(t *testing.T) {
	disk := io.Disk{
		Size:     199,
		Init:     50,
		Sequence: []int{82, 170, 43, 140, 24, 16, 190},
	}
	got := SSF(disk)
	want := 208

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
