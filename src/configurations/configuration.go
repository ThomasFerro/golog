package configurations

import (
	"os"

	"github.com/ThomasFerro/golog/formatters"
	gologgers "github.com/ThomasFerro/golog/loggers"
	"github.com/ThomasFerro/golog/outputs"
)

func getLoggerName() string {
	if filename := os.Getenv("LOGFILE_NAME"); filename == "" {
		return filename
	}

	return "log"
}

var defaultLoggers = []gologgers.Logger{
	gologgers.NewLogger(os.Stdout, formatters.NewKvpFormatter()),
	gologgers.NewLogger(outputs.NewSimpleLogFile(
		"log",
	), formatters.NewKvpFormatter()),
}

var loggers []gologgers.Logger

// SetLoggers Set the loggers to use
func SetLoggers(newLoggers ...gologgers.Logger) {
	loggers = newLoggers
}

// GetLoggers Get the used loggers
func GetLoggers() []gologgers.Logger {
	if len(loggers) == 0 {
		return defaultLoggers
	}
	return loggers
}
