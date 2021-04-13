package processes

type Process struct {
	arrivalTime int
	duration    int
}

type ProcessExecution struct {
	pid        int
	startTime  int
	finishTime int
	duration   int
}
