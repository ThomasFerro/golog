package loggers

import (
	"io"
)

// Logger A logger
type Logger interface {
	Output() io.Writer
	// TODO : formatter
}

// GoLogger A logger
type GoLogger struct {
	output io.Writer
}

// Output Get the logger's output
func (logger GoLogger) Output() io.Writer {
	return logger.output
}

// NewLogger Create a new logger
func NewLogger(output io.Writer) Logger {
	return GoLogger{
		output,
	}
}
