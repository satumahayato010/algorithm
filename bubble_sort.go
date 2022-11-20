package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(nums []int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		for j := 0; j < lenNums-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
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
	output := bubbleSort(nums)
	fmt.Println(output)
}
