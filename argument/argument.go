package argument

import (
	"flag"
	"github.com/youtube-dl-server/err"
)

type Argument struct {
	console *Console
	args    []string
}

func InitCommand() *Argument {
	return &Argument{
		console: InitConsole(),
	}
}

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
) {

}
