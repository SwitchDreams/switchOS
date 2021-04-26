package io

import (
	"bufio"
	"fmt"
	"github.com/switchdreams/switchOS/errors"
	"os"
	"strconv"
)

// Parse parses input file
func Parse(filename string) (Disk, errors.IOSError) {
	var cont int
	var disk Disk

	file, err := os.Open(filename)
	if err != nil {
		return disk, errors.WrapError(err, fmt.Sprintf("Failed to open input file '%s'", filename))
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		Num, _ := strconv.Atoi(line)
		if cont == 0 {
			disk.Size = Num
		} else if cont == 1 {
			disk.Init = Num
		} else {
			disk.Sequence = append(disk.Sequence, Num)
		}
		cont++
	}

	return disk, nil
}
