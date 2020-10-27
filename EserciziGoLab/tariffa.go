package main

import "fmt"

func main() {
	var età int
	var stud bool
  tariffa := "5 euro"

	fmt.Print("età: ")
	fmt.Scan(&età)
	fmt.Print("studente? (t/f) ")
	fmt.Scan(&stud)

	if stud == false{
		if età >= 0 && età < 9 {
			tariffa = "gratis"
		} else if età >= 18 && età < 26 || età >= 65 {
			tariffa = "7.5 euro"
		} else if età >= 26 && età < 65 {
			tariffa = "10 euro"
    }
	}
  fmt.Println("ingresso",tariffa)
}
