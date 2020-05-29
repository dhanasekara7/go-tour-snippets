package main

import "fmt"

func main() {
	fmt.Printf("%v\n" , add(10, 20))
	fmt.Printf("%v\n" , add2(10, 20))
}


func add(x int, y int) int {
	return x + y
}

// x,y int
func add2(x,y int) int {
	return x + y
}