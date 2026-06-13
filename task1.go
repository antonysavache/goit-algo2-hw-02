package main

func FindMinMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		panic("numbers must not be empty")
	}

	return findMinMax(numbers, 0, len(numbers)-1)
}

func findMinMax(numbers []int, left, right int) (int, int) {
	if left == right {
		return numbers[left], numbers[left]
	}

	if right-left == 1 {
		if numbers[left] < numbers[right] {
			return numbers[left], numbers[right]
		}
		return numbers[right], numbers[left]
	}

	middle := left + (right-left)/2
	leftMin, leftMax := findMinMax(numbers, left, middle)
	rightMin, rightMax := findMinMax(numbers, middle+1, right)

	minValue := leftMin
	if rightMin < minValue {
		minValue = rightMin
	}

	maxValue := leftMax
	if rightMax > maxValue {
		maxValue = rightMax
	}

	return minValue, maxValue
}
