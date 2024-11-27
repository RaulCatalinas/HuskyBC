package utils

import (
	"fmt"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
)

type WriteMessageProps struct {
	Type    enums.MessageType
	Message string
}

func WriteMessage(props WriteMessageProps) {
	color, exists := constants.MESSAGE_COLORS[props.Type]

	if !exists {
		color = constants.DEFAULT_COLOR
	}

	fmt.Println(color + props.Message + constants.DEFAULT_COLOR)
}
