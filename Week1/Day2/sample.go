package main

import "fmt"

func main() {

	messages := make(chan string, 1)

	messages <- "buffered"
	fmt.Println(<-messages)
	messages <- "channel"

	fmt.Println(<-messages)
}
