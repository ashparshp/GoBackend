package main

func main() {
	message := make(chan string)
	message <- "ping"

	msg := <-message
	println(msg)

}