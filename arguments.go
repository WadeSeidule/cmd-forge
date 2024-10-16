package cli

import "fmt"

type PositonalArg struct {
	Name       string
	Type       string
	Required   bool
	Help       string
	Validator Validator
}

func (a *PositonalArg) Validate(c *Command) []error {
	errs := []error{}
	if !IsTypeValid(a.Type) {
		errs = append(errs, fmt.Errorf("invalid type %s for argument %s", a.Type, a.Name))
	}

	errs = append(errs, a.Validator.Validate(c)...)
	return errs
}

func (a *PositonalArg) HelpString() string {
	required := ""
	if a.Required {
		required = ", required"
	}
	return fmt.Sprintf("Argument: %s (%s%s) %s", a.Name, a.Type, required, a.Help)
}

type FlagArg struct {
	Name       string
	ShortName  string
	Type       string
	Required   bool
	Help       string
	Validator Validator
}

func (f *FlagArg) Validate(c *Command) []error {
	errs := []error{}
	if !IsTypeValid(f.Type) {
		errs = append(errs, fmt.Errorf("invalid type %s for flag %s", f.Type, f.Name))
	}
	errs = append(errs, f.Validator.Validate(c)...)
	return errs
}

func (f *FlagArg) HelpString() string {
	required := ""
	if f.Required {
		required = ", required"
	}
	shortName := ""
	if f.ShortName != "" {
		shortName = ", -" + f.ShortName
	}

	return fmt.Sprintf("{--%s%s} (%s%s) %s", f.Name, shortName, f.Type, required, f.Help)
}
