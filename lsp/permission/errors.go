package permission

import "errors"

var (
	ErrPermissionExist    = errors.New("already exist")
	ErrPermissionNotExist = errors.New("not exist")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrDisabled           = errors.New("disabled")
)

func IsPermissionError(err error) bool {
	if err == ErrDisabled || err == ErrPermissionDenied || err == ErrPermissionExist || err == ErrPermissionNotExist {
		return true
	}
	return false
}
