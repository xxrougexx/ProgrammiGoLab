package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Salutare n volte ")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Println("Ciao")
	}
}
