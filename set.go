package mobile

import (
	"code.olapie.com/conv"
	"code.olapie.com/types"
)

type IntSet struct {
	types.Set[int]
}

func (s *IntSet) String() string {
	return conv.MustJSONString(s.Slice())
}

type Int64Set struct {
	types.Set[int64]
}

func (s *Int64Set) String() string {
	return conv.MustJSONString(s.Slice())
}

type Float64Set struct {
	types.Set[float64]
}

func (s *Float64Set) String() string {
	return conv.MustJSONString(s.Slice())
}

type StringSet struct {
	types.Set[string]
}

func NewStringSet() *StringSet {
	return new(StringSet)
}
