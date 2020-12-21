package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var data []string
	var input string
	var mese, giorni int
	fmt.Print("Inserire data (formato gg-mm-aaaa): ")
	fmt.Scan(&input)
	data = strings.Split(input, "-")
	mese, _ = strconv.Atoi(data[1])
	giorni = giorniInMese(mese)
	fmt.Printf("Ci sono %d giorni nel mese inserito ", giorni)
}

func giorniInMese(mese int) (giorni int) {
	switch mese {
	case 2:
		giorni = 28
	case 11, 4, 6, 8:
		giorni = 30
	default:
		giorni = 31
	}
	return
}
