/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	pnum            int
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup
var hwg sync.WaitGroup

var gAskPermission chan int
var gRecPermission [6]chan string

func (p Philo) eat() {

	cl := make(chan string)
	cr := make(chan string)

	// eat three times
	for i := 0; i < 3; i++ {

		// ask permission, send request to eat
		//fmt.Println("asking ", p.pnum)
		gAskPermission <- p.pnum
		// wait for permission
		<-gRecPermission[p.pnum] // philospher numbering starts at 1, so we're using slots 1 to 5

		//look for chopsticks
		go func() { p.leftCS.Lock(); cl <- "locked" }()
		go func() { p.rightCS.Lock(); cr <- "locked" }()

		// wait to pick up the chopsticks, in any order
		for j := 0; j < 2; j++ {
			select {
			case <-cl:
				//	got it, move on
			case <-cr:
				//	got it, move on
			}
		}
		fmt.Println("starting to eat ", p.pnum)
		// yum, tasty
		fmt.Println("finished eating ", p.pnum)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func host() {

	// manage the request and permission queues (channels)
	// so that no more than two philosphers eat at the same time

	for i := 0; i < 7; i++ {

		// we will allow 5 x 3 = 15 dinings

		// take two off the request channel
		p1 := <-gAskPermission
		p2 := <-gAskPermission
		//		fmt.Println("requested ", p1, p2)

		var lwg sync.WaitGroup
		lwg.Add(2)
		go func() { gRecPermission[p1] <- "ok"; lwg.Done() }()
		go func() { gRecPermission[p2] <- "ok"; lwg.Done() }()
		lwg.Wait()
	}
	p1 := <-gAskPermission
	gRecPermission[p1] <- "ok"
	hwg.Done()
}

func main() {

	//initialize global receive permission channels
	for i := range gRecPermission {
		gRecPermission[i] = make(chan string)
	}
	gAskPermission = make(chan int) //buffer size two used to limit the number of simultaneous diners.

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	hwg.Add(1)
	go host() // start host and let it run, don't wait

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()

	//close all the channels
	close(gAskPermission)
	for i := range gRecPermission {
		close(gRecPermission[i])
	}
	hwg.Wait()
}
