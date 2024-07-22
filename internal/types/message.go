package types

type MessageType string

const (
	Success MessageType = "success"
	Error   MessageType = "error"
	Info    MessageType = "info"
	Config  MessageType = "config"
)
