package validator

type ConstraintCheckFn[T any] func(T) bool

type Constraint[T any] interface {
	Check(T) bool
	Message() string
}

type constraint[T any] struct {
	msg   string
	check ConstraintCheckFn[T]
}

func NewConstraint[T any](message string, check ConstraintCheckFn[T]) Constraint[T] {
	return &constraint[T]{message, check}
}

func (c *constraint[T]) Check(v T) bool {
	return c.check(v)
}

func (c *constraint[T]) Message() string {
	return c.msg
}
