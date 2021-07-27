package matchers

import (
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"strings"
)

type Matcher struct {
	command    *args.Command
	manOptions *man.OptionList
}

func NewMatcher(command *args.Command, options *man.OptionList) *Matcher {
	return &Matcher{
		command:    command,
		manOptions: options,
	}
}

func (m *Matcher) Match() *man.OptionList {
	var found []*man.Option

	for _, arg := range m.command.Args {
		// match double dashed args
		if strings.Contains(arg, "--") {
			arg = strings.Trim(arg, "-")

			for _, opt := range m.manOptions.Options() {
				// search for exact match
				if opt.Name == arg || opt.ShortName == arg {
					found = append(found, opt)
				}
			}
		}

		// match single dash args
		if strings.Contains(arg, "-") {
			arg = strings.Trim(arg, "-")

			for _, opt := range m.manOptions.Options() {
				// search for exact match
				if opt.Name == arg || opt.ShortName == arg {
					found = append(found, opt)
				} else {
					// split arg into characters and search
					for _, c := range arg {
						cc := string(c)
						if opt.Name == cc || opt.ShortName == cc {
							found = append(found, opt)
						}
					}
				}
			}
		}
	}

	return man.NewOptionList(found)
}
