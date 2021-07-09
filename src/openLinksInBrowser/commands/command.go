package commands

type Commands int

const (
	RUN Commands = iota
	ADD_LINK
	SHOW_ALL_LINKS
)

func (commands Commands) String() string {
	return [...]string{"RUN", "SHOW_ALL_LINKS"}[commands]
}
