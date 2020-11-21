package aoc

import "errors"

type CalleeType = func([]string) error

type RegistryType struct {
	data map[int]CalleeType
}

func newRegistry() *RegistryType {
	var r RegistryType
	r.data = make(map[int]CalleeType)
	return &r
}

func (r *RegistryType) Register(day int, callee CalleeType) {
	r.data[day] = callee
}

func (r *RegistryType) Invoke(day int, args []string) error {
	if callee, found := r.data[day]; !found {
		return errors.New("invalid day")
	} else {
		return callee(args)
	}
}

var Registry = newRegistry()
