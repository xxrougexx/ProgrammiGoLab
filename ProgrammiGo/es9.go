package main
import "fmt"
import "math/rand"

func main(){
  var somma int
  const (
    K = 10
    TARGET = 50
    SEED = 10
  )
  rand.Seed(SEED)

  for somma < TARGET{
    somma += rand.Intn(10)
  }
  fmt.Println(somma)
}
