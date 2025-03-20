package main

import (
	"fmt"
	"sync"
	"time"
)

func printCh(ch chan int) {

	for val := range ch {
		fmt.Print(val)
	}

}

func addNum(i int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- i

}

func main() {
	ch := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6}
	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go addNum(num, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	//len(ch)
	for i := range ch {
		fmt.Print(i, " ")
	}
	//go printCh(ch)
	time.Sleep(time.Second)

}
