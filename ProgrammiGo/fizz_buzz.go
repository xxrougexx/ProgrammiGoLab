package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%3 == 0 {
		s = "Fizz "
	}
	if n%5 == 0 {
		s += "Buzz"
	}
	fmt.Println(s)
}
