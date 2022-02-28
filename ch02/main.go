package main

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/mat"
)

func createDataSet() (mat.Matrix, []string) {
	return mat.NewDense(4, 2, []float64{1.0, 1.1, 1.0, 1.0, 0, 0, 0, 0.1}), []string{"A", "A", "B", "B"}
}

func classify0(in []float64, dataSet mat.Matrix, labels []string, k int) string {
	dataSetSize, _ := dataSet.Dims()
	tileMat := mat.NewDense(dataSetSize, len(in), nil)
	for r := 0; r < tileMat.RawMatrix().Rows; r++ {
		tileMat.SetRow(r, in)
	}
	fmt.Printf("tileMat: %v\n", mat.Formatted(tileMat))
	tileMat.Sub(tileMat, dataSet)
	fmt.Printf("diffMat: %v\n", mat.Formatted(tileMat))

	tileMat.Apply(func(i, j int, v float64) float64 {
		return v * v
	}, tileMat)
	fmt.Printf("diffMat: %v\n", mat.Formatted(tileMat))
	sqDistance := mat.NewDense(dataSetSize, 1, nil)
	for r := 0; r < tileMat.RawMatrix().Rows; r++ {
		d := float64(0)
		row := tileMat.RowView(r)
		for c := 0; c < row.Len(); c++ {
			d += row.At(c, 0)
		}

		sqDistance.SetRow(r, []float64{d})
	}
	fmt.Printf("sqDistance: \n%v\n", mat.Formatted(sqDistance))

	sqDistance.Apply(func(i, j int, v float64) float64 {
		return math.Sqrt(v)
	}, sqDistance)
	fmt.Printf("sqDistance: \n%v\n", mat.Formatted(sqDistance))

	distance := make([]float64, sqDistance.RawMatrix().Rows)
	copy(distance, sqDistance.RawMatrix().Data)

	sort.Sort(sort.Float64Slice(distance))

	return ""
}

func main() {
	group, labels := createDataSet()
	fmt.Printf("%v\n", mat.Formatted(group))
	fmt.Println(labels)

	fmt.Printf("classify: %v\n", classify0([]float64{0.0, 0.0}, group, labels, 1))
}
