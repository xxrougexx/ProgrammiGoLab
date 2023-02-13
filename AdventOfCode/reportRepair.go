package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  var report []int
  var sum2020, num1, num2 int
  b := bufio.NewScanner(os.Stdin)
  for b.Scan() {
    val, _ := strconv.Atoi(b.Text())
    report = append(report, val)
  }

  for i := 0; i <= len(report) - 1; i++ {
    for j := 0; j <= len(report) - 1; j++ {
      num1, num2 = report[i], report[j]
      sum2020 = s2020(num1, num2, report)
      if sum2020 != -1 {
        break
      }
    }
    if sum2020 != -1 {
      break
    }
  }
  fmt.Println(sum2020)
}

func s2020(n1, n2 int, x []int) int {
  for i := range x {
    if x[i] != n1 && x[i] != n2 && x[i] + n1 + n2 == 2020 {
      return x[i] * n1 * n2
    }
  }
  return -1
}
