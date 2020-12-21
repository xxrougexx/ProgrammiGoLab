package person

import "time"

func (p Data) Age() int {
  var t time.Time
  return t.Year() - p.BirthYear
}
