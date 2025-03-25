package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/types"
)

type defaultHelpFormatter struct {
	writer io.Writer
}

func newDefaultHelpFormatter(writer io.Writer) *defaultHelpFormatter {
	if writer == nil {
		writer = os.Stdout
	}
	return &defaultHelpFormatter{writer: writer}
}

func (defaultHelpFormatter *defaultHelpFormatter) Format(options []types.Option) string {
	var builder strings.Builder

	builder.WriteString("Usage: huskybc [options]\n\n")
	builder.WriteString("Command line for easy Husky configuration\n\n")
	builder.WriteString("Options:\n")

	const (
		nameWidth      = 15
		aliasWidth     = 5
		formatTemplate = "%-*s %-*s %s\n"
	)

	for _, option := range options {
		fmt.Fprintf(
			&builder,
			"  "+formatTemplate,
			nameWidth, option.Name,
			aliasWidth, option.Alias,
			option.Description,
		)
	}

	builder.WriteString("\n")

	return builder.String()
}

func ShowHelp(options []types.Option) {
	formatter := newDefaultHelpFormatter(os.Stdout)
	helpText := formatter.Format(options)

	fmt.Fprint(formatter.writer, helpText)
}

func ShowHelpWithWriter(options []types.Option, writer io.Writer) {
	formatter := newDefaultHelpFormatter(writer)
	helpText := formatter.Format(options)
	fmt.Fprint(writer, helpText)
}
