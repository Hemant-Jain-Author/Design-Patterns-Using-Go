package main

import (
	"fmt"
	"sync"
)

// Barrier represents a synchronization barrier
type Barrier struct {
	count       int
	barrierLock sync.Mutex
	barrierCond *sync.Cond
}

// NewBarrier creates a new Barrier
func NewBarrier(count int) *Barrier {
	b := &Barrier{
		count:       count,
		barrierCond: sync.NewCond(&sync.Mutex{}),
	}
	return b
}

// Wait waits for all goroutines to reach the barrier
func (b *Barrier) Wait() {
	b.barrierLock.Lock()
	defer b.barrierLock.Unlock()

	b.count--

	if b.count > 0 {
		b.barrierCond.Wait()
	} else {
		b.barrierCond.Broadcast()
	}
}

func worker(b *Barrier, id int) {
	fmt.Printf("Worker %d started\n", id)
	// Simulating some work
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d working...\n", id)
		// Simulating some computation
	}
	fmt.Printf("Worker %d finished\n", id)
	b.Wait()
}

func main() {
	// Create a barrier with the specified number of workers
	numWorkers := 3
	barrier := NewBarrier(numWorkers)

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(barrier, i)
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers finished. Proceeding to the next step.")
}
