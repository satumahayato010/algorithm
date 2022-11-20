package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(nums []int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		minIdx := i
		for j := i + 1; j < lenNums; j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}

func main() {
	var nums []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}

	fmt.Println(nums)
	fmt.Println("############################")
	output := selectionSort(nums)
	fmt.Println(output)
}
