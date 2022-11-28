package main

import (
	"fmt"
	"strings"
)

func MeanMedian(arrayInput []float64) (float64, float64) {

	var mean, median float64

	//mean = panjang / 2

	for i := 0; i < len(arrayInput); i++ {
		mean += arrayInput[i]
	}
	mean = mean / float64(len(arrayInput)) // len <- tipenya int

	var p float64 = float64(len(arrayInput)+1) / 2 //menampung array = 4
	S := fmt.Sprintf("%f", p)                      // string ganjil / genap
	if strings.Contains(S, ".5") {                 // genap
		l := p - 1
		median = (arrayInput[int(l-0.5)] + arrayInput[int(l+0.5)]) / 2
	} else { //ganjil
		median = arrayInput[int(p+1)/2]
	}
	return mean, median

}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
