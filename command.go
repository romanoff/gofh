package gofh

import "strings"

type Command struct {
	Pattern  string
	Callback func(map[string]string)
}

func (self *Command) Matches(args []string) map[string]string {
	patternArgs := strings.Split(self.Pattern, " ")
	if len(args) == 0 && len(patternArgs) > 0 {
		return nil
	}
	options := make(map[string]string)
	for i, arg := range args {
		if patternArgs[i][0] == ':' {
			options[patternArgs[i][1:]] = arg
		} else if patternArgs[i] != args[i] {
			return nil
		}
	}
	return options
}
