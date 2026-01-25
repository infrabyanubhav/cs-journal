package main

import (
	"fmt"
	"sync"
)

func add(n *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*n = *n + 2
}

func sub(n *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*n = *n - 10
}

func main() {
	var wg sync.WaitGroup
	var num, counter int
	fmt.Scan(&num)
	counter = num
	for i := 0; i < 10000; i++ {
		num = counter
		wg.Add(2)
		go add(&num, &wg)
		go sub(&num, &wg)
		wg.Wait()
		fmt.Print(num)
	}

}

// add modifies the shared variable `num` concurrently with other goroutines.
// This causes a data race because the read-modify-write operation (*n = *n + 2)
// is not atomic and can interleave with other writes (like sub),
// leading to unpredictable results depending on scheduling.

// sub also modifies the shared variable concurrently with add.
// Without synchronization (mutex/atomic/channel), these concurrent writes
// create a race condition.
