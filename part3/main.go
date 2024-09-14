package main

import "fmt"

func main() {
	people := map[string]int{"bob": 18, "dobi": 31}

	delete(people, "bob")

	for k, v := range people {
		fmt.Printf("name:%s age:%d\n", k, v)
	}
}
