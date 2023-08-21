package validator

type field[T any] struct {
	name  string
	value T
}

type rule[T any] struct {
	field[T]
	constraints []Constraint[T]
}

type Rule interface {
	Check() error
}

func NewRule[T any](name string, value T, constraints []Constraint[T]) Rule {
	return &rule[T]{field[T]{name, value}, constraints}
}

func (r *rule[T]) Check() error {
	for _, constraint := range r.constraints {
		if !constraint.Check(r.field.value) {
			return &validationError{r.field.name, constraint.Message()}
		}
	}

	return nil
}
