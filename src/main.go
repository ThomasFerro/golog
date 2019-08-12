package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Fields Log fields
type Fields map[string]interface{}

// Entry Log entry
type Entry interface {
	WithFields(fields Fields) Entry
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	WriteLog(level string, message string)
}

// LogEntry Log entry
type LogEntry struct {
	fields Fields
}

// NewLogEntry Create a new log entry
func NewLogEntry() Entry {
	return LogEntry{}
}

// WithFields Create a new entry with provided fields
func WithFields(fields Fields) Entry {
	return LogEntry{
		fields,
	}
}

// WithFields Add fields to the entry
func (entry LogEntry) WithFields(fields Fields) Entry {
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
	log.Panicln(entry.formatLog("error", message))
}

// Fatal Log a fatal message
func (entry LogEntry) Fatal(message string) {
	// TODO : Passer par log.Fatal
	entry.WriteLog("fatal", message)
}

func formatStringMessage(key string, value interface{}) string {
	return fmt.Sprintf("%v=\"%v\"", key, value)
}

func (entry LogEntry) formatLog(level string, message string) string {
	formattedMetadata := make([]string, len(entry.fields) + 2)
	index := 2

	formattedMetadata[0] = fmt.Sprintf("level=%v", level)
	formattedMetadata[1] = formatStringMessage("message", message)
	for key, value := range entry.fields {
		if reflect.TypeOf(value).String() == "string" {
			formattedMetadata[index] = formatStringMessage(key, value)
		} else {
			formattedMetadata[index] = fmt.Sprintf("%v=%v", key, value)
		}
		index++
	}

	return strings.Join(formattedMetadata, " ") 
}

// WriteLog Write the log
func (entry LogEntry) WriteLog(level string, message string) {
	// TODO : Plusieurs sorties ? Avec un format texte avec couleur pour le terminal et un json ?
	log.Println(entry.formatLog(level, message))
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
