package main

import (
	"fmt"
)

type PNode *Node
type Node struct {
	freq  int
	child map[rune]PNode
}

func newNode() PNode {
	return &Node{0, make(map[rune]PNode)}
}

func insertWord(words string, node PNode) bool {
	letters := words
	var tmpNode PNode
	tmpNode = node
	for _, letter := range letters {
		_node, error := tmpNode.child[letter]
		if error == false {
			tmpNode.child[letter] = newNode()
			tmpNode = tmpNode.child[letter]
		} else {
			tmpNode.child[letter] = _node
		}
	}
	tmpNode.freq++
	return true
}

func findWord(words string, node PNode) int {
	letters := words
	tmpNode := node
	for _, letter := range letters {
		_, error := tmpNode.child[letter]
		if error == false {
			return -1
		}
		tmpNode = tmpNode.child[letter]
	}
	return tmpNode.freq
}
func main() {
	var test PNode = newNode()
	fmt.Println(test.freq)
	insertWord("abc", test)
	fmt.Println(findWord("abc", test))
	findWord("ab", test)
}
