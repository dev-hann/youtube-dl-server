package argument

import (
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

}

func (c *Console) ShowLogo() {

}
