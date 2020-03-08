package main

import "fmt"

type Queue struct {
	items [5]interface{}
	depth int
	head int
}

type Stack struct {
	items [5]interface{}
	depth int
	head int
}

func main(){
	a := new(Queue)
	a.New(5)

	b := new(Stack)
	b.New(2)

	val, err := b.Pop()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(val)

	err = b.Push(5)
	err = b.Push("string test")
	fmt.Println("b.times: ", b.items)
	fmt.Println("b.head: ", b.head)

	err = b.Push(4)
	if err != nil{
		fmt.Println(err)
	}

	val, err = b.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("b.items: ", b.items)
	fmt.Println("b.head: ", b.head)
	fmt.Println("val: ", val)
}

func (q *Queue) New(length int) {
	q.depth = length
	q.head = -1
	return
}

func (q *Queue) Push(value int) (err error){
	if q.head == q.depth-1 {
		err = fmt.Errorf("Queue is full.")
		return
	}

	q.head = q.head +1
	q.items[q.head] = value
	return
}

func (q *Queue) Pop() (value interface{}, err error){
	if q.head < 0 {
		err = fmt.Errorf("Queue is empty.")
		return
	}

	value = q.items[0]
	for i:=0; i<len(q.items)-1; i++ {
		q.items[i] = q.items[i+1]
	}
	q.head = q.head -1
	return
}

func (s *Stack) New(length int) {
	s.depth = length
	s.head = -1
	return
}

func (s *Stack) Push(value interface{}) (err error){
	if s.head == s.depth-1 {
		err = fmt.Errorf("Stack is full")
		return
	}

	s.head = s.head +1
	s.items[s.head] = value
	return
}

func (s *Stack) Pop() (value interface{}, err error){
	if s.head < 0 {
		err = fmt.Errorf("Stack is empty")
		return
	}

	value = s.items[s.head]
	s.head = s.head -1
	return
}
