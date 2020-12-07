package aoc

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{l: list.New()}
}

func (s *Stack) Init() {
	s.l.Init()
}

func (s *Stack) Len() int {
	return s.l.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}

func (s *Stack) Add(v interface{}) {
	s.l.PushFront(v)
}

func (s *Stack) Head() interface{} {
	e := s.l.Front()
	if e != nil {
		s.l.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) LookHead() interface{} {
	e := s.l.Front()
	if e != nil {
		return e.Value
	}
	return nil
}
