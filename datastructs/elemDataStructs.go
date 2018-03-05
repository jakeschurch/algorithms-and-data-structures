package datastructs

import "errors"

// Stack struct is implementation of LIFO container object.
type Stack struct {
	top   int
	stack []Node
	limit interface{}
}

func newStack(limit interface{}) *Stack {
	return &Stack{top: 0, limit: limit}
}

// Pop func is implementation of delete operation.
func (stack *Stack) Pop() (node Node, err error) {
	node = stack.stack[stack.top]
	if stack.top != 0 {
		stack.top--
		stack.stack = stack.stack[:stack.top]
		return node, nil
	} else {
		return node, errors.New("buffer underflow")
	}
}

// Push func is implementation of insert operation.
func (stack *Stack) Push(node Node) error {
	if stack.top != stack.limit {
		stack.top++
		stack.stack[stack.top] = node
		return nil
	}
	return errors.New("buffer overflow")
}

// NewQueue instantiates a new Queue.
func NewQueue(limit int) *Queue {
	return &Queue{head: 0, tail: 1, limit: limit, queue: make([]Node, 0, limit)}
}

// Queue is an implementation of a FIFO container type.
type Queue struct {
	head  int
	tail  int
	queue []Node
	limit int
}

func (queue *Queue) GetQueue() []Node {
	if queue.tail < queue.head {
		nodes := append(queue.queue[:queue.tail], queue.queue[queue.head:]...)
		return nodes
	}
	return queue.queue
}

// Enqueue stores value in the queue.
func (queue *Queue) Enqueue(node Node) error {
	switch queue.tail {
	case queue.limit + 1:
		queue.tail = 0
		// return errors.New("buffer overflow")
	}
	queue.queue[queue.tail] = node
	queue.tail++
	return nil
}

// Dequeue removes and returns value from the queue.
func (queue *Queue) Dequeue() (Node, error) {
	node := queue.queue[queue.head]
	switch queue.head {
	case queue.limit + 1:
		return node, errors.New("buffer overflow")
	case queue.tail + 1:
		return node, errors.New("buffer underflow")
	default:
		queue.head++
	}
	return node, nil
}

// Node represents data stored in a container.
type Node struct {
	data interface{}
}

// NewNode instantiates a new Node.
func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

// Ex 10.1-6: Show how to implement a queue using two stacks.
// Analyze the running time of the queue operations.
//
// Stacks: S_1[1...j...n], S_2[1...i...n]
//
// Enqueue()	O(n)
// 1. 		S_2[1...i] <- S_1[1...j]			n
// 2. 		S_2[i + 1] <- valIn					1
// 3. 		S_1[1...j + 1] <- S_2[1...i + 1]	n
//
// Dequeue()	O(n)
// 1. 		S_2[1...i] <- S_1[1...j - 1]		n
// 2. 		valOut <- S_1[j]					1
// 3. 		S_1[1...j - 1] <- S_2[1...i - 1]	n
//
// Ex 10.1-7: Show how to implement a stack[LIFO] using two queues[FIFO].
// Analyze the running time of the stack operations.
//
// Queues: Q_1[1...j...n], Q_2[1...i...n]
//
// Push()	O(1)
// 1.		S_1[j] <- S_2[1]	// S_2[1] now empty	1
// 2.		S_2[1] <- valIn							1
//
// Pop()	O(1)
// 1. 			valOut <- S_1[1]					1
// 2. 			S_1[1...j - 1] <- S_1[2...j]		1

func newLinkedList() *linkedList {
	return &linkedList{prev: nil, key: 0, next: nil}
}

func insert(L, new *linkedList) {
	new.next = L.prev
	if L.prev != nil {
		L.prev.prev = new
	}
	L.prev = new
	new.prev = nil
}

type listNode struct {
	key  interface{}
	prev *listNode
	next *listNode
}

// doublyLinkedLists have both prev and next
type linkedList struct {
	key  uint
	prev *linkedList
	next *linkedList
}
