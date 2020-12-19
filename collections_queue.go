package aoc

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (q *Queue) Init() {
	q.l.Init()
}

func (q *Queue) Len() int {
	return q.l.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *Queue) Add(v interface{}) {
	q.l.PushBack(v)
}

func (q *Queue) Head() interface{} {
	e := q.l.Front()
	if e != nil {
		q.l.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) LookHead() interface{} {
	e := q.l.Front()
	if e != nil {
		return e.Value
	}
	return nil
}

func (q *Queue) Clone() *Queue {
	t := NewQueue()
	for !q.IsEmpty() {
		t.Add(q.Head())
	}
	c := NewQueue()
	for !t.IsEmpty() {
		v := t.Head()
		q.Add(v)
		c.Add(v)
	}
	return c
}
