package main

import "fmt"

// Sending
/* 2
func processChan(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processng...", num)
		time.Sleep(time.Second)
	}

	// for a single value...
	// fmt.Println("Processng...", <-numChan)
}
*/

// Receive 3
/*
func sum(result chan int, num1, num2 int) {
	numResult := num1 + num2

	result <- numResult
}
*/

/* 4
func task(done chan bool) {
	defer func () {done <- true}()
	fmt.Println("Process...")
}
*/

/*
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func () {done <- true}()
	for email := range emailChan {
		fmt.Println("Sending email to: ", email)
		time.Sleep(time.Second)
	}
}
*/

func main() {


	/* 1
	message := make(chan string)
	message <- "ping" // Channels are blocking unless other side ready to receive

	msg := <-message
	fmt.Println(msg)
	*/

	/* 2
	numChannel := make(chan int)

	go processChan(numChannel)

	// numChannel <- 5

	for {
		numChannel <- rand.Intn(100)
	}
	*/

	/* 3
	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result // blocking
	fmt.Println(res)
	*/

	/* 4
	done := make(chan bool)

	go task(done)

	<- done
	*/

	/*
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 1; i <= 5; i++ {
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}
	fmt.Println("Done sending!")

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	close(emailChan) // important
	<-done
	*/

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func ()  {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "P..."
	}()

	for i := 0; i <2; i ++ {
		select {
		case chan1Vale := <- chan1:
			fmt.Println("Received data from chan1", chan1Vale)
		case chan2Vale := <- chan2:
			fmt.Println("Received data from chan2", chan2Vale)
		}
	}

}