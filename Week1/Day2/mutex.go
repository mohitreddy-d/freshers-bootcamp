package main

import (
	"fmt"
	"sync"
)

type Node struct {
	m   sync.Mutex
	val int
}

func (node *Node) inc() {
	//for i := 0; i < n; i++ {
	//	node.val++
	//}
	//node.m.Lock()
	//defer node.m.Unlock()
	node.val++
}

func doInc(n int, node *Node, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		node.inc()
	}
}

func main() {

	node := Node{
		val: 0,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go doInc(100000, &node, &wg)
	go doInc(100000, &node, &wg)
	go doInc(100000, &node, &wg)

	wg.Wait()
	//time.Sleep(2 * time.Second)
	fmt.Print(node.val)

}
