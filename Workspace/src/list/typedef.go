//Pacchetto per la gestione e creazione di liste
package list

//Le liste sono costituite da nodi;
//ognuno ha un contenuto string
//e un puntatore che punta al prossimo nodo 
type Node struct {
  Content string
  Next *Node
}

type List = *Node
