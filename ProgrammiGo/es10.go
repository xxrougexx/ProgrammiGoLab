package main
import "fmt"

func main(){
  var h, min int

  for {
    fmt.Print("ore e minuti: ")
    fmt.Scan(&h,&min)
    if (h >= 0 && h <= 23) && (min >= 0 && min <= 59) {
        fmt.Println("valori validi")
        break
    }
  }
}
