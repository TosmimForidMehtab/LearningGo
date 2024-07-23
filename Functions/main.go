package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func getNames() (string, string) {
	return "foo", "bar"
}

func checkEven(a int) bool {
	return a%2 == 0
}

func filter(a []int, f func(int) bool) []int {
	var result []int
	for _, v := range a {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	println(add(1, 2))
	println(sub(1, 2))
	l1, l2 := getNames()
	println(l1, l2)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(filter(arr, checkEven))

}
