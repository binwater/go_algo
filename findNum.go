package main

import "fmt"

func findTarget(arr []int, target int) int{
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left+right)/2
		if arr[mid]==target{
			return mid
		} else if arr[mid] < target{
			left = mid+1
		} else {
			right = mid-1
		}
	}
	return -1
}

func main() {
	nums := []int{1,2,3,4,5,6}
	fmt.Println(findTarget(nums, 1))

	fmt.Println(findTarget(nums, 6))

	fmt.Println(findTarget(nums, 4))
}
