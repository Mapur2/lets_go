package main

import "fmt"

const Size = 5

type Node struct {
	Left, Right *Node
	Val         string
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := Node{}
	tail := Node{}

	head.Right = &tail
	tail.Left = &head
	head.Left = nil
	tail.Right = nil
	return Queue{Head: &head, Tail: &tail}
}

func (c *Cache) CheckWork(key string) {
	node := &Node{}
	if val, ok := c.Hash[key]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: key}
	}
	c.Add(node)
	c.Hash[key] = node
}

func (c *Cache) Remove(val *Node) *Node {
	fmt.Printf("remove %s \n", val.Val)
	left := val.Left
	right := val.Right

	left.Right = right
	right.Left = left

	c.Queue.Size -= 1
	delete(c.Hash, val.Val)
	return val
}

func (c *Cache) Add(m *Node) {
	fmt.Printf("add %s \n", m.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = m
	m.Left = c.Queue.Head
	m.Right = tmp
	tmp.Left = m

	c.Queue.Size++
	if c.Queue.Size > Size {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (q *Queue) Display() {
	t := q.Head.Right
	if t == q.Tail {
		fmt.Println("Queue is empty")
	} else {
		for t != q.Tail {
			fmt.Print(t.Val, " ")
			t = t.Right
		}
		fmt.Println()
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func main() {
	fmt.Println("Starting cache")
	cache := NewCache()

	for _, word := range []string{"parrot", "avacado", "tiger", "parrot", "cat", "dog", "elephant", "horse"} {
		cache.CheckWork(word)
		cache.Display()
	}
}
