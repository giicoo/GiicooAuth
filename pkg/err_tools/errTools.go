package errTools

import "net/http"

type APIError interface {
	APIError() (int, string)
}

var (
	// Status: 500
	ErrInternalError = &sentinelAPIError{status: http.StatusInternalServerError, msg: "internal error"}

	// Status: 400
	ErrEmailUsed = &sentinelAPIError{status: http.StatusBadRequest, msg: "email already used"}

	// Status: 400
	ErrInvalidJSON = &sentinelAPIError{status: http.StatusBadRequest, msg: "invalid json"}

	// Status: 400
	ErrWrongPassword = &sentinelAPIError{status: http.StatusBadRequest, msg: "invalid password"}

	// Status: 400
	ErrInvalidTokenSyntax = &sentinelAPIError{status: http.StatusBadRequest, msg: "invalid token: invalid syntax"}

	// Status: 400
	ErrInvalidTokenExpired = &sentinelAPIError{status: http.StatusBadRequest, msg: "invalid token: expired"}
)

type sentinelAPIError struct {
	status int
	msg    string
}

func (e sentinelAPIError) Error() string {
	return e.msg
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.msg
}

type sentinelWrappedError struct {
	error
	sentinel *sentinelAPIError
}

func (e sentinelWrappedError) Is(err error) bool {
	return e.sentinel == err
}

func (e sentinelWrappedError) APIError() (int, string) {
	return e.sentinel.APIError()
}

func WrapError(err error, sentinel *sentinelAPIError) error {
	return sentinelWrappedError{error: err, sentinel: sentinel}
}
