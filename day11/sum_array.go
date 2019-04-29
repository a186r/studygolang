package main

import "fmt"

func main() {
	array1 := []float32{1.0, 2.0, 3.3}
	sum1 := Sum(array1)
	fmt.Printf("和是：%f\n", sum1)

	a := []int{1, 2, 3, 4, 5}

	sum, avg := SumAndAverage(a)
	fmt.Printf("和是:%d，平均数是:%f\n", sum, avg)
}

func Sum(arrF []float32) (sum float32) {

	for _, val := range arrF {
		sum += val
	}

	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, val := range a {
		sum += val
	}

	return sum, float32(sum / len(a))
}
