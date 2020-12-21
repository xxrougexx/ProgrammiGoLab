package main
import "fmt"

func main() {
  var num float64
  var arrotondato int
  fmt.Println("Inserire numero con la virgola: ")
  fmt.Scan(&num)

  arrotondato = int(num)

  fmt.Println("Il numero arrotondato e': ", arrotondato)

}
