# Week3 THREADS IN GO 

### GOROUTINES
- goroutines are threads that run concurrently w other goroutines
- 1 goroutine is created automatically to execute the main() function
- in a standard func call; main goroutine `BLOCKS` on the func call; while on a goroutine; this is handled by the `go runtime scheduler` 
- when the main goroutine ends; all other goroutines are forced to exit regardless if they completed or not

### EXITING GOROUTINES
- delayed exit w timer hack; will get intermittent errors/ varying results; non-deterministic behavior
    ```go fmt.Printf("something")
       time.Sleep(100 * time.Millisecond ) //timing is non-deterministic bc runtime scheduler may be using time for another process or OS for another app
       fmt.Println("finished.")
     ```

### BASIC SYNCHRONIZATION (formal synchronization constructs)
- Synchronization is when multiple threads agree on a timing of a global event
- Using global events whose execution is viewed by all threads; simultaneously
- X: Task1: x=1;x=x+1;<trigger global event> 
    Task2: if <global event> { print x }
    we are forcing for x=x+1 to execute before printing x; rather than allowing interleaving to trigger print x at x=1; thus forcing it to be deterministic
- synchronization thus reduces performance/ efficiency bc it behaves sequentially; but good in cases where sequence is required

### WAIT GROUPS
- Sync package includes WaitGroups
- sync.WaitGroup forces a goroutine to wait for other goroutines
- invoke with counter of routines
- decrease counter when each routine completes
- methods:
    * ADD(3)
    * DONE()
    * WAIT() //blocks until counter down to 0


### COMMUNICATION
- Channels are used between goroutines to communicate
- Channels are typed
- Use make() to create a channel 
    X: c:= make(chan int)
- Send and receive using <- operator
    X: Send c <- 3 //send 3 to channel c
    X: Receive x:= <- c //receive c into x

- 2 ways to send data to a goroutine: 1 via params, 2 via channels

### BLOCKING CHANNELS
- by default channels are `UNBUFFERED`  (cannot hold data in transit)
- Sending `BLOCKS` until data is `RECEIVED`
- Receiving `BLOCKS` until data is `SENT`
- USECASE1 X: @TASK1 c <-3  => (1 hour later...) => @TASK2 reaches x:= <-c 
    ~ @T1 will block for 1 hour until @T2 reads that channel; then @T1 can then continue
    >> @T1 block and send (c<-3) 
    >> wait... until T2 receives
    >> @T2 receives
    >> @T1 unblocks and continues exec

- USECASE2 X: Reverse 
    >> @T2 blocks and receives (but T1 has not sent)
    >> waits for T1 to send
    >> @T1 blocks and sends 
    >> @T2 receives
    >> @T1 continues 

- Therefore can also use as a synchronization method; and disregard message
    >> @T1 c <-3
    >> wait...
    >> @T2  <-c  //no variable to accept result; simply used to sync @T1 w @T2 (to catch up)

### BUFFERED CHANNELS
- Channels can maintain a limited number of objects 
- Capacity is the number of objects it can hold in transit thru the MAKE optional argument; default is 0
 X: make (chan int, `3` )
 >> SENDER ONLY BLOCKS IF BUFFER IS FULL 
- Useful when the producer @T1 is producing a lot of data and @T2 consumer is slower at processing that data; so a buffer will hold that data in route until consumed.
  and same in reverse when @T2 consumer is consuming too much
  