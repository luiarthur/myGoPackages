package myStat

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
)

func M1(row int, col int) mat64.Dense {
	v := make([]float64, row*col)
	M := mat64.NewDense(row, col, v)
	return *M
}

func PrintMatrix(m mat64.Dense) (n int, err error) {
	r, c := m.Caps()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%f  ", m.At(i, j))
		}
		fmt.Printf("\n")
	}
	return
}
