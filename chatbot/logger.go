package chatbot

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var onceLogger sync.Once

func initLogger(level zapcore.Level, isConsole bool, logpath string) *zap.Logger {
	loglevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	if isConsole {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebugging := zapcore.Lock(os.Stdout)
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleDebugging, loglevel),
		)

		cl := zap.New(core)
		// defer cl.Sync()
		return cl
	}

	return nil
}

// InitLogger - initializes a thread-safe singleton logger
func InitLogger(level zapcore.Level, isConsole bool, logpath string) {
	// once ensures the singleton is initialized only once
	onceLogger.Do(func() {
		logger = initLogger(level, isConsole, logpath)
	})

	return
}

// // Log a message at the given level with given fields
// func Log(level zap.Level, message string, fields ...zap.Field) {
// 	singleton.Log(level, message, fields...)
// }

// Debug logs a debug message with the given fields
func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

// Info logs a debug message with the given fields
func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

// Warn logs a debug message with the given fields
func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

// Error logs a debug message with the given fields
func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

// Fatal logs a message than calls os.Exit(1)
func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

// SyncLogger - sync logger
func SyncLogger() {
	logger.Sync()
}
