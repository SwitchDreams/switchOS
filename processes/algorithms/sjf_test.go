package algorithms

import (
	"reflect"
	"testing"

	p "github.com/switchdreams/switchOS/processes"
)

func TestSjf(t *testing.T) {
	processes := []p.Process{
		{ID: 0, ArrivalTime: 0, Duration: 20}, {ID: 1, ArrivalTime: 0, Duration: 10},
		{ID: 2, ArrivalTime: 4, Duration: 6}, {ID: 3, ArrivalTime: 4, Duration: 8},
	}
	processesExecution := []p.ProcessExecution{
		{Pid: 1, StartTime: 0, FinishTime: 10},
		{Pid: 2, StartTime: 10, FinishTime: 16},
		{Pid: 3, StartTime: 16, FinishTime: 24},
		{Pid: 0, StartTime: 24, FinishTime: 44},
	}
	got, _ := Sjf(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSjfWithIdleTime(t *testing.T) {
	processes := []p.Process{
		{ID: 0, ArrivalTime: 0, Duration: 5}, {ID: 1, ArrivalTime: 7, Duration: 10},
	}
	processesExecution := []p.ProcessExecution{
		{Pid: 0, StartTime: 0, FinishTime: 5},
		{Pid: 1, StartTime: 7, FinishTime: 17},
	}
	got, _ := Sjf(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
