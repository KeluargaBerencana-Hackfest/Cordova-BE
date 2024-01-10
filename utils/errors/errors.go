package errors

import "errors"

var (
	ErrSigningJWT = errors.New("FAILED_SIGNING_JWT")

	ErrClaimsJWT = errors.New("FAILED_GET_CLAIMS_FROM_JWT")

	ErrConnDatabase = errors.New("FAILED_CONNECT_DATABASE")

	ErrInvalidRequest = errors.New("INVALID_REQUEST")

	ErrBadRequest = errors.New("BAD_REQUEST")

	ErrRequestTimeout = errors.New("REQUEST_TIMEOUT")

	ErrParsingHTML = errors.New("FAILED_PARSING_HTML")

	ErrSendMail = errors.New("FAILED_TO_SEND_MAIL")

	ErrPasswordNotSame = errors.New("PASSWORD_IS_NOT_SAME")

	ErrFailedCreateAccount = errors.New("FAILED_TO_CREATE_NEW_USER")

	ErrUserNotVerified = errors.New("USER_IS_NOT_VERIFIED")
)
