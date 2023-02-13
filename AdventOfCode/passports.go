package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "regexp"
)

func main() {
  var count int

  param := make(map[string]string)
  b := bufio.NewScanner(os.Stdin)
  b.Split(bufio.ScanWords)

  for b.Scan() {
    field := strings.Split(b.Text(), ":")
    param[field[0]] = field[1]
    fmt.Println(param)
    count += countPassports(param)
    if count != 0 {
      fmt.Println(count)
    }
  }
  fmt.Println(count)
}


func countPassports(m map[string]string) (valid int) {
  v := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
  if len(m) == 8 && v.MatchString(m["byr"]) && v.MatchString(m["iyr"]) && v.MatchString(m["eyr"]) && v.MatchString(m["hgt"]) && v.MatchString(m["hcl"]) && v.MatchString(m["ecl"]) && v.MatchString(m["pid"]){
    valid++
    for k := range m {
      delete(m, k)
    }
  }
  return valid
}
