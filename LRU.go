package main

import "fmt"

type LRU struct {
	head *lrunode
	size int
	data map[string]*lrunode
}

type lrunode struct {
	value string
	prev *lrunode
	next *lrunode
}

func main() {
	lru := &LRU{data: map[string]*lrunode{}, size: 5}
	lru.push("하나")
	lru.push("국민")
	lru.push("우리")
	lru.push("우리")
	lru.push("우리")
	lru.push("우리")
	lru.push("하나")
	lru.push("우리")
	fmt.Println(lru)
	lru.printLRU()

}

func (l *LRU) printLRU() {
	for node := l.head; node != nil; node = node.next {
		fmt.Print(node.value)
	}
}

func (l *LRU) push(bank string) {
	node := &lrunode{value: bank}

	value, exist := l.data[bank]
	if !exist {
		// 여기가 값이 없은 곳. 추가해야해서 len 체크를 해줘야 함.
		if len(l.data) == l.size {
			// size가 꽉 찼으면 데이터 삭제해야 함 마지막에 있는거 삭제.
			// map에서 먼저 삭제
			// 삭제 구현하기.

			//todo: 여기만 구현하면됨.
		} else {
			// size가 아직 max가 아니면 데이터 삭제 안해도 됨 헤드에만 추가하면 되는 케이스.
			// tail은 그대로 tail
		}
		l.data[bank] = node

	} else {
		// 값이 있으면 len 체크는 안해도 되는데 중간 노드 사이 연결을 변경해주어야 한다.
		// 근데 값이 제일 앞에 있을 때랑 뒤에 있을 때 생각해야함. 먼저, 값이 제일 앞에 있을 경우에는 prev가 비어있을 것임. valve.prev == nil 임
		// 값이 제일 뒤에 있을 경우에는 next가 비어있을 것임. value. next == nil 임.
		// 값이 있으면 이 값 전에꺼를 그 다음 노드로 이어줘야함. map에서 삭제할필요도 없음.
		if value.prev == nil { // 제일 앞에 있는 경우에는 할 거 없음.
			return
		} else if value.next == nil { // 제일 뒤에 있는 경우에는 연결 끊어줘야 함.
			fmt.Println("여기")
			value.prev.next = nil
		} else { // 그 외 여기서 문제가 발생한 것 같음.

			value.prev.next = value.next
			value.next.prev = value.prev
		}
	}
	// 기존 헤드 prev에 node 추가 후 head를 value로 변경

	if l.head != nil {
		l.head.prev = node
	}
	node.next = l.head
	l.head = node
}


