package main
import "fmt"
import "math/rand"

func main(){
  const (
    K = 10
    SEED = 10
  )
  rand.Seed(SEED)

  for i := 0; i < K; i++{
    fmt.Println(rand.Intn(K))
  }
}
