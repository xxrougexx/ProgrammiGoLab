package main


import "fmt"

func main() {
	var numero int
	fmt.Scan(&numero)
	max := numero
	count := 1
	for i := 1; i < 10; i++ {
		fmt.Scan(&numero)
		if numero == max {
			count++
		} else if numero > max {
			max = numero
			count = 1
		}
	}
	fmt.Println(max)
	fmt.Println(count)
}
