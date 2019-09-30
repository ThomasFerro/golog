package configurations_test

import (
	"os"
	"testing"

	gologgers "github.com/ThomasFerro/golog/loggers"
	configurations "github.com/ThomasFerro/golog/configurations"
	"github.com/ThomasFerro/golog/outputs"
)

func TestShouldProvideADefaultConfigurationWithStd(t *testing.T) {
	loggers := configurations.GetLoggers()

	if len(loggers) < 1 {
		t.Error("Could not find the standard output logger")
		return
	}

	if loggers[0].Output() != os.Stdout {
		t.Error("The first default logger's output is not the standard output")
	}
}

func TestShouldProvideASecondDefaultConfigurationWithLogFile(t *testing.T) {
	loggers := configurations.GetLoggers()

	if len(loggers) < 2 {
		t.Errorf("Should have two default configurations but found %v", len(loggers))
		return
	}

	if _, typeOk := loggers[1].Output().(outputs.SimpleLogFile); !typeOk {
		t.Error("The second default logger's output is a log file writer")
	}
}

func TestSecondDefaultLoggerShouldBeNamedLogByDefault(t *testing.T) {
	loggers := configurations.GetLoggers()

	if len(loggers) < 2 {
		t.Errorf("Should have two default configurations but found %v", len(loggers))
		return
	}

	simpleLogFile, typeOk := loggers[1].Output().(outputs.SimpleLogFile);
	if !typeOk {
		t.Error("The second default logger's output is a log file writer")
		return
	}

	if simpleLogFile.Filename() != "log" {
		t.Errorf("The rotating log file has the wrong name")
	}
}

// TODO : Test that the default SimpleLogFile's filename is based on env var

func TestShouldUseTheProvidedLoggers(t *testing.T) {
	configurations.SetLoggers(
		gologgers.NewLogger(os.Stderr, nil),
	)

	loggers := configurations.GetLoggers()

	if len(loggers) < 1 {
		t.Error("Could not find the provided logger")
		return
	}

	if loggers[0].Output() != os.Stderr {
		t.Error("The provided logger was not set")
	}
}
