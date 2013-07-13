package gofh

import "strings"

type Command struct {
	Pattern  string
	Callback func(map[string]string)
	Options  []*Option
}

type ParsedOptions map[string]string
type CommandArguments []string

func (self *Command) Matches(args []string) map[string]string {
	patternArgs := strings.Split(self.Pattern, " ")
	if len(args) == 0 && len(patternArgs) > 0 {
		return nil
	}
	ca := CommandArguments(args)
	commandArguments := &ca
	options := &ParsedOptions{}
	for _, commandOption := range self.Options {
		extractCommandOption(commandOption, commandArguments, options)
	}
	for i, arg := range *commandArguments {
		if len(patternArgs) <= i {
			return nil
		}
		if patternArgs[i][0] == ':' {
			(*options)[patternArgs[i][1:]] = arg
		} else if patternArgs[i] != args[i] {
			return nil
		}
	}
	return *options
}

func extractCommandOption(option *Option, args *CommandArguments, parsedOptions *ParsedOptions) {
	for i, arg := range *args {
		if arg == "--"+option.Name {
			if option.Boolean {
				copy((*args)[i:], (*args)[i+1:])
				(*args) = (*args)[:len(*args)-1]
				(*parsedOptions)[option.Name] = "true"
			}
		}
	}
}
