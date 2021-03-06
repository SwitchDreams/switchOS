package memory

import (
	"bufio"
	"fmt"
	"github.com/switchdreams/switchOS/errors"
	"os"
	"strconv"
)

// Parse parses input file
func Parse(filename string) (Memory, errors.IOSError) {
	var cont int
	var memory Memory

	file, err := os.Open(filename)
	if err != nil {
		return memory, errors.WrapError(err, fmt.Sprintf("Failed to open input file '%s'", filename))
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		Num, _ := strconv.Atoi(line)
		if cont == 0 {
			memory.Size = Num
		} else {
			memory.Sequence = append(memory.Sequence, Num)
		}
		cont++
	}

	return memory, nil
}
