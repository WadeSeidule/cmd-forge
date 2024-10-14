package cli

import "fmt"

type FlagArg struct {
	Name       string
	ShortName  string
	Type       string
	Required   bool
	Help       string
	Validators []ValidatorFunc
}

func (f *FlagArg) HelpString() string {
	required := ""
	if f.Required {
		required = "required"
	}
	shortName := ""
	if f.ShortName != "" {
		shortName = ", -" + f.ShortName
	}

	return fmt.Sprintf("{--%s%s} (%s%s) %s", f.Name, shortName, f.Type, required, f.Help)
}
