package main_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	golog "github.com/ThomasFerro/golog"
)

/*
- [x] Afficher le niveau de log
- [x] Afficher les métadonnées
- [x] Afficher le message de log (fait parti des métadonnées ?)
- [x] Appel panic sur erreur
- [ ] Appel os.Exit sur fatal (comment le tester proprement et garder le code coverage ?)
- [x] Afficher les métadonnées imbriquées
- [x] Afficher les métadonnées multilignes
QOL:
- [ ] Afficher les logs debug en bleu par défaut
- [ ] Afficher les logs info en vert par défaut
- [ ] Afficher les logs warn en jaune/orange par défaut
- [ ] Afficher les logs error en rouge par défaut
- [ ] Afficher les logs fatal en rouge par défaut
- [ ] Permettre de surcharger les couleurs par défaut
*/
type DumbWriter struct { 
	messages []string;
}

func (d *DumbWriter) Write(p []byte) (int, error) {
	d.messages = append(d.messages, string(p))
	return len(p), nil
}

func (d DumbWriter) LastLogContains(message string) bool {
	lastLog := d.messages[len(d.messages) - 1]
	fmt.Println(lastLog)
	return strings.Contains(lastLog, message)
}

func NewDumbWriter() *DumbWriter {
	return &DumbWriter{
		messages: make([]string, 0),
	}
}

func NewDumbWriterAsLogOutput() *DumbWriter {
	dumbWriter := NewDumbWriter()
	log.SetOutput(dumbWriter)
	return dumbWriter
}

func TestShouldDisplayTheDebugLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Debug("Test")
	
	if !dumbWriter.LastLogContains("level=debug") {
		t.Error("The log is not written with debug level")
	}
}

func TestShouldDisplayTheInfoLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Info("Test")

	if !dumbWriter.LastLogContains("level=info") {
		t.Error("The log is not written with info level")
	}
}

func TestShouldDisplayTheWarnLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Warn("Test")

	if !dumbWriter.LastLogContains("level=warn") {
		t.Error("The log is not written with warn level")
	}
}

func TestShouldDisplayTheErrorLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	defer func(dumbWriter *DumbWriter) {
		recover()
		if !dumbWriter.LastLogContains("level=error") {
			t.Error("The log is not written with error level")
		}
	}(dumbWriter)

	golog.Error("Test")
}

func TestShouldDisplayTheFatalLogLevel(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Fatal("Test")

	if !dumbWriter.LastLogContains("level=fatal") {
		t.Error("The log is not written with fatal level")
	}
}

func TestShouldDisplayTheDebugLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Debug("Test")

	if !dumbWriter.LastLogContains("message=\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func TestShouldDisplayTheInfoLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Info("Test")

	if !dumbWriter.LastLogContains("message=\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func TestShouldDisplayTheWarnLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Warn("Test")

	if !dumbWriter.LastLogContains("message=\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func TestShouldDisplayTheErrorLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	defer func(dumbWriter *DumbWriter) {
		recover()

		if !dumbWriter.LastLogContains("message=\"Test\"") {
			t.Errorf("The log message is not written")
		}
	}(dumbWriter)

	golog.Error("Test")
}

func TestShouldDisplayTheFatalLogMessage(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()

	golog.Fatal("Test")

	if !dumbWriter.LastLogContains("message=\"Test\"") {
		t.Errorf("The log message is not written")
	}
}

func fakeMetadata() golog.Fields {
	return golog.Fields{
		"test": "oui",
		"otherData": 42.42,
	}
}

func checkIfFakeMetadataAreLogged(t *testing.T, dumbWriter *DumbWriter, level string) {
	if !dumbWriter.LastLogContains("test=\"oui\"") || !dumbWriter.LastLogContains("otherData=42.42") {
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

func TestShouldDisplayAStringMetadataQuoted(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()
	
	golog.WithFields(golog.Fields{
		"test": "a metadata",
	}).Debug("Test")

	if !dumbWriter.LastLogContains("test=\"a metadata\"") {
		t.Errorf("The log is not written with the quoted string metadata")
	}
}

func TestShouldDisplayMultilineStringMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()
	
	golog.WithFields(golog.Fields{
		"test": `a test 
		metadata`,
	}).Debug("Test")

	if !dumbWriter.LastLogContains(`test="a test 
		metadata"`) {
		t.Errorf("The log is not written with the multiline string metadata")
	}
}

func TestShouldDisplayChainedMetadata(t *testing.T) {
	dumbWriter := NewDumbWriterAsLogOutput()
	
	golog.WithFields(golog.Fields{
		"first": "a metadata",
	}).WithFields(fakeMetadata()).Debug("Test")
	
	if !dumbWriter.LastLogContains("test=\"oui\"") ||
		!dumbWriter.LastLogContains("otherData=42.42") ||
		!dumbWriter.LastLogContains("first=\"a metadata\"") {
		t.Errorf("The log is not written with the provided chained metadata %v", dumbWriter.messages[len(dumbWriter.messages) - 1])
	}
}

func TestErrorLogShouldPanic(t *testing.T) {
	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The error log did not panic")
        }
	}()
	
	golog.Error("Test")
}
