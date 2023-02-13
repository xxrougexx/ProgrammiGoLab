package main

import (
  "fmt"
  "math/rand"
  "os"
  "bufio"
  "strconv"
)

func main() {
  strNumAmici := os.Args[1]
  numAmici, _ := strconv.Atoi(strNumAmici)
  amico := make([]int, numAmici + 1)


  for i := 1; i <= numAmici; i++ {
    amico[i] = i
  }

  f, err := os.Create("estrazione.txt")
  if err != nil {
    panic(err)
  }
  w := bufio.NewWriter(f)

  defer f.Close()

  for n := numAmici; n > 0; n-- {
    pos := rand.Intn(n)
    amico[pos], amico[n-1] = amico[n-1], amico[pos]

    if amico[n] == n {
      if pos == 0 {
        pos++
      }
      amico[pos], amico[n] = amico[n], amico[pos]
    }
    if amico[n] == 0 {
      amico[n] = numAmici
    }
    fmt.Println(n, "->", amico[n])
    fmt.Fprintf(w,"%d -> %d\n", n,amico[n])
  }
  w.Flush()
}
