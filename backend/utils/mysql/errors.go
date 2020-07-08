package mysql

import (
	"fmt"
	"strings"
)

// custom error
var (
	ErrNoneAffected = fmt.Errorf("none affected")
)

// IsDuplicateError return true if err is Error 1062: Duplicate entry
func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062: Duplicate entry")
}
