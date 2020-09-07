package main

type LRU struct {
	size int
	list *CDLL
	table map[string] *Node
}

func NewLRU(sz int) *LRU {
	return &LRU{
		sz,
		&CDLL{sz, 0, nil},
		make(map[string] *Node),
	}
}

func (lru *LRU) insertKeyValuePair(k string, v int) {
	var cur *Node
	if _, ok := lru.table[k]; ok {
		cur = lru.table[k]
		cur.Value = v
		lru.list.UpdateHead(cur)
	} else {
		cur = lru.list.AddAtHead(k, v)
		lru.table[k] = cur
		if lru.list.Len > lru.list.Cap {
			lru.list.Len--
			lru.list.Head.Prev.Prev.Next = lru.list.Head
			delete(lru.table, lru.list.Head.Prev.Key)
			lru.list.Head.Prev = lru.list.Head.Prev.Prev
		}
	}
}

func (lru *LRU) getMostRecentKey() interface{} {
	if lru.list.Len > 0 {
		return lru.list.Head.Key
	}
	return nil
}

func (lru *LRU) getValueFromKey(k string) interface{} {
	if cur, ok := lru.table[k]; ok {
		lru.list.UpdateHead(cur)
		return cur.Value
	}
	return nil
}