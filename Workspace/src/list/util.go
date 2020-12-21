package list

//Ritorna la lunghezza della lista inserita
func (x List) Length() int {
  c := 0
  for x != nil {
    x = x.Next
    c++
  }
  return c
}

//Cancella un nodo della posizione desiderata
//dalla lista
func (x List) DeleteNode(pos int) (bool, List) {
  var curr, prev List
  var ok bool
  curr = x
  for i := 0; i < pos - 1; i++ {
    prev = curr
    curr = curr.Next
    if curr.Next == nil {
      ok = false
    }
  }
  prev.Next = curr.Next
  curr.Next = nil
  return ok, x
}
