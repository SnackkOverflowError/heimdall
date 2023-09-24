package main

import (
	"heimdall/memtable/rbtree"
)

// go run main.go | dot -Tpng  > test.png && open test.png

func main() {
	rbtree.RBDisplayTest()
}
