package main

import (
	"fmt"
)
/*
 What is a Race Condition? 

 A race condition happens when two concurrent processes try to modify the same set of data or the 
 same resources.
 
 Below depending on which function is being executed, a we cannot be sure of the actual value the go function 
 will print. 
 */

func incrementNumber(value int) int {
	go func() {
		value++;
		fmt.Println(value);
	}()
	fmt.Println(value);
	return value;
}

func main() {
	incrementNumber(4);
	incrementNumber(3);
}

