# Advent of Code 2023

Last year I decided I would try advent of code, solved the first two days and then sort of gave up.

This year will be different™️!

## Note the odd use of `pacakage main`

All cmd/puzzleXX directories are their own `main` package. Luckily Go is smart (apparently) and you can just cd into those directories and `go run .` and everything is scoped to that directory. I was a little worried that puzzle01::parseLine would leak into the main package of puzzle02, but it doesn't.

I'm pretty sure this is not how you're supposed to do things but the point was to learn, and this works for now. I'll refactor if I have to but it's AoC, not a production application.

## Stuff I learned

### Go doesn't have a while loop

Except it sort of does:
```go
x := 0
for x < 5 {
  fmt.Println(x)
  x++
}
```

### Goroutines are really easy to use if you want to distribute looped work
The trick when spawning multiple threads is to use waitgroups 

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func processElement(waitTime int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done() // `defer` means wg.Done() gets called when the function exits

	time.Sleep(time.Duration(waitTime) * time.Second)
	resultChan <- waitTime
}

func main() {
	resultChan := make(chan int)
	var wg sync.WaitGroup

	// This behaves like a loop of sleep(time.Second) calls but is actually parallel
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processElement(i, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		// This will loop until you close the channel - which happens when the wg is done
		fmt.Println(result)
	}
}
```
