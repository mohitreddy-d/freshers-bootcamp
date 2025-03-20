package main

import (
	"fmt"
	"sync"
)

func print(s string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for i := 1; i <= 100; i++ {
		fmt.Println(s, i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go print("first", &wg)

	//print("====", nil)
	go print("second", &wg)
	//time.Sleep(time.Second)
	//fmt.Println(wg)
	wg.Wait()
}
