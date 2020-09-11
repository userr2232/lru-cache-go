package main

type Node struct {
	Key interface{}
	Value interface{}
	Prev *Node
	Next *Node
}

// Circular Double Linked List
type CDLL struct {
	Cap int
	Len int
	Head *Node
}

// Sirve para enviar un nodo no Head adelante
// es decir, volverlo Head O(1)
func (l *CDLL) placeFront(cur *Node) {
	cur.Next = l.Head
	cur.Prev = l.Head.Prev
	l.Head.Prev.Next = cur
	l.Head.Prev = cur
	l.Head = cur
}

// Sirve para agregar un nuevo nodo a la CDLL como Head O(1)
func (l *CDLL) AddAtHead(k interface{}, v interface{}) *Node {
	l.Len++
	cur := &Node{k, v, nil, nil}
	if l.Head == nil {
		l.Head = cur
		cur.Next = cur
		cur.Prev = cur
		return cur
	}
	l.placeFront(cur)
	return cur
}

// Remueve un nodo de la CDLL O(1) manteniéndola conectada
// y luego posiciona el nodo removido como Head O(1)
func (l *CDLL) UpdateHead(cur *Node) {
	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev
	l.placeFront(cur)
}