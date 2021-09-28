package errors

import (
	"errors"
	"fmt"
)

var (
	InternalServerError = fmt.Errorf("internal server error")
	NotFound            = fmt.Errorf("not found")
	AccessDenied        = fmt.Errorf("access denied")
)

func Is(err, internal error) bool {
	return errors.Is(err, internal)
}
