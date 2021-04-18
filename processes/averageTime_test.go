package processes

import (
	"reflect"
	"testing"
)

func TestCalcFifo(t *testing.T) {
	processes := []Process{
		{ID: 0, ArrivalTime: 0, Duration: 20}, {ID: 1, ArrivalTime: 0, Duration: 10},
		{ID: 2, ArrivalTime: 4, Duration: 6}, {ID: 3, ArrivalTime: 4, Duration: 8},
	}
	processesExecution := []ProcessExecution{
		{Pid: 0, StartTime: 0, FinishTime: 20},
		{Pid: 1, StartTime: 20, FinishTime: 30},
		{Pid: 2, StartTime: 30, FinishTime: 36},
		{Pid: 3, StartTime: 36, FinishTime: 44},
	}
	ttTime, responseTime, waitTime := AverageTime(processes, processesExecution, false)
	got := []float32{ttTime, responseTime, waitTime}
	want := []float32{30.5, 19.5, 19.5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalcSjf(t *testing.T) {
	processes := []Process{
		{ID: 0, ArrivalTime: 0, Duration: 20}, {ID: 1, ArrivalTime: 0, Duration: 10},
		{ID: 2, ArrivalTime: 4, Duration: 6}, {ID: 3, ArrivalTime: 4, Duration: 8},
	}
	processesExecution := []ProcessExecution{
		{Pid: 1, StartTime: 0, FinishTime: 10},
		{Pid: 2, StartTime: 10, FinishTime: 16},
		{Pid: 3, StartTime: 16, FinishTime: 24},
		{Pid: 0, StartTime: 24, FinishTime: 44},
	}
	ttTime, responseTime, waitTime := AverageTime(processes, processesExecution, false)
	got := []float32{ttTime, responseTime, waitTime}
	want := []float32{21.5, 10.5, 10.5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalcRoundRobin(t *testing.T) {
	processes := []Process{
		{ID: 0, ArrivalTime: 0, Duration: 20},
		{ID: 1, ArrivalTime: 0, Duration: 10},
		{ID: 2, ArrivalTime: 4, Duration: 6},
		{ID: 3, ArrivalTime: 4, Duration: 8},
	}

	processesExecution := []ProcessExecution{
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
	ttTime, responseTime, waitTime := AverageTime(processes, processesExecution, true)
	got := []float32{ttTime, responseTime, waitTime}
	want := []float32{31.5, 2.0, 20.5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
