package main

import (
	"log"
	"sync"
)

/*
ROSS LEON
2023.06.09
UCI GOLANG COURSE3
MODULE 4 ACTIVITY

============  DINING PHILOSOPHERS PROBLEM
Implement the dining philosopher’s problem with the following constraints/modifications.

1) There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

2) Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

3) The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

4) In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

5) The host allows no more than 2 philosophers to eat concurrently.

6) Each philosopher is numbered, 1 through 5.

7) When a philosopher starts eating (after it has obtained necessary locks) it
prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

8) When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.

============
The classic Dining Philosophers problems represents an OS dealing with resource allocation; where Philosophers represent processes
and the chopsticks represent shared limited resources which must be shared between processes in a synchronized manner.
*/


type ChopStick struct {
	sync.Mutex
}
type Philosopher struct {
	number int	// number of the philosopher
	ateQty int // number of times a philosopher has eaten
	LeftChopstick, RightChopStick *ChopStick
}
var wg, wg2, wg3 sync.WaitGroup
//create list of chopsticks
var chopSticks = make( []*ChopStick, 5)
//create list of philo
var philosophers = make( []*Philosopher, 5)

func (p *Philosopher) eat() {	
	//can only eat 3x 	  
	if p.ateQty <3 {
		p.LeftChopstick.Lock()
		p.RightChopStick.Lock()
		log.Printf("Philosopher-%v starting to eat (%vx)", p.number , p.ateQty +1)
		p.RightChopStick.Unlock()
		p.LeftChopstick.Unlock()		 
		log.Printf("Philosopher-%v finished eating (%vx)", p.number , p.ateQty +1)
		//update ate qty
		p.ateQty = p.ateQty +1 
	} 
	wg3.Done()
}

func removeFromQueue(s []*Philosopher, i int) {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1] 
}

func hostDining(philosophers []*Philosopher) { 
	//start hosting dining
	i:=0  

	//iterate until all philosophers have eaten 3x
	for len(philosophers) > 0 { 
		//grant permission to eat; only 2 phiolosophers at a time 
		//check if have eaten less than 3x?
		if philosophers[i].ateQty <3 {
				log.Println("------------------ maximum (2) philosophers may eat at a time")
				wg3.Add(1)
				//permission to eat
				go philosophers[i].eat() 			 
		} else {
				//remove philosopher from the dining table
				removeFromQueue(philosophers, i)
				philosophers = philosophers[:len(philosophers)-1] 
				i=0
				continue
			}
		
		//allow 2 philosophers to eat at a time
		if len(philosophers) != 1 {
				if philosophers[(i+2)% len(philosophers)].ateQty <3 {
					wg3.Add(1)
					go philosophers[(i+2)%5].eat()  
				} 
		}

		//wait for 2 philosophers to finish eating
		wg3.Wait()		 
		i++
		i = i%len(philosophers) 		 
	} 
	//finish hosting
	wg2.Done()
}

func main (){ 
	log.Printf("\n\n(5) philosophers can eat maximum 3x; only 2 at a time.\n\n")
	//init list of chopsticks
	for i:=0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	} 
	//init list of philosophers
	for i:=0; i < 5; i++ {
		philosophers[i] = &Philosopher{ i+1, 0, chopSticks[i], chopSticks[ (i+1) %5] }
		}

	wg2.Add(1)
	go hostDining(philosophers)
	//wait until all philosophers have eaten before exiting
	wg2.Wait() 
	log.Printf("\n\nAll philosophers finished eating 3x.\n\n")
}