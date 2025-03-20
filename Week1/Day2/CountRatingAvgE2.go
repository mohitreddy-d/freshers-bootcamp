package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takeRating(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(time.Second))
	ch <- rand.Intn(10) + 1
}

func main() {
	noOfStudents := 200
	ch := make(chan int)
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < noOfStudents; i++ {
		wg.Add(1)
		go takeRating(ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	var avg int
	sum := 0
	for rating := range ch {
		fmt.Println(rating)
		sum += rating
	}
	avg = sum / noOfStudents

	fmt.Print("average rating is:", avg)

}
