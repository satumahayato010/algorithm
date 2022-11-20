package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(nums []int, low, high int) int {
	i := low - 1
	pivot := nums[high]
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}

	}
	nums[i+1], nums[high] = nums[high], nums[i+1]

	return i + 1
}

func quickSort(nums []int, low, high int) []int {
	if low < high {
		partitionIndex := partition(nums, low, high)
		quickSort(nums, low, partitionIndex-1)
		quickSort(nums, partitionIndex+1, high)
	}
	return nums
}

func quickSortStart(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func main() {
	var nums []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}

	fmt.Println(nums)
	fmt.Println("############################")
	output := quickSortStart(nums)
	fmt.Println(output)
}
