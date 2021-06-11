package lfserve

import "log"

type Mappings struct {
	Pairs map[int64]int64
	Free  queue
}

type queue struct {
	head  *node
	tail  *node
	count int64
}

type node struct {
	val  int64
	next *node
}

var m *Mappings = nil

func NewMap() {
	log.Printf("new mapping")
	m = &Mappings{
		Pairs: make(map[int64]int64),
		Free: queue{
			head:  &node{val: 0, next: nil},
			tail:  nil,
			count: 0,
		},
	}
	m.Free.tail = m.Free.head
}

func GetMap() *Mappings {
	return m
}

func (q *queue) add(val int64) {
	q.tail.next = &node{val: val, next: nil}
	q.tail = q.tail.next
	q.count = q.count + 1
	log.Printf("added %d, count %d", val, q.count)
}

func (q *queue) remove() int64 {
	if q.count == 0 {
		return -1
	}
	tmp := q.head.next
	q.head.next = tmp.next
	q.count = q.count - 1
	log.Printf("removed %d, count %d", tmp.val, q.count)
	if q.count == 0 {
		q.tail = q.head
	}
	return tmp.val
}

func (q *queue) size() int64 {
	return q.count
}

func (q *queue) peek() int64 {
	if q.count == 0 {
		return -1
	}
	return q.head.next.val
}
