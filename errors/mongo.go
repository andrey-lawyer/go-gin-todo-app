package errors

import "strings"

func IsMongoError(err error) bool {
	return strings.Contains(err.Error(), "mongo")
}
