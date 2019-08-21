package main_test

import (
	"fmt"
	"strings"
	"testing"

	golog "github.com/ThomasFerro/golog"
	"github.com/ThomasFerro/golog/entries"
	"github.com/ThomasFerro/golog/formatters"
	"github.com/ThomasFerro/golog/loggers"
)

type DumbWriter struct { 
	messages []string;
}

func (d *DumbWriter) Write(p []byte) (int, error) {
	d.messages = append(d.messages, string(p))
	return len(p), nil
}

func (d DumbWriter) LastLogContains(key string, value string) bool {
	lastLog := d.messages[len(d.messages) - 1]
	fmt.Println(lastLog)
	return strings.Contains(lastLog, fmt.Sprintf("%v=%v", key, value))
}

func NewDumbWriter() *DumbWriter {
	return &DumbWriter{
		messages: make([]string, 0),
	}
}

func NewDumbWriterAsLogOutput() *DumbWriter {
	dumbWriter := NewDumbWriter()
	formatter := formatters.NewKvpFormatter()
	golog.SetLoggers(
		loggers.NewLogger(dumbWriter, formatter),
	)
	return dumbWriter
}

func TestShouldDisplayTheDebugLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Debug("Test")
	
	if !dumbWriter.LastLogContains("level", "debug") {
		t.Error("The log is not written with debug level")
	}
}

func TestShouldDisplayTheInfoLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Info("Test")

	if !dumbWriter.LastLogContains("level", "info") {
		t.Error("The log is not written with info level")
	}
}

func TestShouldDisplayTheWarnLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Warn("Test")

	if !dumbWriter.LastLogContains("level", "warn") {
		t.Error("The log is not written with warn level")
	}
}

func TestShouldDisplayTheErrorLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	defer func(dumbWriter *DumbWriter) {
		recover()
		if !dumbWriter.LastLogContains("level", "error") {
			t.Error("The log is not written with error level")
		}
	}(dumbWriter)

	golog.Error("Test")
}

func TestShouldDisplayTheFatalLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Fatal("Test")

	if !dumbWriter.LastLogContains("level", "fatal") {
		t.Error("The log is not written with fatal level")
	}
}

func TestShouldDisplayTheDebugLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Debug("Test")

	if !dumbWriter.LastLogContains("message", "\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func TestShouldDisplayTheInfoLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Info("Test")

	if !dumbWriter.LastLogContains("message", "\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func TestShouldDisplayTheWarnLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Warn("Test")

	if !dumbWriter.LastLogContains("message", "\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func TestShouldDisplayTheErrorLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	defer func(dumbWriter *DumbWriter) {
		recover()

		if !dumbWriter.LastLogContains("message", "\"Test\"") {
			t.Errorf("The log message is not written")
		}
	}(dumbWriter)

	golog.Error("Test")
}

func TestShouldDisplayTheFatalLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Fatal("Test")

	if !dumbWriter.LastLogContains("message", "\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func fakeMetadata() entries.Fields {
	return entries.Fields{
		"test": "oui",
		"otherData": 42.42,
	}
}

func checkIfFakeMetadataAreLogged(t *testing.T, dumbWriter *DumbWriter, level string) {
	if !dumbWriter.LastLogContains("test", "\"oui\"") || !dumbWriter.LastLogContains("otherData", "42.42") {
		t.Errorf("The %v log is not written with the provided metadata %v", level, dumbWriter.messages[len(dumbWriter.messages) - 1])
	}
}

func TestShouldDisplayTheDebugMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.WithFields(fakeMetadata()).Debug("Test")

	checkIfFakeMetadataAreLogged(t, dumbWriter, "debug")
}

func TestShouldDisplayTheInfoMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.WithFields(fakeMetadata()).Info("Test")

	checkIfFakeMetadataAreLogged(t, dumbWriter, "info")
}

func TestShouldDisplayTheWarnMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.WithFields(fakeMetadata()).Warn("Test")

	checkIfFakeMetadataAreLogged(t, dumbWriter, "warn")
}

func TestShouldDisplayTheErrorMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	defer func(dumbWriter *DumbWriter) {
		recover()

		checkIfFakeMetadataAreLogged(t, dumbWriter, "error")
	}(dumbWriter)

	golog.WithFields(fakeMetadata()).Error("Test")
}

func TestShouldDisplayTheFatalMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.WithFields(fakeMetadata()).Fatal("Test")

	checkIfFakeMetadataAreLogged(t, dumbWriter, "fatal")
}

func TestShouldDisplayChainedMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()
	
	golog.WithFields(entries.Fields{
		"first": "a metadata",
	}).WithFields(fakeMetadata()).Debug("Test")
	
	if !dumbWriter.LastLogContains("test", "\"oui\"") ||
		!dumbWriter.LastLogContains("otherData", "42.42") ||
		!dumbWriter.LastLogContains("first", "\"a metadata\"") {
		t.Errorf("The log is not written with the provided chained metadata %v", dumbWriter.messages[len(dumbWriter.messages) - 1])
	}
}

func TestShouldAllowForMultipleOutputs(t *testing.T) {
	firstOutput := NewDumbWriter()
	secondOutput := NewDumbWriter()
	formatter := formatters.NewKvpFormatter()
	golog.SetLoggers(
		loggers.NewLogger(firstOutput, formatter),
		loggers.NewLogger(secondOutput, formatter),
	)

	golog.Debug("Test")

	if len(firstOutput.messages) == 0 || len(secondOutput.messages) == 0 {
		t.Errorf("The log is not written in every output")
	}
}

const (
	FIRST_FORMATTER_USED = "FIRST_FORMATTER_USED"
	SECOND_FORMATTER_USED = "SECOND_FORMATTER_USED"
)

type FirstFormatter struct {}

func (formatter FirstFormatter) Format(fields entries.Fields, level string, message string) string {
	return FIRST_FORMATTER_USED
}

type SecondFormatter struct {}

func (formatter SecondFormatter) Format(fields entries.Fields, level string, message string) string {
	return SECOND_FORMATTER_USED
}

func TestShouldAllowForMultipleFormatters(t *testing.T) {
	firstOutput := NewDumbWriter()
	secondOutput := NewDumbWriter()
	firstFormatter := FirstFormatter{}
	secondFormatter := SecondFormatter{}
	golog.SetLoggers(
		loggers.NewLogger(firstOutput, firstFormatter),
		loggers.NewLogger(secondOutput, secondFormatter),
	)

	golog.Debug("Test")

	if len(firstOutput.messages) == 0 || firstOutput.messages[0] != FIRST_FORMATTER_USED {
		t.Errorf("The first formatter was not used")
	}

	if len(secondOutput.messages) == 0 || secondOutput.messages[0] != SECOND_FORMATTER_USED {
		t.Errorf("The second formatter was not used")
	}
}
