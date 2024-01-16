package errors

import "errors"

var (
	ErrConnDatabase = errors.New("FAILED_CONNECT_DATABASE")

	ErrInvalidRequest = errors.New("INVALID_REQUEST")

	ErrBadRequest = errors.New("BAD_REQUEST")

	ErrRequestTimeout = errors.New("REQUEST_TIMEOUT")

	ErrParsingHTML = errors.New("FAILED_PARSING_HTML")

	ErrSendMail = errors.New("FAILED_TO_SEND_MAIL")

	ErrPasswordNotSame = errors.New("PASSWORD_IS_NOT_SAME")

	ErrFailedCreateAccount = errors.New("FAILED_TO_CREATE_NEW_USER")

	ErrUserNotVerified = errors.New("USER_IS_NOT_VERIFIED")

	ErrFailedCountEmailUser = errors.New("FAILED_COUNT_EMAIL_USER")

	ErrEmailAlreadyExist = errors.New("EMAIL_ALREADY_EXIST")

	ErrFailedSaveAccount = errors.New("FAILED_SAVE_ACCOUNT")

	ErrFailedGetAccount = errors.New("FAILED_GET_ACCOUNT")

	ErrFailedVerifyAccount = errors.New("FAILED_VERIFY_ACCOUNT")

	ErrNotYetActivityDone = errors.New("NOT_YET_ACTIVITY_DONE")

	ErrRegisterRequestNotValid = errors.New("REGISTER_REQUEST_NOT_VALID")

	ErrAllActivityAlreadyDone = errors.New("ALL_ACTIVITY_ALREADY_DONE")

	ErrCantUnchelcklistActivity = errors.New("CANT_UNCHECKLIST_ACTIVITY")
)
