package aoc

import "testing"

func TestNewStack(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		s := NewStack()
		if s.Len() != 0 {
			t.Errorf("empty stack should have length 0, got = '%v', want = '%v'", s.Len(), 0)
		}
		if !s.IsEmpty() {
			t.Errorf("empty stack should have be empty, got = '%v', want = '%v'", s.IsEmpty(), true)
		}
	})

	t.Run("Add one", func(t *testing.T) {
		s := NewStack()
		s.Add("something")
		if s.Len() != 1 {
			t.Errorf("stack should have length 1, got = '%v', want = '%v'", s.Len(), 1)
		}
		if s.IsEmpty() {
			t.Errorf("stack should have be empty, got = '%v', want = '%v'", s.IsEmpty(), false)
		}
	})

	t.Run("Init will reset", func(t *testing.T) {
		s := NewStack()
		s.Add("something")
		s.Init()
		if s.Len() != 0 {
			t.Errorf("empty stack should have length 0, got = '%v', want = '%v'", s.Len(), 0)
		}
	})

	t.Run("Add two", func(t *testing.T) {
		s := NewStack()
		s.Add("something")
		s.Add("else")
		if s.Len() != 2 {
			t.Errorf("stack should have length 1, got = '%v', want = '%v'", s.Len(), 2)
		}
		if s.IsEmpty() {
			t.Errorf("stack should have be empty, got = '%v', want = '%v'", s.IsEmpty(), false)
		}
	})

	t.Run("Add and remove", func(t *testing.T) {
		s := NewStack()
		s.Add("something")
		s.Add("else")
		if s.Len() != 2 {
			t.Errorf("stack should have length 1, got = '%v', want = '%v'", s.Len(), 2)
		}
		if head := s.Head(); head == nil {
			t.Errorf("stack head should be 1st, got = '%v', want = '%v'", head, "else")
		} else if head != "else" {
			t.Errorf("stack head should be 1st, got = '%v', want = '%v'", head, "else")
		}
		if s.Len() != 1 {
			t.Errorf("stack should have length 1, got = '%v', want = '%v'", s.Len(), 1)
		}
		if head := s.Head(); head == nil {
			t.Errorf("stack head should be 2nd, got = '%v', want = '%v'", head, "something")
		} else if head != "something" {
			t.Errorf("stack head should be 2nd, got = '%v', want = '%v'", head, "something")
		}
		if s.Len() != 0 {
			t.Errorf("empty stack should have length 0, got = '%v', want = '%v'", s.Len(), 0)
		}
	})

}
