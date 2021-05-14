package algorithms

import (
	io "github.com/switchdreams/switchOS/io"
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	disk := io.Disk{
		Size:     199,
		Init:     53,
		Sequence: []int{98, 183, 37, 122, 14, 124, 65, 67},
	}
	got := Scan(disk)
	want := 236

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestScanGeek(t *testing.T) {
	disk := io.Disk{
		Size:     200,
		Init:     50,
		Sequence: []int{176, 79, 34, 60, 92, 11, 41, 114},
	}
	got := Scan(disk)
	want := 226

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestScanOnlyLeft(t *testing.T) {
	disk := io.Disk{
		Size:     199,
		Init:     53,
		Sequence: []int{40, 30},
	}
	got := Scan(disk)
	want := 23

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestScanOnlyRight(t *testing.T) {
	disk := io.Disk{
		Size:     199,
		Init:     53,
		Sequence: []int{60},
	}
	got := Scan(disk)
	want := 111

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
