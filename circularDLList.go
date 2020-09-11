package main

type Node struct {
	Key interface{}
	Value interface{}
	Prev *Node
	Next *Node
}

type CDLL struct {
	Cap int
	Len int
	Head *Node
}

func (l *CDLL) placeFront(cur *Node) {
	cur.Next = l.Head
	cur.Prev = l.Head.Prev
	l.Head.Prev.Next = cur
	l.Head.Prev = cur
	l.Head = cur
}

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

func (l *CDLL) UpdateHead(cur *Node) {
	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev
	l.placeFront(cur)
}