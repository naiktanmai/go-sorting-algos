package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := (len(arr) / 2)
		left := arr[:mid]
		right := arr[mid:]

		left = mergeSort(left)
		right = mergeSort(right)

		i := 0
		j := 0

		result := []int{}

		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		}

		for i < len(left) {
			result = append(result, left[i])
			i++
		}

		for j < len(right) {
			result = append(result, right[j])
			j++
		}
		return result
	}
	return arr
}

func main() {
	result := mergeSort([]int{12, 11, 13, 5, 6, 7})
	fmt.Println(result)
}
