package cache

import "fmt"

const size = 5

type (
	Node struct {
		Data string
		Left *Node
		Right *Node
	}
	Queue struct {
		Head *Node
		Tail *Node
		Length int
	}
	Hash map[string]*Node
	Cache struct {
		queue Queue
		hash Hash
	}
)

func New () Cache {
	return Cache {
		queue: NewQueue(),
		hash: NewHash(),
	}
}

func NewQueue () Queue {
	head := &Node {}
	tail := &Node {}

	head.Right = tail
	tail.Left = head

	return Queue { Head: head, Tail: tail }
}

func NewHash () Hash {
	return Hash {}
}

func (c *Cache) Save (data string) {
	var node *Node

	if val, ok := c.hash[data]; ok {
		node = c.remove(val)
	} else {
		node = &Node { Data: data }
	}

	c.add(node)
	c.hash[data] = node
}

func (c *Cache) remove (n *Node) *Node {
	left, right := n.Left, n.Right

	right.Left = left
	left.Right = right

	c.queue.Length -= 1
	delete(c.hash, n.Data)

	return n
}

func (c *Cache) add (n *Node) *Node {
	tmp := c.queue.Head.Right

	c.queue.Head.Right = n
	n.Left = c.queue.Head
	n.Right = tmp
	tmp.Left = n

	c.queue.Length += 1
	c.hash[n.Data] = n
	if c.queue.Length > size {
		c.remove(c.queue.Tail.Left)
	}

	return n
}

func (c *Cache) Display () {
	c.queue.Display()
}

func (q *Queue) Display () {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Data)
		if i < q.Length - 1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}