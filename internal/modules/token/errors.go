package token

import "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")

func IsErrInvalidAccessToken(err error) bool {
	return err == ErrInvalidAccessToken
}
