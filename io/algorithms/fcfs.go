package algorithms

import (
	"math"

	io "github.com/switchdreams/switchOS/io"
)

func Fcfs(disk io.Disk) int {
	// Inicializando com o valor Init (primeira posição - ponto de partida)
	countCylinders := math.Abs(float64(disk.Sequence[0] - disk.Init))
	for i := 0; i < (len(disk.Sequence) - 1); i++ {
		countCylinders += math.Abs(float64(disk.Sequence[i] - disk.Sequence[i+1]))
	}

	return int(countCylinders)
}
