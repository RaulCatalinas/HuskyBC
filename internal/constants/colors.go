package constants

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

const colorPrefix = "\033["
const colorSuffix = "m"
const DEFAULT_COLOR = colorPrefix + string(enums.MessageColorDefault) + colorSuffix

var MESSAGE_COLORS = map[enums.MessageType]string{
	enums.MessageTypeSuccess: colorPrefix + string(enums.MessageColorGreen) + colorSuffix,
	enums.MessageTypeInfo:    colorPrefix + string(enums.MessageColorBlue) + colorSuffix,
	enums.MessageTypeError:   colorPrefix + string(enums.MessageColorRed) + colorSuffix,
	enums.MessageTypeConfig:  colorPrefix + string(enums.MessageColorWhite) + colorSuffix,
	enums.MessageTypeWarning: colorPrefix + string(enums.MessageColorYellow) + colorSuffix,
}
