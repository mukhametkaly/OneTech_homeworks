package main

import "fmt"
type Node struct {
	next *Node
	prev *Node
	value int
}

type Deque struct {
	start *Node
	end *Node
	size int
}

func (dq *Deque) peek(isEnd bool)  {
	if dq.size == 0 {
		return
	}else if isEnd {
		fmt.Println(dq.end.value)
	}else {
		fmt.Println(dq.start.value)
	}
}




func (deque *Deque) push_back(i int)  {
	if deque.start == nil {
		node := Node{
			value: i,
		}
		deque.start = &node
		deque.end = &node
		deque.size = 1
	}else {
		node := Node{
			prev: deque.end,
			value: i,
		}
		deque.end.next = &node
		deque.end = &node
		deque.size++
	}

}

func (deque *Deque) push_front(i int)  {
	if deque.start == nil {
		node := Node{ value: i,}
		deque.start = &node
		deque.end = &node
		deque.size = 1
	}else {
		node := Node{
			value: i,
			next: deque.start,
		}

		deque.start.prev = &node
		deque.start = &node
		deque.size++
	}

}

func (dq *Deque)  pop_front(){
	if dq.start == nil {
		return
	} else if dq.size == 1 {
		dq.start = nil
		dq.end = nil
		dq.size--
	} else {
		dq.start.next.prev = nil
		dq.start = dq.start.next
		dq.size--
	}
}

func (dq *Deque)  pop_back(){
	if dq.end == nil {
		return
	}else if dq.size == 1 {
		dq.start = nil
		dq.end = nil
		dq.size--
	} else {
		dq.end.prev.next = nil
		dq.end = dq.end.prev
		dq.size--
	}
}


func main()  {
	dq := Deque{}
	dq.push_front(1)
	dq.peek(false)
	dq.pop_back()
	dq.peek(false)

	dq.push_front(2)
	dq.push_front(1)
	dq.peek(false)
	dq.peek(true)
	dq.push_back(5)
	dq.peek(true)
	dq.pop_back()
	dq.peek(true)

	dq.push_front(10)
	dq.push_front(20)
	dq.peek(false)
	dq.pop_front()
	dq.peek(false)
	dq.pop_front()
	dq.peek(false)
	fmt.Println(dq.size)
	dq.pop_back()
	dq.pop_back()
	fmt.Println(dq.size)
	dq.pop_back()

}