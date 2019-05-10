package validation

type Validator interface {
	Validate() error
}

func Validate(v interface{}) error {
	if _, ok := v.(Validator); !ok {
		return nil
	}
	return v.(Validator).Validate()
}
