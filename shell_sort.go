package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shellSort(nums []int) []int {
	lenNums := len(nums)
	gap := lenNums / 2

	for gap > 0 {
		for i := gap; i < lenNums; i++ {
			temp := nums[i]
			j := i
			for j >= gap && nums[j-gap] > temp {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = temp
		}
		gap /= 2
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
	output := shellSort(nums)
	fmt.Println(output)
}
