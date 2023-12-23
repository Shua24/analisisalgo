package main

import (
	"fmt"
)

type hrstruct struct {
	hrs, mins int
}

func Isearch(arr []hrstruct, targetHrs int) int {
	low, high := 0, len(arr)-1
	for low <= high && targetHrs >= arr[low].hrs && targetHrs <= arr[high].hrs {
		pos := low + ((targetHrs - arr[low].hrs) * (high - low) / (arr[high].hrs - arr[low].hrs))
		if arr[pos].hrs == targetHrs {
			return pos
		}
		if arr[pos].hrs > targetHrs {
			high = pos - 1
		} else {
			low = pos + 1
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
	// fmt.Println(rand.Intn(24))

	targetHrs := 11
	result := Isearch(hrArray, targetHrs)

	if result != -1 {
		fmt.Printf("Target hour %d found at index %d\n", targetHrs, result)
	} else {
		fmt.Printf("Target hour %d not found in the array\n", targetHrs)
	}
}
