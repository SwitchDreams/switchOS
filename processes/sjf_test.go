package processes

import (
	"reflect"
	"testing"
)

func TestSjf(t *testing.T) {
	processes := []Process{
		{id: 0, arrivalTime: 0, duration: 20}, {id: 1, arrivalTime: 0, duration: 10},
		{id: 2, arrivalTime: 4, duration: 6}, {id: 3, arrivalTime: 4, duration: 8},
	}
	processesExecution := []ProcessExecution{
		{pid: 1, startTime: 0, finishTime: 10},
		{pid: 2, startTime: 10, finishTime: 16},
		{pid: 3, startTime: 16, finishTime: 24},
		{pid: 0, startTime: 24, finishTime: 44},
	}
	got := Sjf(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSjfWithIdleTime(t *testing.T) {
	processes := []Process{
		{id: 0, arrivalTime: 0, duration: 5}, {id: 1, arrivalTime: 7, duration: 10},
	}
	processesExecution := []ProcessExecution{
		{pid: 0, startTime: 0, finishTime: 5},
		{pid: 1, startTime: 7, finishTime: 17},
	}
	got := Sjf(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
