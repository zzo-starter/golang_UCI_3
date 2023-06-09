package main

//go run -race race_condition1.go

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

func main() {

	counter := 0

	const num = 15
	var wg sync.WaitGroup //wait for a collection of goroutines to finish
	wg.Add(num) //wait for 15 goroutines


    var mu sync.Mutex  //lock access to counter variable until finished updating

	for i := 0; i < num; i++ {
		log.Println("@i-", i , "counter: ", counter )
		go func() {
			mu.Lock() //----- LOCK 
			log.Println("counter: ", counter)
			temp := counter
			log.Println("within gr> runtine.Gosched()")
			runtime.Gosched() //allow other goroutines to run; yields the processor
			temp++
			counter = temp
			mu.Unlock() //---- UNLOCK
			log.Println("within gr> counter:", counter, " decreasing wg by 1 ")
			wg.Done() //decrements the wg counter by 1
		}()
	}

	wg.Wait() //block until all go routines have finished
	fmt.Println("count:", counter)
}

 