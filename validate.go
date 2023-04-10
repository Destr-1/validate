package validate

import (
	"errors"
)

var ErrBadRequest = errors.New("bad request")

func Validate(title string, text string) error {
	if title == "" || len(title) > 100 || text == "" || len(text) > 500 {
		return ErrBadRequest
	}
	return nil
}
