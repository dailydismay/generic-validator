package validator

func Validate(rules ...Rule) error {
	for _, r := range rules {
		if ve := r.Check(); ve != nil {
			return ve
		}
	}

	return nil
}
