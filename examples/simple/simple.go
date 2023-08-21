package main

import (
	"fmt"

	validator "github.com/dailydismay/generic-validator"
	"github.com/dailydismay/generic-validator/common"
)

type Payload struct {
	S string
	N int
}

func main() {
	p := Payload{}
	// Validate accepts Rule list
	err := validator.Validate(
		validator.NewRule("S", p.S, []validator.Constraint[string]{
			validator.NewConstraint(
				"must not be empty",
				common.Required,
			)}),
		validator.NewRule("N", p.N, []validator.Constraint[int]{
			validator.NewConstraint(
				"must be greater than zero",
				common.Gte(1),
			)}),
	)

	ve := validator.MustAsValidationError(err)
	// S must not be empty
	fmt.Printf("%s %s\n", ve.Field(), ve.Message())
}
