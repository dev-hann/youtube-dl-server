package argument

import (
	"flag"
	"github.com/youtube-dl-server/err"
)

type Argument struct {
	console *Console
	Option  *Option
	args    []string
}

func InitArgument() *Argument {
	return &Argument{
		Option:  InitOption(),
		console: InitConsole(),
	}
}

/// Use Cobra
func (c *Argument) Parse() error {
	flag.Parse()
	c.args = flag.Args()
	if len(c.args) == 0 {
		return err.NoArgument
	}
	return nil
}

func (c *Argument) Run(
	startFunc func(configPath string, console *Console),
	upgradeFunc func(console *Console),
	versionFunc func(console *Console),
) {
	console := c.console
	arg := c.args[0]
	switch arg {
	case "start":
		startFunc(c.Option.config, console)
		break
	case "upgrade":
		upgradeFunc(console)
		break
	case "version":
		versionFunc(console)
		break
	}
}
