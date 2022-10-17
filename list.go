package mobile

import (
	"code.olapie.com/mobile/nomobile"
	"code.olapie.com/types"
)

type StringList struct {
	nomobile.List[string]
}

func NewStringList() *StringList {
	l := &StringList{}
	return l
}

type IntList struct {
	nomobile.List[int]
}

func NewIntList() *IntList {
	return new(IntList)
}

type Int64List struct {
	nomobile.List[int64]
}

func NewInt64List() *Int64List {
	return new(Int64List)
}

type Float64List struct {
	nomobile.List[float64]
}

func NewFloat64List() *Float64List {
	return new(Float64List)
}

type BoolList struct {
	nomobile.List[bool]
}

func NewBoolList() *BoolList {
	return new(BoolList)
}

type ImageList struct {
	nomobile.List[*Image]
}

func NewImageList() *ImageList {
	return new(ImageList)
}

type PhoneNumberList struct {
	nomobile.List[*PhoneNumber]
}

func NewPhoneNumberList() *PhoneNumberList {
	return new(PhoneNumberList)
}

func (l *PhoneNumberList) Contains(phoneNumber *PhoneNumber) bool {
	for _, pn := range l.Elements {
		if pn.String() == (*types.PhoneNumber)(phoneNumber).String() {
			return true
		}
	}
	return false
}

func (l *PhoneNumberList) ContainsString(phoneNumber string) bool {
	for _, pn := range l.Elements {
		if pn.String() == phoneNumber {
			return true
		}
	}
	return false
}

type AnyList struct {
	nomobile.List[*Any]
}

func NewAnyList() *AnyList {
	return new(AnyList)
}

func (a *AnyList) FirstImage() *Image {
	for _, e := range a.Elements {
		te := (*types.Any)(e)
		if img, ok := te.Value().(*types.Image); ok {
			return (*Image)(img)
		}
	}
	return nil
}

type CountryList struct {
	nomobile.List[*Country]
}
