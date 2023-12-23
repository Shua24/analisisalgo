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
		{1, 0},
		{1, 2},
		{1, 3},
		{1, 3},
		{1, 3},
		{1, 3},
		{1, 4},
		{1, 4},
		{2, 5},
		{2, 6},
		{2, 6},
		{2, 7},
		{2, 7},
		{2, 7},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 9},
		{3, 10},
		{3, 10},
		{4, 10},
		{4, 10},
		{4, 11},
		{4, 11},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 15},
		{6, 15},
		{6, 15},
		{6, 16},
		{6, 16},
		{6, 17},
		{7, 18},
		{8, 19},
		{8, 19},
		{8, 21},
		{8, 22},
		{8, 23},
		{9, 23},
		{9, 23},
		{9, 23},
		{9, 24},
		{9, 25},
		{10, 25},
		{10, 26},
		{10, 27},
		{10, 29},
		{10, 30},
		{10, 31},
		{10, 32},
		{10, 33},
		{11, 34},
		{12, 34},
		{12, 35},
		{12, 35},
		{12, 35},
		{12, 35},
		{12, 35},
		{12, 35},
		{13, 36},
		{13, 36},
		{13, 37},
		{13, 38},
		{15, 39},
		{15, 39},
		{15, 41},
		{16, 42},
		{16, 42},
		{16, 42},
		{16, 42},
		{16, 43},
		{16, 43},
		{16, 43},
		{16, 43},
		{17, 44},
		{17, 45},
		{18, 47},
		{18, 47},
		{19, 47},
		{19, 48},
		{19, 48},
		{20, 48},
		{20, 49},
		{20, 51},
		{21, 51},
		{21, 53},
		{21, 53},
		{22, 53},
		{22, 54},
		{22, 54},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 57},
		{23, 59},
		{23, 59},
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
