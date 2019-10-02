package formatters

import (
	"fmt"

	"github.com/ThomasFerro/golog/entries"
)

func formatStringMessage(key string, value interface{}, separator string) string {
	return fmt.Sprintf("%v%v\"%v\"", key, separator, value)
}

// Formatter A log formatter
type Formatter interface {
	Format(fields entries.Fields, level string, message string) string
}
