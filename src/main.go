package main

import (
	"log"

	"github.com/ThomasFerro/golog/entries"
	"github.com/ThomasFerro/golog/formatters"
)

// LogEntry Log entry
type LogEntry struct {
	fields entries.Fields
}

// NewLogEntry Create a new log entry
func NewLogEntry() entries.Entry {
	return LogEntry{}
}

// WithFields Create a new entry with provided fields
func WithFields(fields entries.Fields) entries.Entry {
	return LogEntry{
		fields,
	}
}

// Fields Get the log's fields
func (entry LogEntry) Fields() entries.Fields {
	return entry.fields
}

// WithFields Add fields to the entry
func (entry LogEntry) WithFields(fields entries.Fields) entries.Entry {
	logEntry := LogEntry{
		fields,
	}
	for key, value := range entry.fields {
		logEntry.fields[key] = value
	}

	return logEntry
}

// Debug Log a debug message
func (entry LogEntry) Debug(message string) {
	entry.WriteLog("debug", message)
}

// Info Log an info message
func (entry LogEntry) Info(message string) {
	entry.WriteLog("info", message)
}

// Warn Log a warning message
func (entry LogEntry) Warn(message string) {
	entry.WriteLog("warn", message)
}

// Error Log an error message
func (entry LogEntry) Error(message string) {
	entry.WriteLog("error", message)
}

// Fatal Log a fatal message
func (entry LogEntry) Fatal(message string) {
	// TODO : Passer par log.Fatal
	entry.WriteLog("fatal", message)
}

// WriteLog Write the log
func (entry LogEntry) WriteLog(level string, message string) {
	// TODO : Loop in the loggers
	// TODO : Use the configured loggers' formatters
	formattedMessage := formatters.NewKvpFormatter().Format(entry, level, message)
	if level == "error" {
		log.Panicln(formattedMessage)
	}
	log.Println(formattedMessage)
}

// Debug Log a debug message
func Debug(message string) {
	NewLogEntry().Debug(message)
}

// Info Log an info message
func Info(message string) {
	NewLogEntry().Info(message)
}

// Warn Log a warning message
func Warn(message string) {
	NewLogEntry().Warn(message)
}

// Error Log an error message
func Error(message string) {
	NewLogEntry().Error(message)
}

// Fatal Log a fatal message
func Fatal(message string) {
	NewLogEntry().Fatal(message)
}
