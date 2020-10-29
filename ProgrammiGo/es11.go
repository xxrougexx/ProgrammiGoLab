package main
import "fmt"
import "math/rand"

func main(){
  var random, somma int
  const (
    K = 21
    SEED = 10
  )
  rand.Seed(SEED)

  for {
    random = rand.Intn(K)
    fmt.Println(random)
    somma += random
    if random == K-1 {
      break
    }
  }
  fmt.Println(somma)
}
