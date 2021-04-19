package processes

import (
	"fmt"
	"os"

	"github.com/switchdreams/switchOS/errors"
)

// ExecutionToFile outputs the process execution list to a file
func ExecutionToFile(processesExecutionList []ProcessExecution, filename string) errors.IOSError {
	f, err := os.Create("./output/" + filename)
	if err != nil {
		return errors.WrapError(err, fmt.Sprintf("Failed to create output file '%s'", filename))
	}

	defer f.Close()

	for _, processExecution := range processesExecutionList {
		_, err := f.WriteString(fmt.Sprintf("Rodar processo [%d] de [%d] ate [%d]\n",
			processExecution.Pid, processExecution.StartTime, processExecution.FinishTime))
		if err != nil {
			return errors.WrapError(err, "Failed to write to output file")
		}
	}

	return nil
}
