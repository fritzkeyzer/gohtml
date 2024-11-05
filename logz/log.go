package logz

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/fatih/color"
)

var Level slog.Level = slog.LevelInfo

//LevelDebug Level = -4
//	LevelInfo  Level = 0
//	LevelWarn  Level = 4
//	LevelError Level = 8

func Debug(msg string, pairs ...any) {
	if Level > slog.LevelDebug {
		return
	}
	color.New(color.FgCyan).Printf("DEBUG: %s", msg)
	color.New(color.Reset).Println(formatPairs(pairs...))
}

func Info(msg string, pairs ...any) {
	if Level > slog.LevelInfo {
		return
	}
	color.New(color.FgGreen).Printf("INFO: %s", msg)
	color.New(color.Reset).Println(formatPairs(pairs...))
}

func Warn(msg string, pairs ...any) {
	if Level > slog.LevelWarn {
		return
	}
	color.New(color.FgHiYellow).Printf("WARN: %s", msg)
	color.New(color.Reset).Println(formatPairs(pairs...))
}

func Error(err error, msg string, pairs ...any) {
	if Level > slog.LevelError {
		return
	}
	color.New(color.FgRed).Printf("ERROR: %s: %v", msg, err)
	color.New(color.Reset).Println(formatPairs(pairs...))
}

func formatPairs(pairs ...any) string {
	line := ""
	for i := 0; i < len(pairs); i += 2 {

		line += fmt.Sprintf("\n\t%v: %s", pairs[i],
			strings.TrimSpace(fmt.Sprintf("%+v", pairs[i+1])),
		)
	}

	return line
}
