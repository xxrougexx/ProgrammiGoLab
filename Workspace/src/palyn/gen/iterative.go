package gen

import(
  "math/rand"
)

func genCharacter() string {
  return string('a' + rune(rand.Intn('z' - 'a' + 1)))
}

func IterativeGenPalyn(n int) string {
  t := ""
  if n % 2 != 0 {
    t = string(genCharacter())
  }
  for i := 0;i < n/2; n++{
    c := genCharacter()
    t = c + t + c
  }
  return t
}
