package types

type Option struct {
	Name        string
	Alias       string
	Description string
}

type Command struct {
	Name    string
	Alias   string
	Handler func()
}
