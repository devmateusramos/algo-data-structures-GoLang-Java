package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	result := Solution2(array, 10)
	fmt.Println(result)

}

// Solution1 TimeO(n²) | Space O(1)
func Solution1(array []int, targetSum int) []int {
	for i := 0; i < len(array); i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if targetSum == firstNum+secondNum {
				result := []int{firstNum, secondNum}
				return result
			}
		}
	}

	return []int{}
}

// Solution2 Time O(n) | Space O(n)
func Solution2(array []int, targetSum int) []int {
	nums := make(map[int]bool)
	for _, num := range array {
		potentialMatch := targetSum - num
		if nums[potentialMatch] {
			result := []int{num, potentialMatch}
			return result
		}
		nums[num] = true
	}

	return []int{}
}

// Solution3 Time O(n * log(n)) | Space O(1)
func Solution3(array []int, targetSum int) []int {
	sort.Ints(array)
	left := 0
	right := len(array) - 1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == targetSum {
			result := []int{array[left], array[right]}
			return result
		} else if currentSum < targetSum {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}
