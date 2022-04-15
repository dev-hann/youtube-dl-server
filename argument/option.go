package argument

import "flag"

type Option struct {
	config string
}

/// Use cobra
func InitOption() *Option {
	c := flag.String("config", "./config.yaml", "Server Config file path.")
	return &Option{
		*c,
	}
}
