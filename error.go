package mob

import (
	"errors"
	"reflect"

	"go.olapie.com/types"
	"go.olapie.com/utils"
)

type Error struct {
	types.Error
}

var _ error = (*Error)(nil)

func NewError(code int, message string) *Error {
	return &Error{
		*types.NewError(code, message),
	}
}

func ToError(err error) *Error {
	if err == nil {
		return nil
	}

	if v := reflect.ValueOf(err); !v.IsValid() || v.IsZero() {
		return nil
	}

	var e *Error
	if errors.As(err, &e) {
		return e
	}

	code := utils.GetErrorCode(err)
	return NewError(code, err.Error())
}
