package golog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/therealriteshkudalkar/golog/constants"
)

type LogHandler struct {
	level       slog.Level
	writeToFile bool
	fileWriter  *os.File
}

func (logHandler LogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return logHandler.level <= level
}

func (logHandler LogHandler) Handle(ctx context.Context, record slog.Record) error {
	message := record.Message

	record.Attrs(func(attr slog.Attr) bool {
		message += fmt.Sprintf(" %v: %v", attr.Key, attr.Value)
		return true
	})

	var stamp string
	var writer io.Writer = os.Stderr
	if logHandler.writeToFile {
		if logHandler.fileWriter != nil {
			writer = logHandler.fileWriter
		}
		stamp = fmt.Sprintf("[%v]", record.Level)
	} else {
		switch record.Level {
		case slog.LevelDebug:
			stamp = fmt.Sprintf("[%v%v%v]", constants.Blue, record.Level, constants.Reset)
		case slog.LevelInfo:
			stamp = fmt.Sprintf("[%v%v%v]", constants.Green, record.Level, constants.Reset)
		case slog.LevelWarn:
			stamp = fmt.Sprintf("[%v%v%v]", constants.Yellow, record.Level, constants.Reset)
		case slog.LevelError:
			stamp = fmt.Sprintf("[%v%v%v]", constants.Red, record.Level, constants.Reset)
		default:
			panic("unreachable")
		}
	}

	timestamp := record.Time.Format(time.RFC3339)
	count, err := fmt.Fprintf(writer, "%v %v %v\n", timestamp, stamp, message)
	if err != nil {
		println("Write Count: ", count, ", Error: ", err.Error())
	}
	return err
}

func (logHandler LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("implement me")
}

func (logHandler LogHandler) WithGroup(name string) slog.Handler {
	panic("implement me")
}

func (logHandler LogHandler) CloseLoggerWriter() {
	if logHandler.fileWriter == nil {
		slog.Error("Could not find file writer to close.")
		return
	}
	err := logHandler.fileWriter.Close()
	if err != nil {
		slog.Error("Error while closing the log file.", "Error", err)
		return
	}
}
