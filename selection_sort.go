package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(unsortedArr []int) []int {
	var l int = len(unsortedArr)

	for i := 0; i < l; i++ {
		indOfSmallest := i
		for j := i; j < l; j++ {
			if unsortedArr[j] < unsortedArr[indOfSmallest] {
				indOfSmallest = j
			}
		}
		unsortedArr[i], unsortedArr[indOfSmallest] = unsortedArr[indOfSmallest], unsortedArr[i]
	}
	return unsortedArr
}

func createArr(elemsCount int) []int {
	arr := make([]int, elemsCount)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < elemsCount; i++ {
		arr[i] = rand.Intn(elemsCount)
	}
	return arr
}

const amount = 100

func main() {
	slice := createArr(amount)

	fmt.Printf("Raw arr: %v\n", slice)
	fmt.Printf("Sorted arr: %v\n", selectionSort(slice))
}
