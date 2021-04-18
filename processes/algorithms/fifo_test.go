package algorithms

import (
	"reflect"
	"testing"

	p "github.com/switchdreams/switchOS/processes"
)

func TestFifo(t *testing.T) {
	processes := []p.Process{
		{ID: 0, ArrivalTime: 0, Duration: 20}, {ID: 1, ArrivalTime: 0, Duration: 10},
		{ID: 2, ArrivalTime: 4, Duration: 6}, {ID: 3, ArrivalTime: 4, Duration: 8},
	}
	processesExecution := []p.ProcessExecution{
		{Pid: 0, StartTime: 0, FinishTime: 20},
		{Pid: 1, StartTime: 20, FinishTime: 30},
		{Pid: 2, StartTime: 30, FinishTime: 36},
		{Pid: 3, StartTime: 36, FinishTime: 44},
	}
	got, _ := Fifo(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFifoWithIdleTime(t *testing.T) {
	processes := []p.Process{
		{ID: 0, ArrivalTime: 0, Duration: 5}, {ID: 1, ArrivalTime: 7, Duration: 10},
	}
	processesExecution := []p.ProcessExecution{
		{Pid: 0, StartTime: 0, FinishTime: 5},
		{Pid: 1, StartTime: 7, FinishTime: 17},
	}
	got, _ := Fifo(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
