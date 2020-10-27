package main

import "fmt"

func main() {
	var voto int
	var votoS string
	fmt.Print("Voto: ")
	fmt.Scan(&voto)
	if voto < 0 || voto > 100 {
		votoS = "non valido"
	} else {
		if voto > 0 && voto < 60 {
			votoS = "F"
		} else if voto >= 60 && voto < 70 {
			votoS = "D"
		} else if voto >= 70 && voto < 80 {
			votoS = "C"
		} else if voto >= 80 && voto < 90 {
			votoS = "B"
		} else {
      votoS = "A"
    }
	}
  fmt.Println(votoS)
}
