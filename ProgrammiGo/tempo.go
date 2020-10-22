package main

import "fmt"

func main() {
  var numSecondi int
  const (
    SEC_MINUTO = 60
    SEC_ORA = SEC_MINUTO * 60
    SEC_GIORNO = SEC_ORA * 24
  )

  fmt.Print("Numero di secondi: ")
  fmt.Scan(&numSecondi)

  resto_sec := numSecondi % SEC_GIORNO
  resto_min := resto_sec % SEC_ORA
  giorni := numSecondi / SEC_GIORNO
  ore := resto_sec / SEC_ORA
  resto_sec = numSecondi % SEC_ORA
  minuti := resto_sec / SEC_MINUTO
  secondi := resto_min % SEC_MINUTO

  fmt.Println("g:h:m:s =", giorni, ":", ore, ":",
              minuti, ":", secondi)

}
