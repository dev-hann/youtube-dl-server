package argument

import (
	"fmt"
	"io"
	"os"
)

const Logo = "\nY88b   d88P                888             888               8888888b.  888 \n Y88b d88P                 888             888               888  \"Y88b 888 \n  Y88o88P                  888             888               888    888 888 \n   Y888P  .d88b.  888  888 888888 888  888 88888b.   .d88b.  888    888 888 \n    888  d88\"\"88b 888  888 888    888  888 888 \"88b d8P  Y8b 888    888 888 \n    888  888  888 888  888 888    888  888 888  888 88888888 888    888 888 \n    888  Y88..88P Y88b 888 Y88b.  Y88b 888 888 d88P Y8b.     888  .d88P 888 \n    888   \"Y88P\"   \"Y88888  \"Y888  \"Y88888 88888P\"   \"Y8888  8888888P\"  888 \n .d8888b.                                                                   \nd88P  Y88b                                                                  \nY88b.                                                                       \n \"Y888b.    .d88b.  888d888 888  888  .d88b.  888d888                       \n    \"Y88b. d8P  Y8b 888P\"   888  888 d8P  Y8b 888P\"                         \n      \"888 88888888 888     Y88  88P 88888888 888                           \nY88b  d88P Y8b.     888      Y8bd8P  Y8b.     888                           \n \"Y8888P\"   \"Y8888  888       Y88P    \"Y8888  888                           \n                                                                            \n                                                                            \n                                                                           "

type Console struct {
	writer io.Writer
}

func InitConsole() *Console {
	return &Console{
		writer: os.Stdout,
	}
}

// Log make separate 'log','info',debug'..
func (c *Console) Log(message ...any) {
	fmt.Fprintln(os.Stdout, message)
}

func (c *Console) ShowLogo() {
	c.Log(Logo)
}
