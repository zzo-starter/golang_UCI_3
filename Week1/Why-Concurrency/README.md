## WEEK1: Why Concurrency? 
- Concurrency is BUILT into Go vs C or Python; can import libraries
- In accademia; everyone learns sequential programming; C, Python, Java, C#; not parallel or not concurrent
- Given todays new CPU architecture with multiple cores (ie this MacMini i7-6core ); wasting those cores and their potential bc apps just use 1 core

- Powerpoint slides 
    ### Parallel Execution
    >When 2 programs execute at EXACTLY the same time; thus saving overall time; increasing bandwidth/throughput
    - `X: 1 vs 2 dishwashers`
    - processor cores (typically) exec 1 instruction at a time 
    - to run in parallell need (2) replicated physical CPUS running that app 
    - writing code for parallelism is `HARD`; manually managing context/memory/states/hardware...
    - to speed up exec, click speeds would be bumped up or even `overclocked`
    - new 'fasster' CPUs continually released every 6 months (Moores Law; transistor density doubled every 2 years; smaller=faster); but irrelevant now

    ### Von Neumann Bottleneck (MEMORY SPEED)
    - Delayed access memory; CPU continually accessing memory for writing/reading code and state, clearing, managing, GC... and memory is `VERY slow`; slower than CPU
        - one solution was `Cache`; CPU Cache to speed up access and rely less on memory access for 
        - Dynamic Power; `P = a * CFV^2`
            - a = % of time switching
            - C = capacitance (capacitor)
            - F = clock frequency
            - V = voltage swing from low to high 

    ### Power Wall (HEAT)
    - Increased transistors consume more power => increase in temperature => cannot cool enough with simple heatsinks and fans; else melt chip
    - therefore;fully maximizing throughput via concurrency (1CPU:n cores) is the answer; sp for laptops 
    
    ### Multi-Core 
    - CPUs must sell; therefore manuf increases the # of cores; not really the frequency
    - BUT parallelism is need to fully exploit multi-core systems; and extremely difficult 

    ### Concurrent vs Parallel
    - 
    ### Goroutine race conditions 
    - a timing bug
    - 2 different routines interacting w same data at same time 

    ### Hiding Latency
    - Tasks must periodically WAIT for some transaction
        - X: wait for memory
    - Other tasks can operate while one task is waiting 

- Assessments
    - 1 Quiz
    - 2 Activities 