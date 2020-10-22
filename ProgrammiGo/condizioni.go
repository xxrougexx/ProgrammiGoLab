package main
import "fmt"

func main(){
  var a int
  var x float64
  fmt.Print("Inserisci intero uguale a 10: ")
  fmt.Scan(&a)
  fmt.Println(a == 10)

  fmt.Print("Inserisci intero diverso da 10: ")
  fmt.Scan(&a)
  fmt.Println(a != 10)

  fmt.Print("Inserisci intero diverso da 10 e da 20: ")
  fmt.Scan(&a)
  fmt.Println(a != 10 && a != 20)

  fmt.Print("Inserisci intero diverso da 10 o da 20: ")
  fmt.Scan(&a)
  fmt.Println(a != 10 || a != 20)

  fmt.Print("Inserisci intero maggiore o uguale a 10: ")
  fmt.Scan(&a)
  fmt.Println(a >= 10)

  fmt.Print("Inserisci intero compreso tra [10 e 20]: ")
  fmt.Scan(&a)
  fmt.Println(a >= 10 && a <= 20)

  fmt.Print("Inserisci intero compreso tra (10 e 20): ")
  fmt.Scan(&a)
  fmt.Println(a > 10 && a < 20)

  fmt.Print("Inserisci intero compreso tra [10 e 20): ")
  fmt.Scan(&a)
  fmt.Println(a >= 10 && a < 20)

  fmt.Print("Inserisci intero fuori dall'intervallo tra [10 e 20]: ")
  fmt.Scan(&a)
  fmt.Println(!(a >= 10 && a <= 20))

  fmt.Print("Inserisci intero compreso tra [10 e 20] o tra [30 e 40]: ")
  fmt.Scan(&a)
  fmt.Println((a >= 10 && a <= 20) || (a >= 30 && a <= 40))

  fmt.Print("Inserisci intero multiplo di 4 ma non di 100: ")
  fmt.Scan(&a)
  fmt.Println(a % 4 == 0 && a % 100 != 0)

  fmt.Print("Inserisci intero dispari compreso tra 0 e 100: ")
  fmt.Scan(&a)
  fmt.Println(a % 2 != 0 && a > 0 && a <= 100)

  fmt.Print("Inserisci float vicino a 10.0 (entro 10^-6): ")
  fmt.Scan(&x)
  fmt.Println(x <= 10 + 10e-6 && x >= 10 - 10e-6)
}
