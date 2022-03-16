package main

import "fmt"

func binarySearch(sortedList []int, SearchedValue int) (int, error) {
	var low int = 0
	var high int = len(sortedList) - 1
	var steps int = 0

	for low <= high {
		mid := (low + high) / 2

		steps += 1
		fmt.Println(steps)

		if sortedList[mid] == SearchedValue {
			return mid, nil
		} else if sortedList[mid] < SearchedValue {
			low = mid + 1
		} else if sortedList[mid] > SearchedValue {
			high = mid - 1
		}
	}
	return -1, fmt.Errorf("no value %d in array", SearchedValue)
}

func main() {
	var arr [10000000]int
	for i := 0; i < 10000000; i++ {
		arr[i] = i
	}
	val := 0
	list := arr[:len(arr)]

	pos, err := binarySearch(list, val)
	if err == nil {
		fmt.Printf("Element %d has %d position\n", val, pos)
	} else {
		fmt.Println(err)
	}
}
