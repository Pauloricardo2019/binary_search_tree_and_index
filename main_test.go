package main

import (
	"fmt"
	"testing"
)

func TestGenerateObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		people, err := GenerateObject(1000)
		if err != nil {
			return
		}

		len := len(people)
		n := &node{
			people[0].ID,
			nil,
			nil,
		}
		n.left = &node{
			people[1].ID,
			nil,
			nil,
		}
		n.right = &node{
			people[2].ID,
			nil,
			nil,
		}

		b := bst{
			root: n,
			len:  len,
		}

		newLen := len - 3

		fmt.Println(b)

		for i := 0; i < newLen; i++ {
			b.Add(people[i].ID)
		}

		for i := 0; i < newLen; i++ {
			fmt.Println(b.Search(people[i].ID))
		}
	}
}
