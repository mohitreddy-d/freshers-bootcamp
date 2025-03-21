package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func getData(w http.ResponseWriter, req *http.Request) {
	url := "https://randomuser.me/api/"
	res, err := http.Get(url)
	w.Header().Add("content-type", "application/json")
	scanner := bufio.NewScanner(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for i := 0; scanner.Scan() && i < 2; i++ {
		w.Write([]byte(scanner.Text()))
	}

}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user", getData)
	http.ListenAndServe(":8090", nil)
}
