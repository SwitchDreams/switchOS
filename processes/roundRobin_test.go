package processes

import (
	"reflect"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	processes := []Process{
		{id: 0, arrivalTime: 0, duration: 20},
		{id: 1, arrivalTime: 0, duration: 10},
		{id: 2, arrivalTime: 4, duration: 6},
		{id: 3, arrivalTime: 4, duration: 8},
	}

	processesExecution := []ProcessExecution{
		{pid: 0, startTime: 0, finishTime: 2},
		{pid: 1, startTime: 2, finishTime: 4},
		{pid: 0, startTime: 4, finishTime: 6},
		{pid: 2, startTime: 6, finishTime: 8},
		{pid: 3, startTime: 8, finishTime: 10},
		{pid: 1, startTime: 10, finishTime: 12},
		{pid: 0, startTime: 12, finishTime: 14},
		{pid: 2, startTime: 14, finishTime: 16},
		{pid: 3, startTime: 16, finishTime: 18},
		{pid: 1, startTime: 18, finishTime: 20},
		{pid: 0, startTime: 20, finishTime: 22},
		{pid: 2, startTime: 22, finishTime: 24},
		{pid: 3, startTime: 24, finishTime: 26},
		{pid: 1, startTime: 26, finishTime: 28},
		{pid: 0, startTime: 28, finishTime: 30},
		{pid: 3, startTime: 30, finishTime: 32},
		{pid: 1, startTime: 32, finishTime: 34},
		{pid: 0, startTime: 34, finishTime: 36},
		{pid: 0, startTime: 36, finishTime: 38},
		{pid: 0, startTime: 38, finishTime: 40},
		{pid: 0, startTime: 40, finishTime: 42},
		{pid: 0, startTime: 42, finishTime: 44},
	}
	got, _ := RoundRobin(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRoundRobinDurationNotMod2(t *testing.T) {
	processes := []Process{
		{id: 1, arrivalTime: 0, duration: 6},
		{id: 2, arrivalTime: 0, duration: 3},
		{id: 3, arrivalTime: 0, duration: 1},
		{id: 4, arrivalTime: 0, duration: 7},
	}

	processesExecution := []ProcessExecution{
		{pid: 1, startTime: 0, finishTime: 2},
		{pid: 2, startTime: 2, finishTime: 4},
		{pid: 3, startTime: 4, finishTime: 5},
		{pid: 4, startTime: 5, finishTime: 7},
		{pid: 1, startTime: 7, finishTime: 9},
		{pid: 2, startTime: 9, finishTime: 10},
		{pid: 4, startTime: 10, finishTime: 12},
		{pid: 1, startTime: 12, finishTime: 14},
		{pid: 4, startTime: 14, finishTime: 16},
		{pid: 4, startTime: 16, finishTime: 17},
	}
	got, _ := RoundRobin(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot    %v \nwanted %v", got, want)
	}
}

func TestRoundRobinWithIdleTime(t *testing.T) {
	processes := []Process{
		{id: 1, arrivalTime: 0, duration: 1},
		{id: 2, arrivalTime: 2, duration: 2},
		{id: 3, arrivalTime: 2, duration: 1},
	}

	processesExecution := []ProcessExecution{
		{pid: 1, startTime: 0, finishTime: 1},
		{pid: 2, startTime: 2, finishTime: 4},
		{pid: 3, startTime: 4, finishTime: 5},
	}
	got, _ := RoundRobin(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot    %v \nwanted %v", got, want)
	}
}
