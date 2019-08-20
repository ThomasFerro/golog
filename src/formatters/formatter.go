package formatters

import (
	"github.com/ThomasFerro/golog/entries"
)

// Formatter A log formatter
type Formatter interface {
	Format(entry entries.Entry, level string, message string) string
}
