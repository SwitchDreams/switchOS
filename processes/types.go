package processes

const Quantum = 2

type Process struct {
	ID          int
	ArrivalTime int
	Duration    int
}

type ProcessExecution struct {
	Pid        int
	StartTime  int
	FinishTime int
}
