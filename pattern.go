package mobile

import (
	"code.olapie.com/conv"
	"code.olapie.com/ola/httpkit"
	"code.olapie.com/types"
)

func IsEmailAddress(s string) bool {
	return conv.IsEmailAddress(s)
}

func IsURL(s string) bool {
	return httpkit.IsURL(s)
}

func IsPhoneNumber(phoneNumber string) bool {
	return types.IsPhoneNumber(phoneNumber)
}

func IsDate(s string) bool {
	return conv.IsDate(s)
}
