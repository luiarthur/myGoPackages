package myStat

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"io/ioutil" //ReadMatrix
	"strconv"   //ReadMatrix
	"strings"   //ReadMatrix
)

// Creates a matrix of ones, with parameters row and col.
// Example:
//   The following creates a matrix of 0's of dimension 2x3
//   and assigns it to the variable m
//   	 m := myStat.M1(2,3)
func M0(row int, col int) mat64.Dense {
	v := make([]float64, row*col)
	M := mat64.NewDense(row, col, v)
	return *M
}

// I wrote this out of frustration because I couldn't figure
// out how to use the "mat64.Format" function to print matrices.
// Future work: speed check, argument for digits
// Example:
//   myStat.PrintMatrix(myStat.M1(3,4))
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

func ReadMatrix(file string) mat64.Dense {
	bs, _ := ioutil.ReadFile(file)
	str := string(bs)

	rowStr := strings.Split(str, "\n")
	rows := len(rowStr) - 1
	cols := len(strings.Split(rowStr[0], " "))
	dat := M0(rows, cols)
	for i := 0; i < rows; i++ {
		colStr := strings.Split(rowStr[i], " ")
		for j := 0; j < cols; j++ {
			d, _ := strconv.ParseFloat(colStr[j], 64)
			dat.Set(i, j, d)
		}
	}

	return dat
}
