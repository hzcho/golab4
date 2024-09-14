package main

import "fmt"

func main() {
	people := map[string]int{"bob": 18, "dobi": 31}

	people["oleg"] = 89

	for k, v := range people {
		fmt.Printf("name:%s age:%d\n", k, v)
	}
}
