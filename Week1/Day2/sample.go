package main

import (
	"fmt"
	"strconv"
)

func main() {
	messages := make(chan string)

	//var wg sync.WaitGroup
	go func() {
		for i := 0; i < 100; i++ {
			//wg.Add(1)
			//go func(localI int) {
			//defer wg.Done()
			messages <- strconv.Itoa(i)
			//}(i)

		}
		//wg.Wait()
		close(messages)
	}()

	for i := range messages {
		fmt.Println(i)
	}

	//go func() { messages <- "str" }()
	//go func() { messages <- "str" }()
	//time.Sleep(time.Second)
	//for i := range messages {
	//	fmt.Println(i)
	//}
	//close(messages)
	//message := <-messages
	//fmt.Print(message)
}
