package argument

import (
	"fmt"
	"io"
	"os"
)

const Logo = "YOUTUBEDLSERVER"

type Console struct {
	writer io.Writer
}

func InitConsole() *Console {
	return &Console{
		writer: os.Stdout,
	}
}

func (c *Console) Log(message ...any) {
	fmt.Fprintln(os.Stdout, message)
}

func (c *Console) ShowLogo() {
	c.Log(Logo)
}
