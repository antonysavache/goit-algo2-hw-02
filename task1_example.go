package main

import "fmt"

func RunTask1Example() {
	numbers := []int{8, -3, 15, 4, 0, 23, -10, 7}
	minValue, maxValue := FindMinMax(numbers)

	fmt.Println("Task 1:")
	fmt.Println("Array:", numbers)
	fmt.Println("Min:", minValue)
	fmt.Println("Max:", maxValue)
	fmt.Println()
}
