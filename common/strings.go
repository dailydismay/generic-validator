package common

import (
	"regexp"

	validator "github.com/dailydismay/generic-validator"
)

func Required(s string) bool {
	return s != ""
}

func MaxLen(maxLen int) validator.ConstraintCheckFn[string] {
	return func(s string) bool {
		return len(s) <= maxLen
	}
}

func Regexp(re *regexp.Regexp) validator.ConstraintCheckFn[string] {
	return func(s string) bool {
		return re.MatchString(s)
	}
}
