package processes

import (
	"reflect"
	"testing"
)

func TestCalc(t *testing.T) {
	processes := []Process{
		{id: 0, arrivalTime: 0, duration: 20}, {id: 1, arrivalTime: 0, duration: 10},
		{id: 2, arrivalTime: 4, duration: 6}, {id: 3, arrivalTime: 4, duration: 8},
	}
	processesExecution := []ProcessExecution{
		{pid: 0, startTime: 0, finishTime: 20},
		{pid: 1, startTime: 20, finishTime: 30},
		{pid: 2, startTime: 30, finishTime: 36},
		{pid: 3, startTime: 36, finishTime: 44},
	}
	ttTime, _, waitTime := Calc(processes, processesExecution)
	got := []float32{ttTime, waitTime, waitTime}
	want := []float32{30.5, 19.5, 19.5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
