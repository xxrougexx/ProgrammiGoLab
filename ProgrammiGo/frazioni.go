package main
import "fmt"

func main(){
  var num1, num2, den1, den2 int

  fmt.Print("num e den fraz 1: ")
  fmt.Scan(&num1,&den1)
  fmt.Print("num e den fraz 2: ")
  fmt.Scan(&num2,&den2)

  incrocio1 := num1 * den2
  incrocio2 := num2 * den1

  if incrocio1 == incrocio2 {
    fmt.Println("equivalenti")
  } else {
    fmt.Println("non equivalenti")
  }
}
