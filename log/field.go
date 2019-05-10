package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

func Int(key string, val int) Field {
	return zap.Int(key, val)
}

func String(key string, val string) Field {
	return zap.String(key, val)
}
