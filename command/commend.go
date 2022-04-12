package command

import (
	"github.com/youtube-dl-server/command/argument"
	"github.com/youtube-dl-server/command/option"
)

type Command struct {
	option   *option.Option
	argument *argument.Argument
}

func InitCommand() *Command {
	return &Command{
		option:   option.InitOption(),
		argument: argument.InitArgument(),
	}
}
