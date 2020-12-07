package aoc

import (
	"testing"
)

func TestNewQueue(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		q := NewQueue()
		if q.Len() != 0 {
			t.Errorf("empty queue should have length 0, got = '%v', want = '%v'", q.Len(), 0)
		}
		if !q.IsEmpty() {
			t.Errorf("empty queue should have be empty, got = '%v', want = '%v'", q.IsEmpty(), true)
		}
	})

	t.Run("Add one", func(t *testing.T) {
		q := NewQueue()
		q.Add("something")
		if q.Len() != 1 {
			t.Errorf("queue should have length 1, got = '%v', want = '%v'", q.Len(), 1)
		}
		if q.IsEmpty() {
			t.Errorf("queue should have be empty, got = '%v', want = '%v'", q.IsEmpty(), false)
		}
	})

	t.Run("Init will reset", func(t *testing.T) {
		q := NewQueue()
		q.Add("something")
		q.Init()
		if q.Len() != 0 {
			t.Errorf("empty queue should have length 0, got = '%v', want = '%v'", q.Len(), 0)
		}
	})

	t.Run("Add two", func(t *testing.T) {
		q := NewQueue()
		q.Add("something")
		q.Add("else")
		if q.Len() != 2 {
			t.Errorf("queue should have length 1, got = '%v', want = '%v'", q.Len(), 2)
		}
		if q.IsEmpty() {
			t.Errorf("queue should have be empty, got = '%v', want = '%v'", q.IsEmpty(), false)
		}
	})

	t.Run("Add and remove", func(t *testing.T) {
		q := NewQueue()
		q.Add("something")
		q.Add("else")
		if q.Len() != 2 {
			t.Errorf("queue should have length 1, got = '%v', want = '%v'", q.Len(), 2)
		}
		if head := q.Head(); head == nil {
			t.Errorf("queue head should be 1st, got = '%v', want = '%v'", head, "something")
		} else if head != "something" {
			t.Errorf("queue head should be 1st, got = '%v', want = '%v'", head, "something")
		}
		if q.Len() != 1 {
			t.Errorf("queue should have length 1, got = '%v', want = '%v'", q.Len(), 1)
		}
		if head := q.Head(); head == nil {
			t.Errorf("queue head should be 2nd, got = '%v', want = '%v'", head, "else")
		} else if head != "else" {
			t.Errorf("queue head should be 2nd, got = '%v', want = '%v'", head, "else")
		}
		if q.Len() != 0 {
			t.Errorf("empty queue should have length 0, got = '%v', want = '%v'", q.Len(), 0)
		}
	})

}
