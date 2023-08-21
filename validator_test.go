package validator_test

import (
	"testing"

	validator "github.com/dailydismay/generic-validator"
	"github.com/dailydismay/generic-validator/common"
	"github.com/stretchr/testify/assert"
)

func Test_Validator(t *testing.T) {
	gteThanZeroMessage := "must be gte than zero"
	invalidFieldName := "testint"
	invalidFieldValue := -1
	gteThanZeroCheckFn := common.Gte(0)

	type testCase struct {
		name          string
		rules         []validator.Rule
		assertionErr  assert.ErrorAssertionFunc
		isMatchingErr func(error) bool
	}

	testCases := []testCase{
		{
			name: "Reject Negative Int",
			rules: []validator.Rule{
				validator.NewRule[int](invalidFieldName, invalidFieldValue, []validator.Constraint[int]{
					validator.NewConstraint[int](gteThanZeroMessage, gteThanZeroCheckFn),
				}),
			},
			assertionErr: assert.Error,
			isMatchingErr: func(err error) bool {
				if !validator.IsValidationError(err) {
					return false
				}
				ve := validator.MustAsValidationError(err)
				return ve.Field() == invalidFieldName && ve.Message() == gteThanZeroMessage
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			err := validator.Validate(tc.rules...)
			tc.assertionErr(t, err)

			if err != nil && tc.isMatchingErr != nil {
				assert.True(t, tc.isMatchingErr(err))
			}
		})
	}
}
