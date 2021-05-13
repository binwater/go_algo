package main

import "fmt"

func quickSort(arr []int, p, q int){
	//终止条件
	if p >= q {
		return
	}

	//查找分区位置
	div := findDiv2(arr, p, q)

	//递归查找
	quickSort(arr, p, div-1)
	quickSort(arr, div+1, q)
}

func findDiv1(arr []int, p, q int) int {
	index := p
	divValue := arr[q]

	for i:=p; i<q; i++{
		if arr[i] < divValue {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[q], arr[index] = arr[index], arr[q]
	return index
}

func findDiv2(arr []int, p, q int) int {
	left, right := make([]int, 0), make([]int, 0)
	divValue := arr[q]
	for i:=p; i<q; i++{
		if arr[i] < divValue {
			left = append(left, arr[i])
		} else  {
			right = append(right, arr[i])
		}
	}

	copy(arr[p:p+len(left)], left)
	arr[p+len(left)] = divValue
	copy(arr[p+len(left)+1:q+1], right)
	return p + len(left)
}


func main() {
	nums := []int{3,2,4,1,6,5}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
