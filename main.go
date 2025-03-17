package main

import "fmt"
// import "strconv"

func main() {
	
	// number, err := strconv.Atoi("200o")
	// if(err != nil) {
	// 	fmt.Println(err)
	// }
	n := 5
	var arr1, arr2 []int
	arr1 = make([]int, n)
	arr2 = make([]int, n)
	for i := 0; i < n; i++ {
		arr1[i] = i
		arr2[i] = n - i - 1
	}
	
	fmt.Println("hello world", arr1, arr2)
}