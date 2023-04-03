package custerrors

import (
	"encoding/json"
	"strconv"
)

type custErrors struct {
	errorMessage string
	errorCode    int
}

func (e custErrors) Error() string {
	return e.errorMessage + " " + strconv.Itoa(e.errorCode)
}

func custError(err error) (string, error) {
	// return "", errors.New("something went wrong")
	switch e := err.(type) {
	case *json.SyntaxError:
		return "", custErrors{"json syntax error:" + e.Error() + "offset:" + strconv.FormatInt(e.Offset, 10), 123}
	default:
		return "", custErrors{"cust error output", 123}
	}
	// return "", custErrors{"cust error output", 123}
}
