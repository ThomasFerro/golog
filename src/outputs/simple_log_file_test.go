package outputs_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ThomasFerro/golog/outputs"
)

func removeFile(path string) {
	os.Remove(path);
}

func TestProvideAPathToTheLogFile(t *testing.T) {
	logFile := outputs.NewSimpleLogFile("my-application")

	if logFile.Path() != "/var/log/my-application.log" {
		t.Error("The default log file path is incorrect")
	}
}

func TestCreateALogFile(t *testing.T) {
	logFile := outputs.NewSimpleLogFileWithFolderPath("./", "test1")
	removeFile(logFile.Path())

	_, writeErr := logFile.Write([]byte("Dumb log line"))

	if _, err := os.Stat(logFile.Path()); os.IsNotExist(err) {
		t.Errorf("The log file does not exist, error : %v", writeErr)
	}

	removeFile(logFile.Path())
}

func TestWriteTheLineInTheFile(t *testing.T) {
	logFile := outputs.NewSimpleLogFileWithFolderPath("./", "test2")

	_, writeErr := logFile.Write([]byte("Dumb log line"))

	if writeErr != nil {
		t.Errorf("Could not write in the log file : %v", writeErr)
	}

	content, readErr := ioutil.ReadFile(logFile.Path())

	if readErr != nil {
		t.Errorf("Could not read the log file : %v", readErr)
	}

	if string(content) != "Dumb log line\n" {
		t.Errorf("The log file content does not match with what was written: %v", string(content))
	}

	removeFile(logFile.Path())
}

func TestWriteTheNextLineAtTheEndOfTheFile(t *testing.T) {
	logFile := outputs.NewSimpleLogFileWithFolderPath("./", "test3")

	_, _ = logFile.Write([]byte("Dumb log line"))
	_, writeErr := logFile.Write([]byte("Second log line"))

	if writeErr != nil {
		t.Errorf("Could not write in the log file : %v", writeErr)
	}

	content, readErr := ioutil.ReadFile(logFile.Path())

	if readErr != nil {
		t.Errorf("Could not read the log file : %v", readErr)
	}

	if string(content) != "Dumb log line\nSecond log line\n" {
		t.Errorf("The log file content does not match with what was written: %v", string(content))
	}

	removeFile(logFile.Path())
}
