# Week2 CONCURRENCY BASICS

### PROCESSES & THREADS
    - Process is an instance of a RUNNING PROGRAM that has
        - CODE
        - STACK; a region of memory that handles function calls 
        - HEAP; a region of memory that handles memory allocation 
        - REGISTERS; memory to store values 
        - PROGRAM COUNTER; a register that holds what instruction is being executed/ (the next one to be executed)
        - DATA REGISTERS; 
        - STACK POINTER; a register that holds where we are on the stack
        - SHARED LIBRARIES; shared between processes 
        - All this is considered the CONTEXT 
        - X: OPERATING SYSTEMS
            - multipple process management
            - process A, B, C may all be accessing the memory address 1000; but it is not actual address 1000; so OS manages actual addresses when sharing CPU
                - CPU is switching roughly every 20ms between processes (SCHEDULING); user has impression of parallelism
                - OS must give processes fair access to resources

### SCHEDULING 
    - X: @OS CPU swapping between processees every 20ms using different scheduling algorithms; x: round-robin; some w higher priority    
    - When a CPU changes from one process to another it is called CONTEXT SWITCHING; stores state (context) somewhere for later use; then loads new state for next process

### THREADS AND GOROUTINES
    - Threads share some context
    - Many threads can exist in one process 
    - X: PROCESS[ [ Virtual Memory File Descriptors ] THREAD1(stack, data registers, code), THREAD2(stack, data registers, code), THREAD3(stack, data registers, code) ]
        Threads read less; share VM resources; thus can exec more 
    - OS schedules threads rather than processes 
    ## GO ROUTINES ##
    - Go routines are threads in Go
    - X: PROCESS [ VM File Descriptors] MAIN THREAD[ GOROUTINE1, GOROUTINE2, GOROUTINE3 ]
        OS schedules 1 Main Thread; and within main thread go sub-schedules goroutines ~ subthreads via the GO RUNTIME SCHEDULER

    - MAIN THREAD > LOGICAL PROCESSOR > GORUNTIME SCHEDULER > { GOROUTINE1, GOROUTINE2, GOROUTINE3}
        - we can also request multiple logical processors depending on how many cores you have in your CPU; 6 cores => 6 logical processors; still scheduled; not parallel

### INTERLEAVINGS 
    - In running task A instructions(1,2,3) and task B instructions(1,2,3), CPU may complete tA i1,2 ==> tB i1 ==> tA i3 
    - Interleaving is non-deterministic: may not even finish all instructions in a task before switching to another task (INTERLEAVING) 

### RACE CONDITIONS
    - A result of all these interleavings may produce non-deterministic results == bad!
    - we want our program to always produce deterministic results
    

