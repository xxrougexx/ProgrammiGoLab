package list

import "fmt"

//Aggiunge un nodo all'inizio della lista.
//Ritorna la nuova lista
func (x List) AddFront(s string) List {
  var n List
  n = new(Node)
  n.Content = s
  n.Next = x
  return n
}

//Aggiunge un nodo alla fine della lista.
//Ritorna la nuova lista
func (x List) AddTail(s string) List {
  var n List
  n = new(Node)
  n.Content = s
  n.Next = nil
  if x == nil {
    return n
  }
  first := x
  for x.Next != nil {
    x = x.Next
  }
  x.Next = n
  return first
}
//Aggiunge un nodo in modo che il suo contenuto sia
//in ordine alfabetico col precedente e successivo
func (x List) AddInOrder(s string) List {
  var n List
  n = new(Node)
  n.Content = s
  var prev, curr List
  prev = nil
  curr = x
  for curr != nil && curr.Content < s {
    prev = curr
    curr = curr.Next
  }
  //Ora lo inseriamo tra prev e curr
  if prev == nil {
    n.Next = curr
    return n
  }
  prev.Next = n
  n.Next = curr
  return x
}
//Aggiunge un nodo alla posizione indicata da input
//(0-prima posizione)
func (x List) AddInPos(s string, pos int) (bool, List) {
  var n List
  var ok bool
  n = new(Node)
  n.Content = s
  var prev, curr List
  curr = x
  for i := 0; i < pos; i++ {
      prev = curr
      curr = curr.Next
      if curr.Next == nil {
        fmt.Println("Errore; posizione inserita supera lunghezza lista attuale")
        break
      }
  }
  prev.Next = n
  n.Next = curr
  return ok, x
}
