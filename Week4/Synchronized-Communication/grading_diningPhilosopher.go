package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	num             int
}

var reservation int = 0

func (p Philo) eat(request chan int, approve chan int, doneEating chan int, wg2 *sync.WaitGroup) {

	for i := 0; i < 3; i++ {
		request <- 1
		<-approve
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat Philo%d\n", p.num)
		fmt.Printf("finishing eating Philo%d\n", p.num)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		doneEating <- 1
	}
	wg2.Done()
}
func host(request chan int, approve chan int, abort chan int, doneEating chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-request:
			if reservation == 2 {
				<-doneEating
				approve <- 1
			} else if reservation < 2 {
				reservation++
				approve <- 1
			}

		case <-doneEating:
			reservation--

		case <-abort:
			fmt.Println("Bye")
			wg.Done()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	request := make(chan int)
	approve := make(chan int)
	doneEating := make(chan int)
	abort := make(chan int)
	wg.Add(1)
	go host(request, approve, abort, doneEating, &wg)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go philos[i].eat(request, approve, doneEating, &wg2)
	}
	wg2.Wait()

	abort <- 1
	wg.Wait()
}
