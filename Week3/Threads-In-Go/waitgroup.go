package main

import (
	"fmt"
	"sync"
)



func foo(wg *sync.WaitGroup){
	fmt.Println("new routine")
	wg.Done()
}


func main(){

	var wg sync.WaitGroup
	wg.Add(3)
	go foo(&wg)
	go foo(&wg)
	go foo(&wg)
	wg.Wait() //wait for all 3 SR to complete
	fmt.Println("finished.")

}