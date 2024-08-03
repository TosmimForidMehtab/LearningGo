package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func addNums[T Number](numbers []T) T {
	var sum T
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := addNums(nums)
	nums2 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	sum2 := addNums(nums2)
	fmt.Println("Int sum -> ", sum)
	fmt.Println("float32 sum -> ", sum2)

}
