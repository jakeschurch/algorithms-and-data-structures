package algos

import (
	"fmt"
)

func main2() {
	l1 := []int{9, 6, 4, 3, 5, 7}

	a := mergeSort(l1)
	fmt.Println(a)
}

func mergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	mid := len(A) / 2
	return Merge(mergeSort(A[:mid]), mergeSort(A[mid:]))
}

// Merge takes an unsorted slice, seperates them into two slice to sort,
// and finally merge them into one sorted slice.
func Merge(left, right []int) []int {
	i := 0
	j := 0

	result := make([]int, len(left)+len(right))

	for k := 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			j++
			continue
		} else if j >= len(right) {
			result[k] = left[i]
			i++
			continue
		}

		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}

// insertion sort in increasing order
func insertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]

		i := j - 1

		for i >= 0 && A[i] > key { // shift
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
	return A
}

// Ex. 2.1-2; insertion sort in decreasing order
func reversedInsertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1

		for i >= 0 && A[i] < key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
		fmt.Println(A)
	}
	return A
}

// Ex. 2.1-3 Linear Search
func linearSearch(A []int, v int) (val int, ok bool) {
	for i := range A {
		if A[i] == v {
			return i, true
		}
	}
	return 0, false
}

func selectionSort(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		min := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		key := A[i]
		A[i] = A[min]
		A[min] = key
	}
	fmt.Println(A)
	return A
}

func selectionSort2(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		min := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		for j := min; j > i; j-- {
			key := A[j]
			A[j] = A[j-1]
			A[j-1] = key
		}
		fmt.Println(A)
	}
	return A
}
