package formatters

import (
	"fmt"
	"strings"

	"github.com/ThomasFerro/golog/entries"
)

// JsonFormatter A JSON formatter
type JsonFormatter struct{}

func formatStringMessageAsJson(key string, value interface{}) string {
	return fmt.Sprintf("\"%v\": \"%v\"", key, value)
}

// Format Format the log into a key-value pair message
func (formatter JsonFormatter) Format(fields entries.Fields, level string, message string) string {
	formattedMetadata := make([]string, len(fields)+2)
	index := 2

	formattedMetadata[0] = formatStringMessageAsJson("level", level)
	formattedMetadata[1] = formatStringMessageAsJson("message", message)
	for key, value := range fields {
		if _, typeOk := value.(string); typeOk {
			formattedMetadata[index] = formatStringMessageAsJson(key, value)
		} else {
			formattedMetadata[index] = fmt.Sprintf("\"%v\": %v", key, value)
		}
		index++
	}

	joinedMetadata := strings.Join(formattedMetadata, ", ")
	return fmt.Sprintf("{ %v }", joinedMetadata)
}

// NewKvpFormatter Create a new key-value pair formatter
func NewJsonFormatter() Formatter {
	return JsonFormatter{}
}
