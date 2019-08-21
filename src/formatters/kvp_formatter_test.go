package formatters_test

import (
	"strings"
	"testing"

	"github.com/ThomasFerro/golog/entries"
	"github.com/ThomasFerro/golog/formatters"
)

func TestShouldFormatTheMessage(t *testing.T) {
	formattedMessage := formatters.NewKvpFormatter().Format(entries.Fields{}, "", "message")

	if !strings.Contains(formattedMessage, "message=\"message\"") {
		t.Errorf("The entry's message is not formatted")
	}
}

func TestShouldFormatTheLevel(t *testing.T) {
	formattedMessage := formatters.NewKvpFormatter().Format(entries.Fields{}, "debug", "")

	if !strings.Contains(formattedMessage, "level=debug") {
		t.Errorf("The entry's level is not formatted")
	}
}

func TestShouldFormatTheMetadata(t *testing.T) {
	fields := entries.Fields{
		"test": 42.42,
		"otherTest": true,
	}

	formattedMessage := formatters.NewKvpFormatter().Format(fields, "", "")

	if !strings.Contains(formattedMessage, "test=42.42") || !strings.Contains(formattedMessage, "otherTest=true") {
		t.Errorf("The entry's metadata are not formatted")
	}
}

func TestShouldFormatAStringMetadataQuoted(t *testing.T) {
	fields := entries.Fields{
		"test": "a metadata",
	}

	formattedFields := formatters.NewKvpFormatter().Format(fields, "", "")

	if !strings.Contains(formattedFields, "test=\"a metadata\"") {
		t.Errorf("The entry is not formatted with the quoted string metadata")
	}
}

func TestShouldFormatMultilineStringMetadata(t *testing.T) {
	fields := entries.Fields{
		"test": `a test 
		metadata`,
	}

	formattedFields := formatters.NewKvpFormatter().Format(fields, "", "")

	if !strings.Contains(formattedFields, `test="a test 
		metadata"`) {
		t.Errorf("The entry is not formatted with the multiline string metadata")
	}
}
