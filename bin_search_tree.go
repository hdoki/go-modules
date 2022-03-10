package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//var Count int

func (n *Node) Insert(k int) {
	if n.Key == k {
		fmt.Println("Node Exists")
		return
	}
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) {

}

func main() {
	tree := &Node{Key: 100}
	fmt.Printf("Root: %d", tree.Key)
	tree.Insert(20)
	tree.Insert(120)
	tree.Insert(100)
	fmt.Println(*tree)

}
