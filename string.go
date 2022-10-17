package mobile

import (
	"code.olapie.com/conv"
)

func SmartLen(s string) int {
	n := 0
	for _, c := range s {
		if c <= 255 {
			n++
		} else {
			n += 2
		}
	}

	return n
}

func SquishString(s string) string {
	return conv.SquishString(s)
}

type StringMap struct {
	m map[string]string
}

func NewStringMap() *StringMap {
	return &StringMap{
		m: make(map[string]string),
	}
}

func (m *StringMap) Len() int {
	return len(m.m)
}

func (m *StringMap) Get(key string) string {
	return m.m[key]
}

func (m *StringMap) Set(key, val string) {
	m.m[key] = val
}

func (m *StringMap) Keys() *StringList {
	l := new(StringList)
	l.Elements = conv.GetMapKeys(m.m)
	return l
}
