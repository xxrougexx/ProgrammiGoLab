package main

import (
  "bufio"
  "errors"
  "os"
  "regexp"
  "strconv"
)

func parseLine(line string) (stud student, err error) {
  nameRE := regexp.MustCompile("([a-z]|[A-Z])([a-z]|[A-Z]|[0-9]|_)+")
  scoreRE := regexp.MustCompile("[1-3][0-9]")
  //Finds first matching substring
  stud.Name = nameRE.FindString(line)
  if len(stud.Name) == 0 {
    err = errors.New("Name of student is missing in line " + line)
    return
  }
  //Finds the 2 indexes delimiting the first matching string
  nameEndingIndex := nameRE.FindStringIndex(line)[1]
  scores := scoreRE.FindAllString(line[nameEndingIndex:], -1)
  for _, x := range scores {
    score, conv := strconv.Atoi(x)
    if score < 18 || score > 32 || score == 31 {
      err = errors.New("The score " + x + " is invalid in line " + line)
      return
    } else if conv != nil {
      err = conv
      return
    }
    stud.Score = append(stud.Score, score)
  }
  if len(stud.Score) == 0 {
    err = errors.New("No scores for student " + stud.Name)
    return
  }
  return
}

func readFile(fileName string) (stud []student, err error) {
  f, e := os.Open(fileName)
  if e != nil {
    err = e
    return
  }
  defer f.Close()
  b := bufio.NewScanner(f)
  for b.Scan() {
    line := b.Text()
    s, e := parseLine(line)
    if e != nil {
      err = e
      return
    }
    stud = append(stud, s)
  }
  return
}
