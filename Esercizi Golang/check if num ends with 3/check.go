package main
import "fmt"

func main() {
  var num, check int
  
  fmt.Println("Inserire cifra da controllare: ")
  fmt.Scan(&num)

  if num < 0 {
    num = -num
  }
  check = (num - 3) % 10

  if check == 0{
    fmt.Println("L'ultima cifra di questo numero e' 3")
  } else {
    fmt.Println("L'ultima cifra di questo numero non e' 3")
  }
}
