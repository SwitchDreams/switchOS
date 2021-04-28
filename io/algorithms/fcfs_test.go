package algorithms

import (
	"reflect"
	"testing"

	io "github.com/switchdreams/switchOS/io"
)

func TestFcfs(t *testing.T) {
	disk := io.Disk{
		Size:     199,
		Init:     53,
		Sequence: []int{98, 183, 37, 122, 14, 124, 65, 67},
	}
	got := Fcfs(disk)
	want := 640

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
