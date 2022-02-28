package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func createDataSet() (mat.Matrix, []string) {
	return mat.NewDense(4, 2, []float64{1.0, 1.1, 1.0, 1.0, 0, 0, 0, 0.1}), []string{"A", "A", "B", "B"}
}

func main() {
	group, labels := createDataSet()
	fmt.Printf("%+v\n", group)
	fmt.Println(labels)
}
