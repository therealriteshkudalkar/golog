package golog

import (
	"log/slog"
	"os"
)

func NewLogHandler(logLevel slog.Level, writeToFile bool, logFilePath string) *LogHandler {
	// Open the file
	var logHandler = &LogHandler{
		level:       logLevel,
		writeToFile: writeToFile,
		fileWriter:  nil,
	}
	if writeToFile {
		logFileWriter, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			slog.Error("Failed to open the log file.", "Error", err)
			return logHandler
		}
		logHandler.fileWriter = logFileWriter
	}
	return logHandler
}
