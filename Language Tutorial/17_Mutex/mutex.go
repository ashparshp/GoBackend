package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p * post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	defer p.mu.Unlock()
	p.mu.Lock()
	p.views += 1
}


func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	// myPost.inc() // views happens concurrently not one user wait for other
	// myPost.inc()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}