package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	url := "https://randomuser.me/api/"
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Fprint(scanner.Text())
	}
}
