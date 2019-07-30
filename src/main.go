package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// FormatLog Format the log message
func FormatLog(level string, message string, metadata interface {}) string {
	typeOfMetadata := reflect.TypeOf(metadata)
	valueOfMetadata := reflect.ValueOf(metadata)

	formattedMetadata := make([]string, typeOfMetadata.NumField() + 1)

	formattedMetadata[0] = fmt.Sprintf("level=%v", level)
	for i := 0; i < typeOfMetadata.NumField(); i++ {
		field := typeOfMetadata.Field(i)
		value := fmt.Sprint(valueOfMetadata.Field(i))

		if (field.Type.String() == "string") {
			value = fmt.Sprintf("\"%v\"", value)
		}

		formattedMetadata[i + 1] = fmt.Sprintf("%v=%v", field.Name, value)
	}

	return fmt.Sprintf("%v -- %v", message, strings.Join(formattedMetadata, " "))
}

// Debug Log a debug message
func Debug(message string) {
	var metadata struct{}
	DebugWithMetadata(message, metadata)
}

// DebugWithMetadata Log a debug message with metadata
func DebugWithMetadata(message string, metadata interface {}) {
	log.Println(
		FormatLog("debug", message, metadata),
	)
}

// Info Log an info message
func Info(message string) {
	var metadata struct{}
	InfoWithMetadata(message, metadata)
}

// InfoWithMetadata Log a info message with metadata
func InfoWithMetadata(message string, metadata interface {}) {
	log.Println(
		FormatLog("info", message, metadata),
	)
}

// Warn Log a warning message
func Warn(message string) {
	var metadata struct{}
	WarnWithMetadata(message, metadata)
}

// WarnWithMetadata Log a warn message with metadata
func WarnWithMetadata(message string, metadata interface {}) {
	log.Println(
		FormatLog("warn", message, metadata),
	)
}

// Error Log an error message
func Error(message string) {
	var metadata struct{}
	ErrorWithMetadata(message, metadata)
}

// ErrorWithMetadata Log a error message with metadata
func ErrorWithMetadata(message string, metadata interface {}) {
	log.Println(
		FormatLog("error", message, metadata),
	)
}

// Fatal Log a fatal message
func Fatal(message string) {
	var metadata struct{}
	FatalWithMetadata(message, metadata)
	log.Println("level=fatal")
}

// FatalWithMetadata Log a fatal message with metadata
func FatalWithMetadata(message string, metadata interface {}) {
	log.Println(
		FormatLog("fatal", message, metadata),
	)
}
