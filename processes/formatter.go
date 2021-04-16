package processes

import (
	"fmt"
	"log"
	"os"
)

func ExecutionToFile(processesExecutionList []ProcessExecution, filename string) {
	f, err := os.Create("./output/" + filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, processExecution := range processesExecutionList {
		_, err := f.WriteString(fmt.Sprintf("Rodar processo [%d] de [%d] ate [%d]\n",
			processExecution.pid, processExecution.startTime, processExecution.finishTime))
		if err != nil {
			return
		}
	}
}
