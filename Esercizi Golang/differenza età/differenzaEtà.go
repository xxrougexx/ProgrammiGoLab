package main
import "fmt"

func main() {
  var anno1, anno2, diffAnni int
  var flag string

  fmt.Println("Inserire primo anno di nascita: ")
  fmt.Scan(&anno1)
  fmt.Println("Inserire secondo anno di nascita: ")
  fmt.Scan(&anno2)

  if anno1 > anno2 {
    diffAnni = anno1 - anno2
    flag = "seconda"
  } else {
    diffAnni = anno2 - anno1
    flag = "prima"
  }

  fmt.Println("La", flag, "persona e' più vecchia di", diffAnni, "anni")
}
