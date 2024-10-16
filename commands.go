package cli

import (
	"fmt"
	"strings"
)

type Commands struct {
	Commands map[string]*Command
	Program  string
}

func (c *Commands) Contains(cmd string) (bool, *Command) {
	cmdStruct, exists := c.Commands[cmd]
	return exists, cmdStruct
}

func (c *Commands) ListCommands() string {
	var cmdNames []string
	for name := range c.Commands {
		cmdNames = append(cmdNames, name)
	}
	return strings.Join(cmdNames, ", ")
}

func (c *Commands) Run(args []string) error {
	cmd, args, err := c.parseArgs(args)
	if err != nil {
		return err
	}
	exists, cmdStruct := c.Contains(cmd)
	if !exists {
		return fmt.Errorf("command '%s' not found. Expected one of '%s'", cmd, c.ListCommands())
	}
	err = cmdStruct.Parse(args)
	if err != nil {
		return err
	}
	return cmdStruct.Run()
}

func (c *Commands) AddCommand(cmd *Command) {
	if c.Commands == nil {
		c.Commands = make(map[string]*Command)
	}
	cmd.Program = c.Program
	c.Commands[cmd.Name] = cmd
}

func (c *Commands) parseArgs(args []string) (string, []string, error) {
	if len(args) == 0 {
		return "", nil, fmt.Errorf("expected a command. Expected one of '%s'", c.ListCommands())
	}
	cmd, cmdArgs := args[0], args[1:]
	if cmd == "help" {
		fmt.Print(c.HelpString())
		return "", nil, nil
	}
	return cmd, cmdArgs, nil
}

func (c *Commands) HelpString() string {
	helpStrings := []string{}
	for _, cmd := range c.Commands {
		cmdHelpString := cmd.HelpString()
		cmdHelpString = fmt.Sprintf("Command: %s\n", cmd.Name) + cmdHelpString
		helpStrings = append(helpStrings, cmdHelpString)
	}
	return strings.Join(helpStrings, "\n\n")
}

func NewCommandSet(program string) *Commands {
	return &Commands{Program: program}
}
