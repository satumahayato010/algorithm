package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(nums []int) []int {
	lenNums := len(nums)

	for i := 1; i < lenNums; i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j -= 1
		}
		nums[j+1] = temp
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
	output := insertionSort(nums)
	fmt.Println(output)
}
