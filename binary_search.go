package main

import "fmt"

func binarySearch(nums []int, value int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == value {
			return mid
		}
		if nums[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(nums []int, value int, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == value {
		return mid
	} else if nums[mid] < value {
		return binarySearchRecursive(nums, value, mid+1, high)
	} else {
		return binarySearchRecursive(nums, value, low, mid-1)
	}
}

func main() {
	nums := []int{0, 1, 7, 8, 10, 13, 14, 18, 24}

	output := binarySearchRecursive(nums, 10, 0, len(nums))
	fmt.Println(output)
}
