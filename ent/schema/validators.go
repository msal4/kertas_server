package schema

import (
	"errors"
	"regexp"
)

var yearExp = regexp.MustCompile("\\d{4}-\\d{4}")

func validateYear(s string) error {
	if !yearExp.Match([]byte(s)) {
		return errors.New("format must be YYYY-YYYY")
	}

	return nil
}
