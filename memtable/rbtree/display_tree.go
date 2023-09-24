package rbtree

import (
	"fmt"

	"github.com/emicklei/dot"
)

func (t *Tree) display() {
	g := dot.NewGraph(dot.Directed)
	if t.root == nil {
		return
	}
	t.root.displayHelper(g)

	fmt.Println(g.String())
}

func (n *Node) displayHelper(graph *dot.Graph) {
	// fmt.Println("key", n.key)
	// fmt.Println("color", n.color)
	// fmt.Println("red", red)
	// fmt.Println("black", black)
	var colorString string
	if n.color == red {
		// fmt.Println("key, red:", n.key)
		colorString = "red"
	} else {
		// fmt.Println("key, black:", n.key)
		colorString = "black"
	}
	me := graph.Node(n.key)
	me = me.Attr("color", colorString)

	// add left child to graph
	if n.leftChild != nil {
		left := graph.Node(n.leftChild.key)
		graph.Edge(me, left, "left")
		// pass control to left child
		n.leftChild.displayHelper(graph)
	}
	// add right child to graph
	if n.rightChild != nil {
		right := graph.Node(n.rightChild.key)
		graph.Edge(me, right, "right")
		// pass control to left child
		n.rightChild.displayHelper(graph)
	}
}

func RBDisplayTest() {
	tree := NewTree()

	tree.Insert("test", "test")

	tree.Insert("zzzztest", "test")
	tree.Insert("aaaatest", "test")
	tree.Insert("yyyytest", "test")
	tree.Insert("wwwwtest", "test")

	tree.display()
}
