package cache

import (
	"fmt"
)

type DoesNotExistError struct {
	Key string
}

func (e *DoesNotExistError) Error() string {
	return fmt.Sprintf("key %q does not exist", e.Key)
}

func IsDoesNotExistError(err error) bool {
	_, ok := err.(*DoesNotExistError)
	return ok
}
