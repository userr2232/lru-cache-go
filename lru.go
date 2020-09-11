package main

type LRU struct {
	size int
	list *CDLL
	table map[interface{}] *Node
}

func NewLRU(sz int) *LRU {
	return &LRU{
		sz,
		&CDLL{sz, 0, nil},
		make(map[interface{}] *Node),
	}
}

func (lru *LRU) insertKeyValuePair(k interface{}, v interface{}) {
	var cur *Node
	if _, ok := lru.table[k]; ok {
		// si el key está en la hash table, se extrae el value del nodo
		// y se actualiza la posición en la cache
		cur = lru.table[k]
		cur.Value = v
		lru.list.UpdateHead(cur)
	} else {
		// si no está en la tabla, se agrega un nodo a la CDLL O(1) y a la hash table O(1)
		cur = lru.list.AddAtHead(k, v)
		lru.table[k] = cur
		// Si la longitud de la CDLL excede la capacidad, 
		// se disminuye en 1 la longitud
		// eliminando el nodo final de la CDLL (el prev al head) O(1) y de la hash table O(1)
		// de esta forma se asegura que sea O(1) en espacio
		if lru.list.Len > lru.list.Cap {
			lru.list.Len--
			lru.list.Head.Prev.Prev.Next = lru.list.Head
			delete(lru.table, lru.list.Head.Prev.Key)
			lru.list.Head.Prev = lru.list.Head.Prev.Prev
		}
	}
}

func (lru *LRU) getMostRecentKey() interface{} {
	// Si la CDLL tiene longitud > 0, se obtiene el Key del Head en O(1)
	if lru.list.Len > 0 {
		return lru.list.Head.Key
	}
	return nil
}

func (lru *LRU) getValueFromKey(k interface{}) interface{} {
	// Si la key existe en la hash table, se retorna el value del nodo en O(1)
	if cur, ok := lru.table[k]; ok {
		lru.list.UpdateHead(cur)
		return cur.Value
	}
	return nil
}