package list

import "fmt"

//Fa la print di tutti gli elementi della lista
func (x List) PrintAllElements() {
  for x != nil {
    fmt.Println(x.Content)
    x = x.Next
  }
}

//Fa la print di tutti gli elementi della lista
//con le frecce
func (x List) PrintWithArrows() {
  for x != nil {
    fmt.Printf("%s", x.Content)
    x = x.Next
    if x != nil {
      fmt.Print(" -> ")
    }
  }
  fmt.Println()
}

//Fa la print di tutti gli elementi della lista
//in modo ricorsivo
func (x List) PrintRecursive() {
  fmt.Println(x.Content)
  x = x.Next
  if x == nil {
    return
  }
  x.PrintRecursive()
}

//Fa la print di tutti gli elementi della lista
//in modo ricorsivo con le frecce
func (x List) PrintRecursiveWithArrows() {
  fmt.Printf("%s", x.Content)
  x = x.Next
  if x != nil {
    fmt.Print(" -> ")
  } else {
    return
  }
  x.PrintRecursiveWithArrows()
}
