package formatters

import (
	"fmt"
	"strings"

	"github.com/ThomasFerro/golog/entries"
)

// KvpFormatter A key-value pair formatter
type KvpFormatter struct{}

func formatStringMessage(key string, value interface{}) string {
	return fmt.Sprintf("%v=\"%v\"", key, value)
}

// Format Format the log into a key-value pair message
func (formatter KvpFormatter) Format(fields entries.Fields, level string, message string) string {
	formattedMetadata := make([]string, len(fields)+2)
	index := 2

	formattedMetadata[0] = fmt.Sprintf("level=%v", level)
	formattedMetadata[1] = formatStringMessage("message", message)
	for key, value := range fields {
		if _, typeOk := value.(string); typeOk {
			formattedMetadata[index] = formatStringMessage(key, value)
		} else {
			formattedMetadata[index] = fmt.Sprintf("%v=%v", key, value)
		}
		index++
	}

	return strings.Join(formattedMetadata, " ")
}

// NewKvpFormatter Create a new key-value pair formatter
func NewKvpFormatter() Formatter {
	return KvpFormatter{}
}
