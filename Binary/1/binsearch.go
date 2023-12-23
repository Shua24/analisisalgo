package main

import (
	"fmt"
)

type hrstruct struct {
	hrs, mins int
}

func BinarySearch(arr []hrstruct, targetHrs int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid].hrs == targetHrs {
			return mid
		} else if arr[mid].hrs < targetHrs {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {

	hrArray := []hrstruct{
		{1, 0},
		{1, 4},
		{2, 16},
		{2, 19},
		{3, 23},
		{5, 32},
		{8, 32},
		{11, 37},
		{13, 46},
		{18, 53},
	}

	targetHrs := 11
	result := BinarySearch(hrArray, targetHrs)

	if result != -1 {
		fmt.Printf("Target hours %d found at index %d\n", targetHrs, result)
	} else {
		fmt.Printf("Target hours %d not found in the array\n", targetHrs)
	}
}
