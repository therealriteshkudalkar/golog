package tests

import (
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/therealriteshkudalkar/golog"
)

func TestGoLogFileWriter(t *testing.T) {
	// Describe the log file path
	logFilePath := "./test.log"

	// Set the logger file
	logHandler := golog.NewLogHandler(slog.LevelDebug, true, logFilePath)
	slog.SetDefault(slog.New(logHandler))

	// Create a log
	slog.Debug("This is a test.", "Key", "Value")
	logHandler.CloseLoggerWriter()

	// Read the file
	data, err := os.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("Failed to open the file.")
	}
	if !strings.HasSuffix(string(data), "This is a test. Key: Value\n") {
		t.Fatalf("The data is not same!")
	}

	// Delete the file
	err = os.Remove(logFilePath)
	if err != nil {
		println("Failed to delete the file. Error", err)
	}
}
