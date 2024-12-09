package utils

import (
	"bytes"
	"fmt"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) String() string {
	var out bytes.Buffer
	for p := range s {
		out.WriteString(fmt.Sprintf("%+v, ", p))
	}
	return out.String()
}

func (s Set[T]) Add(x T) {
	if !s.Contains(x) {
		s[x] = struct{}{}
	}
}

func (s Set[T]) Contains(x T) bool {
	_, exists := s[x]
	return exists
}

func (s Set[T]) Size() int {
	return len(s)
}
