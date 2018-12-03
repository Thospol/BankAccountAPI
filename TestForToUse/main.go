package main

import "fmt"

func main() {
	a := []string{"Hello1", "Hello2", "Hello3"}
	// a = append(a[:0], a[1:] )

	var copy []string

	for _, element := range a {
		if element != "Hello2" {
			copy = append(copy, element)
		}

	}

	fmt.Println(len(copy))
	fmt.Println(copy)

}
