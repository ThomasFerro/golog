package formatters_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ThomasFerro/golog/entries"
	"github.com/ThomasFerro/golog/formatters"
)

func TestFormatTheMessage(t *testing.T) {
	formattedMessage := formatters.NewJsonFormatter().Format(entries.Fields{}, "", "message")

	if !strings.Contains(formattedMessage, "message: \"message\"") {
		t.Errorf("The entry's message is not formatted")
	}
}

func TestFormatInAJsonObject(t *testing.T) {
	formattedMessage := formatters.NewJsonFormatter().Format(entries.Fields{}, "", "message")

	if !strings.HasPrefix(formattedMessage, "{") || !strings.HasSuffix(formattedMessage, "}") {
		t.Errorf("The entry is not formatted in a JSON object")
	}
}

func TestFormatInAJsonObjectWithAttributesSeparatedWithComma(t *testing.T) {
	formattedMessage := formatters.NewJsonFormatter().Format(entries.Fields{}, "info", "message")

	if strings.Count(formattedMessage, ",") != 1 {
		t.Errorf("No comma seperates the attributes")
	}
}

func TestFormatTheMetadata(t *testing.T) {
	fields := entries.Fields{
		"test":      42.42,
		"otherTest": true,
	}

	fmt.Println(formattedMessage)

	if !strings.Contains(formattedMessage, "test: 42.42") || !strings.Contains(formattedMessage, "otherTest: true") {
		t.Errorf("The entry's metadata are not formatted")
	}
}
