package common

import validator "github.com/dailydismay/generic-validator"

func Gte(boundary int) validator.ConstraintCheckFn[int] {
	return func(i int) bool {
		return i >= boundary
	}
}
