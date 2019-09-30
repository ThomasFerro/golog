package outputs

import (
	"fmt"
	"io"
	"os"
)

// LogFile A log file
type LogFile interface {
	io.Writer
	Path() string
}

// SimpleLogFile A simple log file
type SimpleLogFile struct {
	folder   string
	filename string
}

func (r SimpleLogFile) Write(b []byte) (n int, err error) {
	file, err := os.OpenFile(r.Path(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}

	messageToWrite := []byte(fmt.Sprintln(string(b)))

	if _, err := file.Write(messageToWrite); err != nil {
		file.Close()
		return 0, err
	}

	if err := file.Close(); err != nil {
		return 0, err
	}

	return len([]byte(messageToWrite)), nil
}

// Filename Get the log file name
func (r SimpleLogFile) Filename() string {
	return r.filename
}

// Path get the path to the file
func (r SimpleLogFile) Path() string {
	if len(r.folder) > 0 {
		return fmt.Sprintf("%v%v.log", r.folder, r.Filename())
	}
	return fmt.Sprintf("/var/log/%v.log", r.Filename())
}

// NewSimpleLogFile Create a new SimpleLogFile
func NewSimpleLogFile(filename string) LogFile {
	return SimpleLogFile{
		folder:   "",
		filename: filename,
	}
}

// NewSimpleLogFileWithFolderPath Create a new SimpleLogFile with a custom folder
func NewSimpleLogFileWithFolderPath(folder string, filename string) LogFile {
	return SimpleLogFile{
		folder,
		filename,
	}
}
