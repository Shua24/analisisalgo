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
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 2},
		{1, 2},
		{1, 2},
		{1, 2},
		{1, 2},
		{1, 2},
		{1, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 2},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 3},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 4},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{2, 5},
		{3, 5},
		{3, 5},
		{3, 5},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 6},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 7},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{3, 8},
		{4, 8},
		{4, 8},
		{4, 8},
		{4, 8},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 9},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 10},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 11},
		{4, 12},
		{4, 12},
		{4, 12},
		{4, 12},
		{4, 12},
		{4, 12},
		{4, 12},
		{4, 12},
		{4, 12},
		{5, 12},
		{5, 12},
		{5, 12},
		{5, 12},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{5, 13},
		{6, 13},
		{6, 13},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 14},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 15},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 16},
		{6, 17},
		{6, 17},
		{6, 17},
		{6, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 17},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 18},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{7, 19},
		{8, 19},
		{8, 19},
		{8, 19},
		{8, 19},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 20},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 21},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 22},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{8, 23},
		{9, 23},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 24},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 25},
		{9, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 26},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 27},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{10, 28},
		{11, 28},
		{11, 28},
		{11, 28},
		{11, 28},
		{11, 28},
		{11, 28},
		{11, 28},
		{11, 28},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 29},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{11, 30},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 31},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 32},
		{12, 33},
		{12, 33},
		{12, 33},
		{12, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 33},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 34},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{13, 35},
		{14, 35},
		{14, 35},
		{14, 35},
		{14, 35},
		{14, 35},
		{14, 35},
		{14, 35},
		{14, 35},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 36},
		{14, 37},
		{14, 37},
		{14, 37},
		{14, 37},
		{14, 37},
		{14, 37},
		{14, 37},
		{14, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 37},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 38},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{15, 39},
		{16, 39},
		{16, 39},
		{16, 39},
		{16, 39},
		{16, 39},
		{16, 39},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 40},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 41},
		{16, 42},
		{16, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 42},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{17, 43},
		{18, 43},
		{18, 43},
		{18, 43},
		{18, 44},
		{18, 44},
		{18, 44},
		{18, 44},
		{18, 44},
		{18, 44},
		{18, 44},
		{18, 44},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 45},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{18, 46},
		{19, 46},
		{19, 46},
		{19, 46},
		{19, 46},
		{19, 46},
		{19, 46},
		{19, 46},
		{19, 46},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 47},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 48},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{19, 49},
		{20, 49},
		{20, 49},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 50},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 51},
		{20, 52},
		{20, 52},
		{20, 52},
		{20, 52},
		{20, 52},
		{20, 52},
		{20, 52},
		{21, 52},
		{21, 52},
		{21, 52},
		{21, 52},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 53},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 54},
		{21, 55},
		{21, 55},
		{21, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 55},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 56},
		{22, 57},
		{22, 57},
		{22, 57},
		{22, 57},
		{22, 57},
		{22, 57},
		{22, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 57},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 58},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
		{23, 59},
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
	fmt.Print(hrArray[467].hrs)
}
