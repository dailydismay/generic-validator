package main

import (
	validator "github.com/dailydismay/generic-validator"
)

func MustModulo(divider, modulo int) validator.ConstraintCheckFn[int] {
	return func(i int) bool {
		return i%divider == modulo
	}
}

func main() {
	x := 11

	err := validator.Validate(
		validator.NewRule("x", x, []validator.Constraint[int]{
			validator.NewConstraint(
				"some message",
				// you can build custom contraint fns
				MustModulo(2, 1),
			)}),
	)

	if err != nil {
		panic(err)
	}
}
