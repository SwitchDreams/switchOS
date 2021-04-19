package processes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/switchdreams/switchOS/errors"
)

// Parse parses input filr
func Parse(filename string) ([]Process, errors.IOSError) {
	var contador int
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.WrapError(err, fmt.Sprintf("Failed to open input file '%s'", filename))
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var processList []Process

	for fileScanner.Scan() {
		line := fileScanner.Text()

		lineSplit := strings.Split(line, " ")
		arrival, _ := strconv.Atoi(lineSplit[0])
		Duration, _ := strconv.Atoi(lineSplit[1])

		processList = append(processList, Process{
			ID:          contador,
			ArrivalTime: arrival,
			Duration:    Duration,
		})
		contador++
	}

	return processList, nil
}
