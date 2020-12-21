package check


/* Checks if the string s is palindrome, and
returns the corresponding boolean value*/
func IterIsPalyn(s string) bool {
  n := len(s)
  for i := 0;i < n / 2;i++ {
    if s[i] != s[n - i - 1] {
      return false
    }
  }
  return true
}
