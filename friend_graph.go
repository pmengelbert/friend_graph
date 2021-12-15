package main

import "fmt"

type (
	Queue struct {
		size       int
		ringBuffer []int
		front      int // points just below the front of the queue
		rear       int
	}

	ListNode struct {
		Val  int
		Next *ListNode
	}
)

func NewQueue(n int) *Queue {
	return &Queue{
		size: n,
		// one unit of storage is left unused; explanation found in Knuth 2.2.2, exercise 1
		ringBuffer: make([]int, n+1),
		front:      1,
		rear:       1,
	}
}

func (q *Queue) insert(val int) error {
	if q.rear == q.size {
		q.rear = 1
	} else {
		q.rear++
	}

	if q.rear == q.front {
		return fmt.Errorf("overflow detected")
	}

	q.ringBuffer[q.rear] = val

	return nil
}

func (q *Queue) remove() (int, error) {
	if q.front == q.rear {
		return -1, fmt.Errorf("underflow detected")
	}

	if q.front == q.size {
		q.front = 1
	} else {
		q.front++
	}

	return q.ringBuffer[q.front], nil
}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}

func main() {
	friends := [][]int{{1, 2}, {3, 4}, {3, 6}, {6, 7}, {1, 9}}
	n := 7

	f1 := 4
	f2 := 7

	result, _ := isFriend(n, friends, f1, f2)
	fmt.Printf("Are %d and %d friends? -> %t\n", f1, f2, result)

	f1 = 3
	f2 = 6
	result, _ = isFriend(n, friends, f1, f2)
	fmt.Printf("Are %d and %d friends? -> %t\n", f1, f2, result)

	f1 = 1
	f2 = 6
	result, _ = isFriend(n, friends, f1, f2)
	fmt.Printf("Are %d and %d friends? -> %t\n", f1, f2, result)
}

func isFriend(n int, friends [][]int, f1, f2 int) (bool, error) {
	visited := make(map[int]bool)

	graph := makeGraph(n, friends)

	// optionally print the graph by uncommenting the following:
	// printGraph(graph)

	queue := NewQueue(n)

	// execute the search using the queue for breadth-first

	var err error
	nodeIndex := f1
	queue.insert(f1)

	for { // will exit when queue is empty
		if err != nil {
			return false, err
		}

		list := graph[nodeIndex]
		for current := list; current != nil; current = current.Next {
			val := current.Val

			if val == f2 {
				return true, nil
			}

			if !visited[val] {
				queue.insert(val)
				visited[val] = true
			}
		}

		if queue.isEmpty() {
			break
		}

		nodeIndex, err = queue.remove()
	}

	return false, nil
}

func printGraph(g map[int]*ListNode) {
	for k, v := range g {
		fmt.Printf("%d", k)

		for current := v; current != nil; current = current.Next {
			fmt.Printf(" -> %d", current.Val)
		}
		fmt.Printf("\n")
	}
}

// assumes there are no duplicate pairings in the input
// i.e. friends = [[1, 2], [1, 2]] shall not appear in the input
func makeGraph(size int, friends [][]int) map[int]*ListNode {
	nodes := make(map[int]*ListNode)

	graphBuilder := make(map[int][]int)
	for _, pair := range friends {
		f1 := pair[0]
		f2 := pair[1]

		graphBuilder[f1] = append(graphBuilder[f1], f2)
		graphBuilder[f2] = append(graphBuilder[f2], f1)
	}

	// making two passes is not ideal, but I wanted to quickly provide a working version of the code.
	for k, v := range graphBuilder {
		if len(v) == 0 {
			nodes[k] = nil
			continue
		}

		head := &ListNode{
			Val:  v[0],
			Next: nil,
		}
		last := head

		for _, val := range v[1:] {
			newNode := &ListNode{
				Val:  val,
				Next: nil,
			}

			last.Next = newNode
			last = newNode
		}

		nodes[k] = head
	}

	return nodes
}
