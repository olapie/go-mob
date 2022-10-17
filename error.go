package mobile

import (
	"fmt"
	"reflect"

	"code.olapie.com/errors"
	"google.golang.org/grpc/status"
)

type Error errors.Error

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) String() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewClientError(message string) *Error {
	return NewError(-1, message)
}

func ToError(err error) *Error {
	if err == nil {
		return nil
	}

	if v := reflect.ValueOf(err); !v.IsValid() || v.IsZero() {
		return nil
	}

	cause := errors.Cause(err)
	if e, ok := cause.(*Error); ok && e != nil {
		return NewError((*errors.Error)(e).Code, err.Error())
	}

	if e, ok := cause.(*errors.Error); ok && e != nil {
		return NewError(e.Code, e.Message)
	}

	if st, ok := status.FromError(err); ok {
		if e := errors.FromString(st.Message()); e != nil {
			return NewError(e.Code, e.Message)
		}
		return NewError(int(st.Code()), st.Message())
	}

	return NewError(0, err.Error())
}
