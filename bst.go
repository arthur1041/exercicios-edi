package main

import "fmt"

type ITree interface {
	Add(value int)
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	PrintLevels()
	Height() int
	Remove(value int) *BstNode
}

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

type Bst struct {
	root *BstNode
}

func (t *Bst) Add(value int) {
	t.root = addNode(t.root, value)
}

func addNode(node *BstNode, value int) *BstNode {
	if node == nil {
		return &BstNode{value: value}
	}
	if value < node.value {
		node.left = addNode(node.left, value)
	} else if value > node.value {
		node.right = addNode(node.right, value)
	}
	return node
}

func (t *Bst) Search(value int) bool {
	return searchNode(t.root, value)
}

func searchNode(node *BstNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	}
	if value < node.value {
		return searchNode(node.left, value)
	}
	return searchNode(node.right, value)
}

func (t *Bst) Min() int {
	node := t.root
	if node == nil {
		return 0
	}
	for node.left != nil {
		node = node.left
	}
	return node.value
}

func (t *Bst) Max() int {
	node := t.root
	if node == nil {
		return 0
	}
	for node.right != nil {
		node = node.right
	}
	return node.value
}

func (t *Bst) PrintPre() {
	printPre(t.root)
	fmt.Println()
}

func printPre(node *BstNode) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		printPre(node.left)
		printPre(node.right)
	}
}

func (t *Bst) PrintIn() {
	printIn(t.root)
	fmt.Println()
}

func printIn(node *BstNode) {
	if node != nil {
		printIn(node.left)
		fmt.Printf("%d ", node.value)
		printIn(node.right)
	}
}

func (t *Bst) PrintPos() {
	printPos(t.root)
	fmt.Println()
}

func printPos(node *BstNode) {
	if node != nil {
		printPos(node.left)
		printPos(node.right)
		fmt.Printf("%d ", node.value)
	}
}

func (t *Bst) PrintLevels() {
	if t.root == nil {
		return
	}
	queue := []*BstNode{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	fmt.Println()
}

func (t *Bst) Height() int {
	return height(t.root)
}

func height(node *BstNode) int {
	if node == nil {
		return 0
	}
	leftH := height(node.left)
	rightH := height(node.right)
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func (t *Bst) Remove(value int) *BstNode {
	t.root = removeNode(t.root, value)
	return t.root
}

func removeNode(node *BstNode, value int) *BstNode {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = removeNode(node.left, value)
	} else if value > node.value {
		node.right = removeNode(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		minRight := node.right
		for minRight.left != nil {
			minRight = minRight.left
		}
		node.value = minRight.value
		node.right = removeNode(node.right, minRight.value)
	}
	return node
}

func main() {
	tree := &Bst{}
	values := []int{47, 30, 54, 28, 63, 5, 75, 86, 21, 7, 42}
	for _, v := range values {
		tree.Add(v)
	}

	fmt.Println("InOrder:")
	tree.PrintIn()

	fmt.Println("PreOrder:")
	tree.PrintPre()

	fmt.Println("PosOrder:")
	tree.PrintPos()

	fmt.Println("LevelOrder:")
	tree.PrintLevels()

	fmt.Println("Altura:", tree.Height())
	fmt.Println("Min:", tree.Min())
	fmt.Println("Max:", tree.Max())

	fmt.Println("Busca 42:", tree.Search(42))
	fmt.Println("Removendo 30...")
	tree.Remove(30)
	tree.PrintIn()
}
