package mobile

import (
	"code.olapie.com/types"
)

type PhoneNumber types.PhoneNumber

func (pn *PhoneNumber) String() string {
	return (*types.PhoneNumber)(pn).String()
}

func (pn *PhoneNumber) InternationalFormat() string {
	return (*types.PhoneNumber)(pn).InternationalFormat()
}

func NewPhoneNumber(callingCode int32, number int64) *PhoneNumber {
	return &PhoneNumber{
		Code:   callingCode,
		Number: number,
	}
}

func ParsePhoneNumber(s string, code int) *PhoneNumber {
	pn, _ := types.NewPhoneNumberV2(s, code)
	return (*PhoneNumber)(pn)
}
