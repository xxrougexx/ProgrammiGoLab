package main

import "fmt"

func main() {
	const K = 5
	var parola string
	var vocale bool
	for i := 0; i < K; i++ {
		//fmt.Println()
		fmt.Scan(&parola)
		vocale = false
		for _, c := range parola {
			switch c {
			case 'a', 'e', 'i', 'o', 'u':
				fmt.Printf("%c\n", c)
				vocale = true
				break
			}
		}
		if vocale == false {
			fmt.Println("la parola non contiene vocali")
		}
	}
}
