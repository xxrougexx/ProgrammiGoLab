package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  var street string
  var trees, countStreet, start int
  b := bufio.NewScanner(os.Stdin)
  for b.Scan() {
    if countStreet % 2 != 0 {
      street = b.Text()
    } else {
      street = b.Text()
      if start >= len(street) {
        start -= len(street)
      }
      if street[start] == '#' {
        trees++
      }
      start += 1
    }
    countStreet++
  }
  fmt.Println(trees)
}

//1, 3 = 268
//1, 1 = 68
//1, 5 = 73
//1, 7 = 75
//2, 1 = 31 LETS GO!
