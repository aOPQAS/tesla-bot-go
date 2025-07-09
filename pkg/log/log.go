package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *Logger

func init() {
	defaultLogger = getDefaultLogger()
}

func NewMustLogger(encoding, level string) *Logger {
	logger, err := NewLogger(encoding, level)
	if err != nil {
		panic(err)
	}

	return logger
}

func NewLogger(encoding, level string) (*Logger, error) {
	lvl := zapcore.InfoLevel
	if err := lvl.Set(level); err != nil {
		return nil, fmt.Errorf("failed to parse level: %w", err)
	}

	l := &Logger{
		encoding: encoding,
		level:    lvl,
		writer:   os.Stderr,
	}

	encoder, err := getEncoder(l.encoding)
	if err != nil {
		return nil, fmt.Errorf("failed to parse encoder: %w", err)
	}

	l.Logger = zap.New(zapcore.NewTee(
		// info / debug log level -> stdout
		zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level == zapcore.InfoLevel || level == zapcore.DebugLevel
			}),
		),

		// error / fatal log level -> stderr
		zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stderr),
			zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level == zapcore.ErrorLevel || level == zapcore.FatalLevel
			}),
		),
	))

	return l, nil
}

type Logger struct {
	encoding string
	level    zapcore.Level
	writer   io.Writer

	*zap.Logger
}

func (l *Logger) NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, l, fields)
}

func (l *Logger) clone() *Logger {
	logger := *l
	return &logger
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	if ctx == nil {
		return l
	}

	if fields, ok := ctx.Value(l).([]zap.Field); ok && fields != nil {
		logger := l.clone()
		logger.Logger = l.With(fields...)
		return logger
	}

	return l
}

var (
	logLevel    = "INFO"
	logEncoding = "console"
)

func GetLogger() *Logger {
	return defaultLogger
}

func getDefaultLogger() *Logger {
	logger := NewMustLogger(logEncoding, logLevel)
	replaceOtherLoggers(logger)
	return logger
}

func replaceOtherLoggers(logger *Logger) {
	zap.ReplaceGlobals(logger.Logger)

	if _, err := zap.RedirectStdLogAt(logger.Logger, zapcore.InfoLevel); err != nil {
		panic(err)
	}
}

func getEncoder(encoding string) (zapcore.Encoder, error) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:      "caller",
		EncodeCaller:   customEncodeCaller,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	switch encoding {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig), nil
	case "console":
		return zapcore.NewConsoleEncoder(encoderConfig), nil
	default:
		return nil, fmt.Errorf("failed to find encoder: %q", encoding)
	}
}

func SetLogLevel(level string) {
	logLevel = level
	defaultLogger = getDefaultLogger()
}

func SetLogEncoding(enc string) {
	logEncoding = enc
	defaultLogger = getDefaultLogger()
}

func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	return defaultLogger.NewContext(ctx, fields...)
}

func WithContext(ctx context.Context) *Logger {
	return defaultLogger.WithContext(ctx)
}

func Debug(msg string, fields ...zap.Field) { defaultLogger.Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)  { defaultLogger.Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { defaultLogger.Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field) { defaultLogger.Error(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { defaultLogger.Fatal(msg, fields...) }
func Panic(msg string, fields ...zap.Field) { defaultLogger.Panic(msg, fields...) }

func customEncodeCaller(_ zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	file, _, line := findCaller()
	enc.AppendString(file + ":" + strconv.Itoa(line))
}

func findCaller() (string, string, int) {
	const callerSkip = 8

	var (
		pc       uintptr
		file     string
		function string
		line     int
	)

	pc, file, line = getCaller(callerSkip)

	if pc != 0 {
		frames := runtime.CallersFrames([]uintptr{pc})
		frame, _ := frames.Next()
		function = frame.Function
	}

	return file, function, line
}

func getCaller(skip int) (uintptr, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return 0, "", 0
	}

	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}

	return pc, file, line
}
