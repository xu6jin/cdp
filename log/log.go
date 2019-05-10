package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

func Debug(v interface{}, fields ...Field) {
	logger.Debug(fmt.Sprint(v), fields...)
}

func Info(v interface{}, fields ...Field) {
	logger.Info(fmt.Sprint(v), fields...)
}

func Warn(v interface{}, fields ...Field) {
	logger.Warn(fmt.Sprint(v), fields...)
}

func Error(v interface{}, fields ...Field) {
	logger.Error(fmt.Sprint(v), fields...)
}

func Fatal(v interface{}, fields ...Field) {
	logger.Fatal(fmt.Sprint(v), fields...)
}

var atom zap.AtomicLevel
var logger *zap.Logger
var once sync.Once

func SetLevel(level Level) {
	setLevel(level)
}

func Sync() error {
	return logger.Sync()
}

func RegisterLogger() {
	once.Do(func() {
		atom = zap.NewAtomicLevel()
		logger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.Lock(os.Stdout),
			atom,
		),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		)
	})
}
