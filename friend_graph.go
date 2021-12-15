package main

import "fmt"

type (
  Queue struct {
    size int,
    ringBuffer []int,
    front int // points just below the front of the queue
    rear int
  }

  ListNode struct {
    Val int,
    Next *ListNode,
  }
)

func NewQueue(n int) *Queue {
  return &Queue{
    size: n,
    ringBuffer: make([]int, size + 1),
    front: 1,
    rear: 1,
  }
}

func (q *Queue) insert(val int) {
  if q.rear == q.size {
    q.rear = 1
  } else {
    q.rear++
  }

  q.ringBuffer[q.rear] = val
}

func (q *Queue) remove() int {
  if q.front == q.size {
    q.front = 1
  } else {
    q.front++
  }

  return q.ringBuffer[q.front]
}

func (q *Queue) isEmpty() bool {
  return q.front == q.rear
}

func isFriend(n int, friends [][]int, f1, f2 int) bool {
  visited := make(map[int]bool)

  // build the graph without duplicating nodes
  var graph []*ListNode

  queue := NewQueue(n)
  if len(graph) == 0 {
    return false
  }

  // execute the search using the queue for breadth-first
  queue.insert(f1)
  for nodeIndex := queue.remove(); !queue.isEmpty(); nodeIndex = queue.remove() { // determine termination condition -- probably when queue is empty
    list := graph[nodeIndex]
    for current := list; current != nil; current = current.Next {
      val := current.Val

      if val == f2 {
        return true
      }

      if !visited[val] {
        queue.insert(val)
        visited[val] = true
      }
    }
  }

  return false
}

func main() {
  // you can write to stdout for debugging purposes, e.g.
  fmt.Println("This is a debug message")
}

// []*ListNode

// n: # number of students
// V, E


// [[1,2],[3,4],[3,6],[6,7],[1,9]] n = 10 check if 4 and 7 are friends


result := make(map[int]*ListNode)
// for pair := range friends {
  f1 := pair[0]
  f2 := pair[1]


  result[f1] = append(result[f1], f2)
  result[f2] = append(result[f2], f1)
}

// map[int]bool := {
  4: true,
  3: true,
  6: true,
}

4 , 3, 6, 7

// There are n students in a class with id from 0 to n - 1. Some of them are friends, while some of them are not. We use a two dimentional array to represent their relationship.
//
// [[1,2],[3,4],[3,6],[6,7],[1,9]] n = 10 check if 4 and 7 are friends
// [[3, 6], [7,6]]
// 3,4, 3, 6, 6, 7  -> 4, 7 return true
// 1 and 3 return false
//
// bool IsFriend(int n, int[][] friends, int f1, int f2
