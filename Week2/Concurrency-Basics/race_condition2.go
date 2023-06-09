package main

//go run -race race_condition2.go

import (
	"log"
	"time"
)

var count int 

func race() {
    count++
	log.Println("count: ", count)
}

func main() {
    go race()
    go race()
    time.Sleep(1 * time.Second)
}

 