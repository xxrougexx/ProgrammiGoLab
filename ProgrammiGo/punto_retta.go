package main
import "fmt"
import "math"

func main(){
  var x1, y1, m, q float64

  fmt.Print("Inserire coordinate del punto(x y): ")
  fmt.Scan(&x1,&y1)
  fmt.Print("Inserire m e q della retta(formula: y = mx + q): ")
  fmt.Scan(&m,&q)

  if math.Abs(y1 - (m * x1 + q)) < 10e-6 {
    fmt.Println("sulla retta")
  } else if y1 < m * x1 + q {
    fmt.Println("sotto")
  } else {
    fmt.Println("sopra")
  }
}
