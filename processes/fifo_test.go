package processes

import (
	"reflect"
	"testing"
)

func TestFifo(t *testing.T) {
	processes := []Process{
		{arrivalTime: 0, duration: 20}, {arrivalTime: 0, duration: 10},
		{arrivalTime: 4, duration: 6}, {arrivalTime: 4, duration: 8},
	}
	processesExecution := []ProcessExecution{
		{pid: 0, startTime: 0, finishTime: 20},
		{pid: 1, startTime: 20, finishTime: 30},
		{pid: 2, startTime: 30, finishTime: 36},
		{pid: 3, startTime: 36, finishTime: 44},
	}
	got, _ := Fifo(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFifoWithIdleTime(t *testing.T) {
	processes := []Process{
		{arrivalTime: 0, duration: 5}, {arrivalTime: 7, duration: 10},
	}
	processesExecution := []ProcessExecution{
		{pid: 0, startTime: 0, finishTime: 5},
		{pid: 1, startTime: 7, finishTime: 17},
	}
	got, _ := Fifo(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
