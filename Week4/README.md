# Week4 SYNCHRONIZED COMMUNICATION 

>> ITERATING THRU A CHANNEL
When a a producer and subscriber are always communicating via a channel, then we iterate
X: `for i:= range c {  fmt.Println(i) }` //continually read from channel c
- reads continually until producer closes the channel via `CLOSE(c)`
- Close only required when this iterative pattern is used
 

## BLOCKING ON CHANNELS
>> MULTIPLE GOROUTINES 
Multiple channels may be used to receive communication
- T1 ---(c1)--->  T3  <----(c2)--- T2
- both channels will be blocked; ~ z = x + y

 
## SELECT
>> READ FROM MULTIPLE CHANNELS WITH SELECT
Can read channels with OR logic; will wait on the first channel that comes in
X: `select {
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
        ...
    }`
>> READ OR SEND WITH SELECT
Can either send or receive; whichever is unblockedd first; it is executed
X: `select {
    case a = <- inchan:
        fmt.Println("Received a")
    case outchan <- b:
        fmt.Println("sent b")
}`
>> SELECT WITH AN 'ABORT' CHANNEL PATTERN
When a message comes in an the abort channel; this triggers the RETURN
X: `for {
        select {
            case a <- c:
                fmt.Println(a)
            case <- abort:
                return 
        }
    }`

>> SELECT WITH A DEFAULT 
May want a default operation to avoid blocking (if none of the other channels have sent any data)
X: `select {
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
    default:
        fmt.Println("nop")
}`

## MUTUAL EXCLUSIONS
Code segments in different goroutines which cannot execute concurrently
Writing to shared variables should be mutually exclusive 

- Sharing variables concurrently can cause problems ~ produce non-deterministic results due to interleavings/ go runtine scheduler
X: race condition; where result is 1 not desired 2
    var i int =0
    var wg = sync.WaitGroup
    func inc() { i = i + 1 wg.Done() }
    func main(){
        wg.Add(2)
        go inc()
        go inc()
        wg.Wait()
        fmt.Println(i)
    }

## MUTEX (Sync.Mutex)
- To prevent interleaving issues; dont write to shared variable at the same time
- A Mutex ensures mutual exclusion


## SYNC MUTEX METHODS
- Lock() 'puts the flag up'; shared variable is in use
    - if lock is already taken by a goroutine, Lock() blocks until the flag is put down
- Unlock() 'puts the flag down'; done updating shared variable
    - when Unlock() is called, a blocked Lock() can proceed; other waiting goroutines can then use and call Lock themselves to continue processing/ updating
- call Lock() at beginning of this MUTUALLY EXCLUSIVE REGION and Unlock() at end

X: race condition; where result is 1 not desired 2; FIXED WITH MUTEX
    var i int = 0
    var wg = sync.WaitGroup
    var mut sync.Mutex

    func inc() { 
        mut.Lock()
        i = i + 1 
        mut.Unlock()
        wg.Done() }
    func main(){
        wg.Add(2)
        go inc()
        go inc()
        wg.Wait()
        fmt.Println(i)
    }

## SYNC - ONCE SYNCHRONIZATION
- Lets say have a multi-threaded go app
    - and need to initialize; by definition, should happen only once
    - one way is to initialize in the main thread (MAIN)
- but sometimes in MAIN is not an option,thus `ONCE.DO(f)`
    - even if invoked in many goroutines, it will be executed ONLY ONCE
- All calls to once.Do() `BLOCK` until the first returns 
X:  ## pseudo-code:

    func A() {  print('abc'); once.do(); print('123') }
    func B() {  once.do(); print('def') }
    func C() {  print('ghi'); once.do(); print('456') }

    >>lets say func B oncedo is reached first; this will block fA and fC; until fB returns;
    so if fA is reached; @oncedo will be blocked; same with fC

X2: 
    var on sync.Once    //=== define
    var wg sync.WaitGroup

    func setup(){ fmt.Println("Init") } //====== init func

    func dostuff() {
        on.Do(setup) //========== initialize before rest of routine code
        fmt.Println("hello")
        wg.Done()
    }

    func main(){
        wg.Add(2)
        go dostuff()
        go dostuff()
        wg.Wait()
    }

    RESULT:
        Init
        hello
        hello


## DEADLOCK 
Deadlock comes from synchronization dependencies; One goroutine to depend on another; each to depend on each other; A->B->A; ~circular self-referencing

X:

func dostuff(c1 chan int, c2 chan int) {
    <- c1
    c2 <- 1
    wg.Done()
}

func main(){
    ch1 := make (chan int)
    ch2 := make (chan int)
    wg.Add(2)
    go dostuff(ch1, ch2) //!!
    go dostuff(ch2, ch1) //!!
    wg.Wait()
}



## DINING PHILOSOPHERS
- deadlock example of 5 philosophers eating rice w limited chopsticks

- The classic Dining Philosophers problems represents an OS dealing with resource allocation; where Philosophers represent processes
and the chopsticks represent shared limited resources which must be shared between processes in a synchronized manner.

- A semaphore is an abstract datatype used to control access to a common shared resource by multiple threads.
- Semaphores are a type of synchronization primitive.
- Semaphores are a useful tool in the prevention of race conditions.
- Binary sempahores are used to implement locks; 0/1
