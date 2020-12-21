package main

import "fmt"

func main() {
	var p int
	var i int
	mano := make([]string, 3)
	for i < 3 {
		fmt.Printf("Inserire la %d^ carta da briscola: ", i+1)
		fmt.Scan(&mano[i])
		i++
	}
	for j := 0; j < 3; j++ {
		p += punti(mano[j])
	}
	fmt.Printf("\nmano %s:%d punti", mano, p)
}

func punti(s string) int {
	var p int
	switch s {
	case "A":
		p = 11
	case "3":
		p = 10
	case "K":
		p = 4
	case "Q":
		p = 3
	case "J":
		p = 2
	case "7", "6", "5", "4", "2":
		p = 0
	default:
		punti = -1
	}
	return p
}
