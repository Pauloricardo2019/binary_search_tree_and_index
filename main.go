package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type node struct {
	index Person
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

func (n node) String() string {
	return strconv.Itoa(n.index.ID)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inOrderTraversalByNode(sb, root.right)
}

func (b bst) Add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}

func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{
			index: Person{
				ID: value,
			},
		}
	}

	if value < root.index.ID {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	return root
}

func (b bst) Search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}
	if value == root.index.ID {
		return root, true
	} else if value < root.index.ID {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
	return nil, false
}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}
	if value > root.index.ID {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.index.ID {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}

			root.index.ID = temp.index.ID

			root.left = b.removeByNode(root.left, value)
		}
	}

	return root
}

func main() {
	people, err := GenerateObject(100000)
	if err != nil {
		return
	}

	len := len(people)
	n := &node{
		people[0],
		nil,
		nil,
	}
	n.left = &node{
		people[1],
		nil,
		nil,
	}
	n.right = &node{
		people[2],
		nil,
		nil,
	}

	b := bst{
		root: n,
		len:  3,
	}

	newLen := len - 3

	fmt.Println(b)

	for i := 0; i < newLen; i++ {
		b.Add(people[i].ID)
	}

	fmt.Println("Objetos inseridos na  arvore!!!")

	start := time.Now()
	go func() {
		fmt.Println(SearchByIndex(people, people[50000].ID))
	}()
	fmt.Println("Search by index", time.Since(start))

	start = time.Now()
	go func() {
		fmt.Println(b.Search(people[50000].ID))
	}()
	fmt.Println("Search by binary tree", time.Since(start))

	time.Sleep(time.Second * 5)

}
