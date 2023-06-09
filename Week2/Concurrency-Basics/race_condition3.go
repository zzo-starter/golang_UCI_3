// Demonstrate a race condition using goroutines

// When a program has multiple threads executing concurrently, there are many
// potentially possible interleavings of executions. By a race condition, we
// refer to the situation where the program results differ depending on the
// particular interleaving that happens.
//
// In the program below, there are 2 functions foo() and bar() that are executed
// as concurrent goroutines. Each waits a random amount of time before printing
// the value of x and incrementing it.
//
// The program behaves non-deterministically in that it produces different outputs
// depending on the actions of the O/S scheduler and go routine scheduler and the
// random wait times chosen.

// Two sample outputs are shown below. In the first example, bar() executes first,
// whereas in the second example, foo() executes before bar().

// In main() function
// In bar:  0
// In foo:  2
// Exiting main() function 3

// In main() function
// In foo:  1
// In bar:  2
// Exiting main() function 3

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var x int

func foo() {
    time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
    fmt.Println("In foo: ", x)
    x = x + 1
}

func bar() {
    time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
    fmt.Println("In bar: ", x)
    x = x + 1
}

func main() {
    fmt.Println("In main() function")
    go foo()
    go bar()
    time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
    x = x + 1
    time.Sleep(500 * time.Millisecond)
    fmt.Println("Exiting main() function", x)
}
