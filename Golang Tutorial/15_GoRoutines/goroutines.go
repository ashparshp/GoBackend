package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id)
}

func main() {

	var wg sync.WaitGroup


	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}