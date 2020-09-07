package main

import "fmt"

type Node struct {
	Key string
	Value int
	Prev *Node
	Next *Node
}

type CDLL struct {
	Cap int
	Len int
	Head *Node
}

func (l *CDLL) AddAtHead(k string, v int) *Node {
	l.Len++
	cur := &Node{k, v, nil, nil}
	if l.Head == nil {
		l.Head = cur
		cur.Next = cur
		cur.Prev = cur
		return cur
	}
	cur.Next = l.Head
	cur.Prev = l.Head.Prev
	l.Head.Prev.Next = cur
	l.Head.Prev = cur
	l.Head = cur
	return cur
}

func (l *CDLL) UpdateHead(cur *Node) {
	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev
	cur.Next = l.Head
	cur.Prev = l.Head.Prev
	l.Head.Prev.Next = cur
	l.Head.Prev = cur
	l.Head = cur
}

func (l *CDLL) PrintList() {
	fmt.Println("printing list with length ", l.Len)
	fmt.Printf("%s: %d\n", l.Head.Key, l.Head.Value)
	cur := l.Head.Next
	for cur != l.Head {
		fmt.Printf("%s: %d\n", cur.Key, cur.Value)
		cur = cur.Next
	}
	fmt.Println("done")
}