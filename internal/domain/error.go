package domain

const (
	ERR_BAD_REQUEST     = "err_bad_request"
	ERR_NOT_FOUND       = "err_not_found"
	ERR_INTERNAL_SERVER = "err_internal_server"
	ERR_UNAUTHORIZED    = "err_unauthorized"
	ERR_FORBIDDEN       = "err_forbidden"
	ERR_CONFLICT        = "err_conflict"
)

type CustomeError struct {
	Code    string
	Message string
}

func (e *CustomeError) Error() string {
	return e.Message
}

func (e *CustomeError) ErrCode() string {
	return e.Code
}

func NewCustomeError(code string, message string) *CustomeError {
	return &CustomeError{Code: code, Message: message}
}
