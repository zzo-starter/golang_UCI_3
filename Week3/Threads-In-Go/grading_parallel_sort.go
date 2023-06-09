package main

import (
	"fmt"
	"sort"
	"sync"
)

// function called by multiple go routines for parallel sorting
func parallelSort(wg *sync.WaitGroup, start, end int, array []int){
	fmt.Println("Sorting Sub array : ", array[start:end])
	sort.Ints(array[start:end])
	wg.Done()
}

// merging all the buffers/ chunks by main routine
func merge(buffer_sizes , arr []int){
	buffer_index := make([]int, 4)
	sorted_arr := make([]int, len(arr))
	start_index := make([]int, 4)

	for i := 1; i < 4; i++{
		start_index[i] += buffer_sizes[i-1] + start_index[i-1]
	}

	for i := 0; i < len(arr); i++{
		min_val := 1 << 31 - 1
		min_index := 0
		for j := 0; j < 4; j++{
			if buffer_index[j] != -1 && 
				arr[start_index[j] + buffer_index[j]] < min_val {
					min_val = arr[start_index[j] + buffer_index[j]]
					min_index = j
			}
		}
		buffer_index[min_index]++
		if buffer_index[min_index] == buffer_sizes[min_index]{
			buffer_index[min_index] = -1
		}
		sorted_arr[i] = min_val
	}
	fmt.Println("Merged Arr : ", sorted_arr)
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func main(){
	// wait group because merge should be done after each 
	// go routine sorts its chunk/buffer data
	var wg sync.WaitGroup

	var n int
	fmt.Println("Enter Number of elements in array")
	fmt.Scan(&n)

	fmt.Println("Enter elements")
	arr := make([]int, n)
	for i, _ := range arr{
		fmt.Scan(&arr[i])
	}
	fmt.Println("Input array : ", arr)

	const NUM_BUFFERS = 4
	buffer_size := n / NUM_BUFFERS
	buffer_sizes := make([]int, 4)
	for i := 0; i < 4; i++{
		buffer_sizes[i] = buffer_size
	}

	rest := n - buffer_size * NUM_BUFFERS
	for i := 0; i < rest; i++{
		buffer_sizes[i]++
	}

	start := 0
	end := 0
	for i := 0; i < 4; i++{
		end += buffer_sizes[i]
		wg.Add(1)
		go parallelSort(&wg, start, end, arr)
		start = end
	}
	wg.Wait()
	merge(buffer_sizes, arr)
}
