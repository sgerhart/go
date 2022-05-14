package main

import "fmt"

func main() {
	fmt.Println("Hello Arrays, What Is Going On?")

	grades := [3]int{87, 67, 97}
	gradesExample2 := [...]int{87, 67, 97}

	fmt.Printf("Grades %v\n", grades)
	fmt.Printf("Grades %v", gradesExample2)

}
