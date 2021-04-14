package processes

type Process struct {
	id          int
	arrivalTime int
	duration    int
}

type ProcessExecution struct {
	pid        int
	startTime  int
	finishTime int
}
