package slogx

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

var RequestIdKey struct{}

type HandlerRequestID struct {
	slog.Handler
}

func (h HandlerRequestID) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(RequestIdKey).(string); ok {
		r.Add("request_id", slog.StringValue(requestID))
	}
	return h.Handler.Handle(ctx, r)
}

// GetLogLevel returns slog.Level by level string
func GetLogLevel(level string) slog.Level {
	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	return logLevel
}

// GetLogger returns a slog.Logger with HandlerRequestID
// and slog.Source replaced by relative path by logger level string
func GetLogger(level string) *slog.Logger {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var replacer = func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			if file, ok := strings.CutPrefix(source.File, wd); ok {
				source.File = file
			}
		}
		return a
	}

	options := &slog.HandlerOptions{
		Level:       GetLogLevel(level),
		AddSource:   true,
		ReplaceAttr: replacer,
	}

	handler := HandlerRequestID{Handler: slog.NewJSONHandler(os.Stderr, options)}
	logger := slog.New(handler).With()

	return logger
}

// SetDefault sets slog.DefaultLogger
func SetDefault(logger *slog.Logger) {
	slog.SetDefault(logger)
}
