package cli

type Validator interface {
	Validate(*Command) []error
}

func IsTypeValid(t string) bool {
	return t == "string" || t == "int" || t == "bool"
}
