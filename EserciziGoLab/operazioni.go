package main

import "fmt"

func main() {
	var val1, val2 int
	fmt.Print("Inserire due valori: ")
	fmt.Scan(&val1, &val2)
	s, p, d := operazioni(val1, val2)
	fmt.Println(s, p, d)
}

func operazioni(n1 int, n2 int) (int, int, int) {
	somma := n1 + n2
	prodotto := n1 * n2
	differenza := n1 - n2
	return somma, prodotto, differenza
}

/*func operazioni(n1 int, n2 int) (somma int, prodotto int, differenza int) {
  somma = n1 + n2
  prodotto = n1 * n2
  differenza = n1 - n2
  return
}*/
