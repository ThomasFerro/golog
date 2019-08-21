package formatters

import (
	"github.com/ThomasFerro/golog/entries"
)

// Formatter A log formatter
type Formatter interface {
	Format(fields entries.Fields, level string, message string) string
}
