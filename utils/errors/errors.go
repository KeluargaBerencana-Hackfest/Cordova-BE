package errors

import "errors"

var (
	ErrSigningJWT = errors.New("FAILED_SIGNING_JWT")

	ErrClaimsJWT = errors.New("FAILED_GET_CLAIMS_FROM_JWT")
)