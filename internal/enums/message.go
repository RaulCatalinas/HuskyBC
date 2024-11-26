package enums

type MessageType string

const (
	MessageTypeSuccess MessageType = "success"
	MessageTypeError   MessageType = "error"
	MessageTypeInfo    MessageType = "info"
	MessageTypeConfig  MessageType = "config"
	MessageTypeWarning MessageType = "warning"
)
