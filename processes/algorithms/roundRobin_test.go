package algorithms

import (
	"reflect"
	"testing"

	p "github.com/switchdreams/switchOS/processes"
)

func TestRoundRobin(t *testing.T) {
	processes := []p.Process{
		{ID: 0, ArrivalTime: 0, Duration: 20},
		{ID: 1, ArrivalTime: 0, Duration: 10},
		{ID: 2, ArrivalTime: 4, Duration: 6},
		{ID: 3, ArrivalTime: 4, Duration: 8},
	}

	processesExecution := []p.ProcessExecution{
		{Pid: 0, StartTime: 0, FinishTime: 2},
		{Pid: 1, StartTime: 2, FinishTime: 4},
		{Pid: 0, StartTime: 4, FinishTime: 6},
		{Pid: 2, StartTime: 6, FinishTime: 8},
		{Pid: 3, StartTime: 8, FinishTime: 10},
		{Pid: 1, StartTime: 10, FinishTime: 12},
		{Pid: 0, StartTime: 12, FinishTime: 14},
		{Pid: 2, StartTime: 14, FinishTime: 16},
		{Pid: 3, StartTime: 16, FinishTime: 18},
		{Pid: 1, StartTime: 18, FinishTime: 20},
		{Pid: 0, StartTime: 20, FinishTime: 22},
		{Pid: 2, StartTime: 22, FinishTime: 24},
		{Pid: 3, StartTime: 24, FinishTime: 26},
		{Pid: 1, StartTime: 26, FinishTime: 28},
		{Pid: 0, StartTime: 28, FinishTime: 30},
		{Pid: 3, StartTime: 30, FinishTime: 32},
		{Pid: 1, StartTime: 32, FinishTime: 34},
		{Pid: 0, StartTime: 34, FinishTime: 36},
		{Pid: 0, StartTime: 36, FinishTime: 38},
		{Pid: 0, StartTime: 38, FinishTime: 40},
		{Pid: 0, StartTime: 40, FinishTime: 42},
		{Pid: 0, StartTime: 42, FinishTime: 44},
	}
	got, _ := RoundRobin(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRoundRobinDurationNotMod2(t *testing.T) {
	processes := []p.Process{
		{ID: 1, ArrivalTime: 0, Duration: 6},
		{ID: 2, ArrivalTime: 0, Duration: 3},
		{ID: 3, ArrivalTime: 0, Duration: 1},
		{ID: 4, ArrivalTime: 0, Duration: 7},
	}

	processesExecution := []p.ProcessExecution{
		{Pid: 1, StartTime: 0, FinishTime: 2},
		{Pid: 2, StartTime: 2, FinishTime: 4},
		{Pid: 3, StartTime: 4, FinishTime: 5},
		{Pid: 4, StartTime: 5, FinishTime: 7},
		{Pid: 1, StartTime: 7, FinishTime: 9},
		{Pid: 2, StartTime: 9, FinishTime: 10},
		{Pid: 4, StartTime: 10, FinishTime: 12},
		{Pid: 1, StartTime: 12, FinishTime: 14},
		{Pid: 4, StartTime: 14, FinishTime: 16},
		{Pid: 4, StartTime: 16, FinishTime: 17},
	}
	got, _ := RoundRobin(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot    %v \nwanted %v", got, want)
	}
}

func TestRoundRobinWithIdleTime(t *testing.T) {
	processes := []p.Process{
		{ID: 1, ArrivalTime: 0, Duration: 1},
		{ID: 2, ArrivalTime: 2, Duration: 2},
		{ID: 3, ArrivalTime: 2, Duration: 1},
	}

	processesExecution := []p.ProcessExecution{
		{Pid: 1, StartTime: 0, FinishTime: 1},
		{Pid: 2, StartTime: 2, FinishTime: 4},
		{Pid: 3, StartTime: 4, FinishTime: 5},
	}
	got, _ := RoundRobin(processes)
	want := processesExecution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot    %v \nwanted %v", got, want)
	}
}
