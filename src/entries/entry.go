package entries

// Entry Log entry
type Entry interface {
	Fields() Fields
	WithFields(fields Fields) Entry
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	WriteLog(level string, message string)
}
