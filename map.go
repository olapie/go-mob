package mobile

import (
	"code.olapie.com/types"
)

type Map struct {
	m types.M
}

func NewMap() *Map {
	return &Map{m: types.M{}}
}

func (m *Map) GetInt64(key string) int64 {
	return m.m.Int64(key)
}

func (m *Map) GetFloat64(key string) float64 {
	return m.m.Float64(key)
}

func (m *Map) GetString(key string) string {
	return m.m.String(key)
}

func (m *Map) GetBool(key string) bool {
	return m.m.Bool(key)
}

func (m *Map) GetInt64List(key string) *Int64List {
	switch v := m.m[key].(type) {
	case []int64:
		l := NewInt64List()
		l.Elements = v
		return l
	case *Int64List:
		return v
	default:
		return nil
	}
}

func (m *Map) GetStringList(key string) *StringList {
	switch v := m.m[key].(type) {
	case []string:
		l := new(StringList)
		l.Elements = v
		return l
	case *StringList:
		return v
	default:
		return nil
	}
}

func (m *Map) GetPhoneNumber(key string) *types.PhoneNumber {
	return m.m.PhoneNumber(key)
}

func (m *Map) GetLocation(key string) *Point {
	v, _ := m.m[key].(*Point)
	return v
}

func (m *Map) SetInt64(key string, val int64) {
	m.m[key] = val
}

func (m *Map) SetFloat64(key string, val float64) {
	m.m[key] = val
}

func (m *Map) SetString(key string, val string) {
	m.m[key] = val
}

func (m *Map) SetBool(key string, val bool) {
	m.m[key] = val
}

func (m *Map) SetPhoneNumber(key string, val *types.PhoneNumber) {
	m.m[key] = val
}

func (m *Map) SetInt64s(key string, val *Int64List) {
	m.m[key] = val.Elements
}

func (m *Map) SetStrings(key string, val *StringList) {
	m.m[key] = val.Elements
}

func (m *Map) SetLocation(key string, val *Point) {
	m.m[key] = val
}

func (m *Map) M() types.M {
	return m.m
}
