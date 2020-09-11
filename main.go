package main

import "fmt"

func main() {
	lru := NewLRU(3)
	lru.insertKeyValuePair("b", 2)
	lru.insertKeyValuePair("a", 1)
	lru.insertKeyValuePair("c", 3)
	fmt.Println(lru.getMostRecentKey())
	fmt.Println(lru.getValueFromKey("a"))
	fmt.Println(lru.getMostRecentKey())
	lru.insertKeyValuePair("d", 4)
	fmt.Println(lru.getValueFromKey("b"))
	lru.insertKeyValuePair("a", 5.1)
	fmt.Println(lru.getValueFromKey("a"))
}