package main

import "fmt"

func main() {
	s := []string{"A", "B", "C"}

	fmt.Println("Hello World")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for ix, l := range s {
		fmt.Println("Index: ", ix, "Letter:", l)
	}


}

