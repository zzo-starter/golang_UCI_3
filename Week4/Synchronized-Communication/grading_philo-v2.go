/*

Implement the dining philosopherâ€™s problem with the
following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with
one chopstick between each adjacent pair of philosophers.

2. Each philosopher should eat only 3 times (not in an
infinite loop as we did in lecture)

3. The philosophers pick up the chopsticks in any order,
not lowest-numbered first (which we did in lecture).

4. In order to eat, a philosopher must get permission from
a host which executes in its own goroutine.

5. The host allows no more than 2 philosophers to eat
concurrently.

6. Each philosopher is numbered, 1 through 5.

7. When a philosopher starts eating (after it has obtained
necessary locks) it prints â€œstarting to eat <number>â€ on a
line by itself, where <number> is the number of the
philosopher.

8. When a philosopher finishes eating (before it has
released its locks) it prints â€œfinishing eating <number>â€
on a line by itself, where <number> is the number of the
philosopher.

*/

package main

import (
	"fmt"
	"sync"
)

// Defining the chopsticks and the philosophers
type CStick struct{ sync.Mutex }
type Philo struct {
	num     int
	leftCS  *CStick
	rightCS *CStick
}

// Defining the eat() method for philosophers
func (p *Philo) eat(wg *sync.WaitGroup, diningPasses *[]*sync.Mutex) {

    // Acquire a dining pass
    grantedPass := retrieveAvailablePass(diningPasses)

    // Pick up chopsticks
    p.leftCS.Lock()
    p.rightCS.Lock()

    fmt.Printf("Starting to eat: %d\n", p.num)
    fmt.Printf("Finished eating: %d\n", p.num)

    // Drop chopsticks
    p.rightCS.Unlock()
    p.leftCS.Unlock()

    // Signal done eating
    grantedPass.Unlock()

    // Clean up waitblock
    wg.Done()
}

func retrieveAvailablePass (diningPasses *[]*sync.Mutex) *sync.Mutex{
    // Blocks until one of the diningPasses is available
    // and returns it

    c := make(chan int)

    for i := 0; i < len(*diningPasses); i++ {
        go acquirePass((*diningPasses)[i], i, &c)
    }
    selectedPassNum := <- c
    return (*diningPasses)[selectedPassNum]
}

func acquirePass (diningPass *sync.Mutex, passNum int, c *chan int) {
    // Blocks until the diningPass is acquired
    diningPass.Lock()
    *c <- passNum
}


// Debugging: Defining a think() method
func (p *Philo) think(){
    fmt.Printf("I am Philosopher %d\n", p.num)
}

func main() {
	// Init variables
    numPhil := 5
	CSticks := make([]*CStick, numPhil)
	philos := make([]*Philo, numPhil)
    reqChans := make([]chan bool, numPhil)
    approveChans := make([]chan bool, numPhil)
    wg := sync.WaitGroup{}
    diningPasses := []*sync.Mutex{new(sync.Mutex), new(sync.Mutex)}

	// Init chopsticks
	for i := 0; i < 5; i++ {
		CSticks[i] = &CStick{sync.Mutex{}}
	}

    // Init philosophers and corresponding channels
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			num:     i,
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5],
		}
        reqChans[i] = make(chan bool)
        approveChans[i] = make(chan bool)
	}

    // Debugging: Checking that function calls are correct
    for _, philo := range philos {
        philo.think()
    }

	// Specify WaitGroup for 3 loops * 5 philosophers
	wg.Add(15)

    // Let the feast begin!
	for i := 0; i < 3; i++ {
        for _, philo := range philos {
			go philo.eat(&wg, &diningPasses)
		}
	}
	wg.Wait()
}