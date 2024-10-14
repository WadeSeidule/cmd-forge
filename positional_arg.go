package cli

import "fmt"

type PositonalArg struct {
	Name       string
	Type       string
	Required   bool
	Help       string
	Validators []ValidatorFunc
}

func (a *PositonalArg) HelpString() string {
	required := ""
	if a.Required {
		required = ", required"
	}
	return fmt.Sprintf("Argument: %s (%s%s) %s", a.Name, a.Type, required, a.Help)
}
