package main

import (
	"fmt"
	"time"
)

/*
A race condition occurs when two or more threads can access shared data
and they try to change it at the same time. Because the thread
scheduling algorithm can swap between threads at any time,
you don't know the order in which the threads will
attempt to access the shared data. Therefore, the
result of the change in data is dependent on the
thread scheduling algorithm, i.e. both threads
are "racing" to access/change the data.
*/

/*
Run this program multiple time to see the different output due
to race between both increaseByOne & decreaseByOne functions
*/

var count int

func increaseByOne() {
	time.Sleep(1 * time.Millisecond)
	count++
}

func decreaseByOne() {
	count--
}

func main() {
	for i := 0; i < 200; i++ {
		go increaseByOne()
		go decreaseByOne()
	}
	fmt.Println(count)
}