package argument

const (
	start = iota
	upgrade
)

type Argument struct {
	upgrader *Upgrade
}

func InitArgument() *Argument {
	return &Argument{
		upgrader: &Upgrade{},
	}
}

func (a *Argument) upgrade() {
	a.upgrader.Do()
}
