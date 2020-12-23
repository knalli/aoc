package aoc

type LinkElement struct {
	sentinel bool
	root     *LinkedList
	prev     *LinkElement
	next     *LinkElement
	value    interface{}
}

func (e *LinkElement) Prev() *LinkElement {
	if prev := e.prev; prev != e.root.sentinel {
		return prev
	}
	return nil
}

func (e *LinkElement) Next() *LinkElement {
	if next := e.next; next != e.root.sentinel {
		return next
	}
	return nil
}

func (e *LinkElement) Value() interface{} {
	return e.value
}

type LinkedList struct {
	front    *LinkElement
	back     *LinkElement
	size     int
	ring     bool
	sentinel *LinkElement
}

func NewLinkedList(ring bool) *LinkedList {
	l := &LinkedList{
		ring:     ring,
		sentinel: &LinkElement{sentinel: true},
	}
	if ring {
		l.front = l.sentinel
		l.back = l.sentinel
	}
	return l
}

func (l *LinkedList) Len() int {
	return l.size
}
func (l *LinkedList) Front() *LinkElement {
	if front := l.front; front != l.sentinel {
		return front
	}
	return nil
}
func (l *LinkedList) Back() *LinkElement {
	if back := l.back; back != l.sentinel {
		return back
	}
	return nil
}

func (l *LinkedList) Add(value interface{}) *LinkElement {
	if after := l.back; after == l.sentinel {
		l.initFirst(value)
		return l.front
	} else {
		return l.AddAfter(after, value)
	}
}

func (l *LinkedList) AddAfter(after *LinkElement, value interface{}) *LinkElement {
	if after == nil || after == l.sentinel {
		panic("after argument should not be nil or sentinel")
	} else {
		e := &LinkElement{
			root:  l,
			value: value,
			prev:  after,
			next:  after.next,
		}
		after.next = e
		if l.ring {
			if e.next == l.front {
				l.back = e
			}
			e.next.prev = e
		} else {
			if e.next == l.sentinel {
				l.back = e
			} else {
				e.next.prev = e
			}
		}
		l.size++
		return e
	}
}

func (l *LinkedList) Remove(e *LinkElement) {
	if l.front == e && l.back == e {
		l.front = l.sentinel
		l.back = l.sentinel
		l.size = 0
	} else if l.front == e {
		l.front = e.next
		e.next.prev = e.prev
		if l.ring {
			e.prev.next = e.next
		}
		l.size--
	} else if l.back == e {
		l.back = e.prev
		e.prev.next = e.next
		if l.ring {
			e.next.prev = e.prev
		}
		l.size--
	} else {
		e.prev.next = e.next
		e.next.prev = e.prev
		l.size--
	}
}

func (l *LinkedList) Each(iter func(e *LinkElement)) {
	l.EachStartAt(l.front, iter)
}

func (l *LinkedList) EachStartAt(start *LinkElement, iter func(e *LinkElement)) {
	t := start
	for {
		iter(t)
		t = t.next
		if t == start {
			break
		}
	}
}

func (l *LinkedList) Filter(filter func(e *LinkElement) bool) *LinkElement {
	start := l.front
	t := start
	for {
		if filter(t) {
			return t
		}
		t = t.next
		if t == start {
			break
		}
	}
	return nil
}

func (l *LinkedList) initFirst(v interface{}) {
	e := &LinkElement{
		root:  l,
		value: v,
		prev:  l.sentinel,
		next:  l.sentinel,
	}
	l.front = e
	l.back = e
	l.size = 1
	if l.ring {
		e.next = e
		e.prev = e
	}
}

func (l *LinkedList) ToString(transform func(e *LinkElement) string) string {
	result := ""
	l.Each(func(e *LinkElement) {
		result += transform(e)
	})
	return result
}
