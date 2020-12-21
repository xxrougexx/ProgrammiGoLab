package check


/* Checks if the string s is palindrome, and
returns the corresponding boolean value*/
func RecIsPalyn(s string) bool {
  if len(s) < 2 {
    return false
  }
  return s[0]==s[len(s) - 1] && RecIsPalyn(s[1:len(s)-1])
}
