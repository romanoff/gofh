package gofh

func Init() *Flags {
	return &Flags{Commands: make([]*Command, 0)}
}

type Flags struct {
	Commands       []*Command
	DefaultHandler func()
}

type Callback func(map[string]string)

func (self *Flags) HandleCommand(pattern string, callback Callback) {
	self.Commands = append(self.Commands, &Command{Pattern: pattern, Callback: callback})
}

func (self *Flags) SetDefaultHandler(handler func()) {
	self.DefaultHandler = handler
}

func (self *Flags) Parse(args []string) {
	for _, command := range self.Commands {
		if command.Callback != nil {
			if options := command.Matches(args); options != nil && command.Callback != nil {
				command.Callback(options)
				return
			}
		}
	}
	if self.DefaultHandler != nil {
		self.DefaultHandler()
	}
}
