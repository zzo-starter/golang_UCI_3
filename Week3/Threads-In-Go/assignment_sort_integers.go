package main

/*
ROSS LEON
2023.06.07
UCI GOLANG COURSE3
MODULE 3 ACTIVITY

============ SORT ARRAY OF INTEGERS
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.

*/

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)



func sortSubArray(ia []int, c chan []int){
	fmt.Println("will sort array of integers: ", ia) 
	sort.Ints(ia)
	fmt.Println("sorted array of integers: ", ia) 
	c <- ia
}

func getArrayOfIntegers()[]int{
	var ai []int
	var s string  
	var i int =1 

	fmt.Println("Enter 12 numbers to sort:")
	for {
		fmt.Printf("%d.) ",i)
		 _, err := fmt.Scan(&s)
		if err != nil {
			log.Println(err)
		} 

		ii, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		i++
		ai = append(ai, ii)
		if i >=13 {
			break
		}
	}
	return ai
}

func main(){

	aui:= getArrayOfIntegers()

	c:= make( chan []int)
	go sortSubArray(aui[:3],   c)
	go sortSubArray(aui[3:6],  c)
	go sortSubArray(aui[6:9],  c)
	go sortSubArray(aui[9:12], c)

	var sa []int 
	for i:=0; i<=3; i++{
		for _, n:=range <-c {
			sa = append(sa, n)
		}
	}
	sort.Ints(sa)
	fmt.Println("\nSORTED ARRAY OF INTEGERS: ",sa)
}