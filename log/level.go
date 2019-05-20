package log

import "go.uber.org/zap/zapcore"

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func setLevel(level Level) {
	switch level {
	case DebugLevel:
		atom.SetLevel(zapcore.DebugLevel)
	case InfoLevel:
		atom.SetLevel(zapcore.InfoLevel)
	case WarnLevel:
		atom.SetLevel(zapcore.WarnLevel)
	case ErrorLevel:
		atom.SetLevel(zapcore.ErrorLevel)
	case FatalLevel:
		atom.SetLevel(zapcore.FatalLevel)
	default:
		atom.SetLevel(zapcore.InfoLevel)
	}
}
