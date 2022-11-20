package main

import (
	"fmt"
	"math/rand"
	"time"
)

func linearSearch(nums []int, value int) int {
	for i, num := range nums {
		if num == value {
			return i
		}
	}
	return -1
}

func main() {
	var nums []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)

	output := linearSearch(nums, 18)
	fmt.Println(output)
}
