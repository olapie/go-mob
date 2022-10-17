package mobile

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type Handler interface {
	SaveSecret(name, value string) bool
	DeleteSecret(name string) bool
	GetSecret(name string) string
	NeedSignIn()
}

type AuthErrorChecker struct {
	h Handler
}

func NewAuthErrorChecker(h Handler) *AuthErrorChecker {
	return &AuthErrorChecker{
		h: h,
	}
}

func (c *AuthErrorChecker) Check(err error) {
	if err == nil {
		return
	}

	code := ToError(err).Code
	if code == http.StatusUnauthorized || code == int(codes.Unauthenticated) {
		go c.h.NeedSignIn()
	}
}
