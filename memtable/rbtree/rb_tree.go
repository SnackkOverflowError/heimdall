package rbtree

import (
	"errors"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type color bool

const (
	black, red color = true, false
)

type Tree struct {
	root *Node
}

type Node struct {
	color      color
	leftChild  *Node
	rightChild *Node
	key        string
	value      string
}

func NewNode(key string, value string) Node {

	return Node{
		key:        key,
		value:      value,
		color:      red,
		leftChild:  nil,
		rightChild: nil,
	}

}

func NewTree() Tree {
	return Tree{root: nil}
}

func (t *Tree) Search(key string) *Node {

	n := t.root

	for n != nil {
		switch strings.Compare(key, n.key) {
		case 0:
			return n
		case 1:
			n = n.rightChild
		case -1:
			n = n.leftChild
		}
	}

	return nil
}

func (t *Tree) Insert(key, value string) error {

	if t.root == nil {
		node := NewNode(key, value)
		t.root = &node
		t.root.color = black
		return nil
	}

	// find path to insertion point
	n := t.root
	travelStack := stack.New()
	insertLeft := false

	for n != nil {
		travelStack.Push(n)

		switch strings.Compare(key, n.key) {
		case 0:
			// best scenario: node/key exists and we can change its value without any recolouring
			n.value = value
			return nil
		case 1:
			n = n.rightChild
			insertLeft = false
		case -1:
			n = n.leftChild
			insertLeft = true
		}
	}

	if travelStack.Len() == 0 {
		return errors.New("node stack is empty")
	}

	// create node to insert
	if parent, ok := travelStack.Peek().(*Node); ok {
		newNode := NewNode(key, value)
		if insertLeft {
			parent.leftChild = &newNode
		} else {
			parent.rightChild = &newNode
		}

	} else {
		// something has gone horribly wrong, return false
		return errors.New("node stack value cannot be casted to Node")
	}

	travelStack.Push(n)

	// recolor

	return nil
}
