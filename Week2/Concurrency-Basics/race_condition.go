package main

/*
ROSS LEON
2023.06.06
WEEK2 MODULE2-ACTIVITY
Assignment: Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

====== EXPLANATION ======
A race condition refers to unsyncing of shared resources; variable states for example amongst concurrent 'threads'/ gosubroutines.
The race-condition is produced by interleaving, when the processor/scheduler is handling a task's instructions, needs to then switch to another task, then returning to the original's tasks instructions to complete; thus producing non-deterministic results.
What an application/ process/ task needs to produce always are deterministic results. The sharing of a variable and its updated value thus may not be accurately communicated to all other unfinished go routines; thus producing mixed results.

In the example below, both go routines are updating the same variable (tally) which is initially 0 (zero). Since both go routines are executed simultaneously;
they each think it is 0; and they each add 1; thus resulting in 1 which is incorrect.
One should add 1 to 0 resulting in 1; then the second goroutine should add 1 to 1 resulting in 2.


*/

import "log"


func updateTally1(){
	tally ++
	log.Println("updated tally is: ", tally)
}

func updateTally2(){
	tally ++
	log.Println("updated tally is: ", tally)
}

var tally int

func main(){
	//tally should be 0
	log.Println("tally at start: ", tally)

	go updateTally1() //add 1
	go updateTally2() //add 1

	//tally should be 2
	log.Println("tally at end:", tally)
}

/* ======= RESULT
2023/06/06 21:07:19 tally at start:  0
2023/06/06 21:07:19 tally at end: 0
2023/06/06 21:07:19 updated tally is:  1
2023/06/06 21:07:19 updated tally is:  1
*/