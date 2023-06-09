package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

var counter int

func sorting(a []int, ch chan []int) {
	defer wg.Done()
	counter++
	fmt.Println("goroutine", a)
	sort.Ints(a)
	ch <- a
	if counter == 4 {
		abortChan <- struct{}{}
	}
}

var wg sync.WaitGroup
var abortChan = make(chan struct{})

func main() {
	wg.Add(4)
	numsChan := make(chan []int, 4)
	var input string

	var nums []int

	for {
		fmt.Println("entern any number of int", nums, len(nums))
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}

		if input == "x" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		nums = append(nums, num)
	}

	sublength := len(nums) / 4
	subcap := cap(nums) / 4
	sli1 := nums[:sublength:subcap]
	sli2 := nums[sublength : sublength*2 : subcap*2]
	sli3 := nums[sublength*2 : sublength*3 : sublength*3]
	sli4 := nums[sublength*3 : len(nums) : cap(nums)]
	go sorting(sli1, numsChan)
	go sorting(sli2, numsChan)
	go sorting(sli3, numsChan)
	go sorting(sli4, numsChan)
	nums = nums[:0]

Loop:
	for {
		select {
		case v := <-numsChan:
			nums = append(nums, v...)
		case <-abortChan:
			break Loop
		default:
			continue
		}
	}
	sort.Ints(nums)
	fmt.Println(nums)
}