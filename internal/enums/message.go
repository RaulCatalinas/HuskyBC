package enums

type MessageType string
type MessageColor string

const (
	MessageTypeSuccess MessageType = "success"
	MessageTypeError   MessageType = "error"
	MessageTypeInfo    MessageType = "info"
	MessageTypeConfig  MessageType = "config"
	MessageTypeWarning MessageType = "warning"
)

const (
	MessageColorGreen   MessageColor = "32"
	MessageColorBlue    MessageColor = "36"
	MessageColorRed     MessageColor = "31"
	MessageColorWhite   MessageColor = "37"
	MessageColorYellow  MessageColor = "33"
	MessageColorDefault MessageColor = "0"
)
