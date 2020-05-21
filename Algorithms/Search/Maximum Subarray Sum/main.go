package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Key    int64
	Left   *Node
	Right  *Node
	height int64
}

func main() {
	r := bufio.NewScanner(bufio.NewReader(os.Stdin))
	r.Split(bufio.ScanWords)
	t := ReadLong(r)

	for i := int64(0); i < t; i++ {
		n := ReadLong(r)
		m := ReadLong(r)

		a := make([]int64, n)
		for j := int64(0); j < n; j++ {
			a[j] = ReadLong(r)
		}

		fmt.Println(MaximumSum(a, m))
	}
}

func MaximumSum(arr []int64, m int64) int64 {
	var max, sm int64
	var r *Node

	for _, aVal := range arr {
		sm = (sm + aVal) % m
		r = insertOrIgnore(r, sm)
		if sm > max {
			max = sm
		}

		nd := r.next(sm)
		if nd == nil {
			continue
		}

		if d := (sm - nd.Key + m) % m; d > max {
			max = d
		}
	}

	return max
}

func ReadLong(s *bufio.Scanner) int64 {
	s.Scan()
	i, _ := strconv.ParseInt(s.Text(), 10, 64)
	return i
}

func balance(n *Node, key int64) *Node {
	b := n.Left.treeHeight() - n.Right.treeHeight()

	if b > 1 && key < n.Left.Key {
		return n.rotateRight()
	}

	if b < -1 && key > n.Right.Key {
		return n.rotateLeft()
	}

	if b > 1 && key > n.Left.Key {
		n.Left = n.Left.rotateLeft()
		return n.rotateRight()
	}

	if b < -1 && key < n.Right.Key {
		n.Right = n.Right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func createNode(key int64) *Node {
	n := &Node{}
	n.Key = key
	n.height = 1
	return n
}

func insertOrIgnore(n *Node, key int64) *Node {
	if n == nil {
		return createNode(key)
	}

	if key < n.Key {
		n.Left = insertOrIgnore(n.Left, key)
	} else if key > n.Key {
		n.Right = insertOrIgnore(n.Right, key)
	} else if key == n.Key {
		return n
	}

	n.height = max(n.Left.treeHeight(), n.Right.treeHeight()) + 1

	return balance(n, key)
}

func leftMost(n *Node) *Node {
	for {
		if n.Left == nil {
			return n
		}
		n = n.Left
	}
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// next returns node with the next value or nil if the key is maximum
func (n *Node) next(key int64) *Node {
	var prevRight *Node
	c := n
	for {
		if c.Key == key {
			if c.Right != nil {
				return leftMost(c.Right)
			}
			if prevRight != nil {
				return prevRight
			}
			break
		}
		if c.Key < key {
			c = c.Right
			continue
		}
		if c.Key > key {
			prevRight = c
			c = c.Left
			continue
		}
	}

	return nil
}

func (n *Node) rotateLeft() *Node {
	r := n.Right
	t2 := r.Left
	r.Left = n
	n.Right = t2
	n.height = max(n.Left.treeHeight(), n.Right.treeHeight()) + 1
	r.height = max(r.Left.treeHeight(), r.Right.treeHeight()) + 1
	return r
}

func (n *Node) rotateRight() *Node {
	r := n.Left
	t2 := r.Right
	r.Right = n
	n.Left = t2
	n.height = max(n.Left.treeHeight(), n.Right.treeHeight()) + 1
	r.height = max(r.Left.treeHeight(), r.Right.treeHeight()) + 1
	return r
}

func (n *Node) treeHeight() int64 {
	if n == nil {
		return 0
	}
	return n.height
}
