package loggers

import (
	"io"

	"github.com/ThomasFerro/golog/formatters"
)

// Logger A logger
type Logger interface {
	Output() io.Writer
	Formatter() formatters.Formatter
}

// GoLogger A logger
type GoLogger struct {
	output    io.Writer
	formatter formatters.Formatter
}

// Output Get the logger's output
func (logger GoLogger) Output() io.Writer {
	return logger.output
}

// Formatter Get the logger's formatter
func (logger GoLogger) Formatter() formatters.Formatter {
	return logger.formatter
}

// NewLogger Create a new logger
func NewLogger(output io.Writer, formatter formatters.Formatter) Logger {
	return GoLogger{
		output,
		formatter,
	}
}
