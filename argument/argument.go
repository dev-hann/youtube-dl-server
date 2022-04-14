package command

import (
	"flag"
	"github.com/youtube-dl-server/command/argument"
	"github.com/youtube-dl-server/command/option"
	"github.com/youtube-dl-server/err"
)

type Command struct {
	option   *option.Option
	argument *argument.Argument
	args     []string
}

func InitCommand() *Command {
	return &Command{
		option:   option.InitOption(),
		argument: argument.InitArgument(),
	}
}

func (c *Command) Parse() error {
	flag.Parse()
	c.args = flag.Args()
	if len(c.args) == 0 {
		return err.NoArgument
	}
	return nil
}

func (c *Command) Run(
	startFunc func(configPath string),
	upgradeFunc func(),
) {

}
