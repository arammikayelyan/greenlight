package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define an error to use when unable to parse
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// We expect that hte incoming JSON value will be a string in the format
	// "<runtime> mins".
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// Convert int32 to a Runtime type and assign to a receiver.
	*r = Runtime(i)

	return nil
}

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
