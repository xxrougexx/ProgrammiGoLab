package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)
func main() {
  s, err := readFile(os.Args[1])
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  for _, x := range s {
    fmt.Printf("%-20s\t%.2f\t%6.2f", x.Name, average(x), average110(x))
    for _, sc := range x.Score {
      fmt.Printf("\t%2d", sc)
    }
    fmt.Println()
  }
  fmt.Println("Copying the data to " + os.Args[2])
  b, _ := json.Marshal(s)
  //0644 is the permission flag on the file
  ioutil.WriteFile(os.Args[2], b, 0644)
  fmt.Println("Now i read back the json file")
  b, _ = ioutil.ReadFile(os.Args[2])
  var t []student
  json.Unmarshal(b, &t)
  fmt.Println(t)
}
