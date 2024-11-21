package types

type Option struct {
	Name        string
	Alias       string
	Description string
	Handler     func()
}
