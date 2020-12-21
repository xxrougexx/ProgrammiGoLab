package main
import "fmt"

func main() {
  var num, numx0x int
  flag := "non ha"

  fmt.Print("Inserire numero a 3 cifre: ")
  fmt.Scan(&num)

  numx0x = num / 10
  if num % 10 == 0 || numx0x % 10 == 0 {
    flag = "ha"
  }

  fmt.Println("Il numero inserito", flag, "degli zeri")
}
