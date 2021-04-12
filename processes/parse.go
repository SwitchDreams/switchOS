package processes

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Parse parses input filr
func Parse(filename string) ([]Process, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var processList []Process

	for fileScanner.Scan() {
		line := fileScanner.Text()

		lineSplit := strings.Split(line, " ")
		arrival, _ := strconv.Atoi(lineSplit[0])
		duration, _ := strconv.Atoi(lineSplit[1])

		processList = append(processList, Process{
			arrivalTime: arrival,
			duration:    duration,
		})
	}

	return processList, nil
}
