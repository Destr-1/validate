package validate

import "fmt"

var ErrBadRequest = fmt.Errorf("bad request")

func validate(title string, text string) error {
	if title == "" || len(title) > 100 || text == "" || len(text) > 500 {
		return ErrBadRequest
	}
	return nil
}
