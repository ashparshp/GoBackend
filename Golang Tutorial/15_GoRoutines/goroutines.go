package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Task", id)
}

func main() {
	for i := 0; i < 10; i++ {
		go task(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Done")
}