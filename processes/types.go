package processes

// Quantum is the time slice that the process can use to execute
const Quantum = 2

// Process represents a process
type Process struct {
	ID          int
	ArrivalTime int
	Duration    int
}

// ProcessExecution represents a process in execution
type ProcessExecution struct {
	Pid        int
	StartTime  int
	FinishTime int
}
