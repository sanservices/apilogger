package apilogger

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

type CtxKeys struct {
	// TaskName is the name of the scheduled task
	// where the log originated from
	TaskName string

	//UUID is the unique ID to identify a
	// session of the apilogger and its logs
	UUID string

	// StartTime is the context key used to
	// access start time of transaction
	StartTime time.Time
}
type ContextKey string

const (
	//ContextData is the struct that contains
	//the data from the context
	ContextData ContextKey = "context-data"

	// Depth of the callstack - needed to determine
	// the initial caller function
	depth int = 6

	prefixInfo  = "INFO "
	prefixWarn  = "WARN "
	prefixError = "ERROR "
	prefixFatal = "FATAL "
)

// Logger struct
type Logger struct {
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	output     io.Writer
	errOutput  io.Writer

	// requestID  string
	// apiKey     string
	// remoteAddr string
	// session    string
}

type Fields map[string]interface{}

var defaultLogger *Logger

// New returns a new Logger instance.
func New() *Logger {
	defaultLogger = &Logger{
		output:    os.Stdout,
		errOutput: os.Stderr,
	}
	return defaultLogger
}

// Set logger to file
func (l *Logger) SetOutputFile(outPath string) error {
	file, err := os.OpenFile(
		outPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Println("Failed to open log file", err)
		return err
	}

	l.output = io.MultiWriter(file, l.output)
	l.errOutput = io.MultiWriter(file, l.errOutput)
	return nil
}

// prints message.
func (l *Logger) printlnWF(logger *log.Logger, logCat LogCat, startTime time.Time, taskName, uuId string, status StatusCat, fields *Fields) {
	logger.Println(finalMessageWF(logCat, startTime, taskName, uuId, status, fields))
}

func (l *Logger) println(logger *log.Logger, logCat LogCat, startTime time.Time, taskName, uuId string, status StatusCat, v ...interface{}) {
	logger.Println(finalMessage(logCat, startTime, taskName, uuId, status, v...))
}

func (l *Logger) printlnf(logger *log.Logger, logCat LogCat, startTime time.Time, taskName, uuId string, status StatusCat, format string, v ...interface{}) {
	logger.Println(finalMessagef(logCat, startTime, taskName, uuId, status, format, v...))
}

func (l *Logger) Info(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.println(l.infoLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, v...)
}

func (l *Logger) Infof(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.printlnf(l.infoLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, format, v...)
}

func (l *Logger) InfoWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.printlnWF(l.infoLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, fields)
}

func (l *Logger) Printf(status StatusCat, s string, i ...interface{}) {
	l.Infof(context.TODO(), LogCatDebug, status, s, i...)
}

func (l *Logger) Warn(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.println(l.warningLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, v...)
}

func (l *Logger) Warnf(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.printlnf(l.warningLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, format, v...)
}

func (l *Logger) WarnWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.printlnWF(l.warningLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, fields)
}

func (l *Logger) Error(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.println(l.errorLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, v...)
}

func (l *Logger) Errorf(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.printlnf(l.errorLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, format, v...)
}

func (l *Logger) ErrorWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.printlnWF(l.errorLog, logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, fields)
}

func (l *Logger) Fatal(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessage(logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, v...))
}

func (l *Logger) Fatalf(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessagef(logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, format, v...))
}

func (l *Logger) FatalWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	// Extract contextual values
	contextData, _ := ctx.Value(ContextData).(CtxKeys)

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessageWF(logCat, contextData.StartTime, contextData.TaskName, contextData.UUID, status, fields))
}

// Info prints message with logging level of info
func Info(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	defaultLogger.Info(ctx, logCat, status, v...)
}

// Infof prints a message using the specified format.
func Infof(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	defaultLogger.Infof(ctx, logCat, status, format, v...)
}

// InfoWF prints message using Fields struct to pass multiple key=value pairs.
func InfoWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	defaultLogger.InfoWF(ctx, logCat, status, fields)
}

// Warn prints message with logging level of info
func Warn(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	defaultLogger.Warn(ctx, logCat, status, v...)
}

// Warnf prints a message using the specified format.
func Warnf(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	defaultLogger.Warnf(ctx, logCat, status, format, v...)
}

// WarnWF prints message with fields to use multiple key=value pairs.
func WarnWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	defaultLogger.WarnWF(ctx, logCat, status, fields)
}

// Error prints message at error level.
func Error(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	defaultLogger.Error(ctx, logCat, status, v...)
}

// Errorf prints message at error level.
func Errorf(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	defaultLogger.Errorf(ctx, logCat, status, format, v...)
}

// ErrorWF prints message at error level using Fields with multiple key=value pairs.
func ErrorWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	defaultLogger.ErrorWF(ctx, logCat, status, fields)
}

// Fatal prints and calls os.exit(1).
func Fatal(ctx context.Context, logCat LogCat, status StatusCat, v ...interface{}) {
	defaultLogger.Fatal(ctx, logCat, status, v...)
}

// Fatalf prints and calls os.exit(1).
func Fatalf(ctx context.Context, logCat LogCat, status StatusCat, format string, v ...interface{}) {
	defaultLogger.Fatalf(ctx, logCat, status, format, v...)
}

// FatalWF prints and calls os.exit(1) with multiple key=value pairs.
func FatalWF(ctx context.Context, logCat LogCat, status StatusCat, fields *Fields) {
	defaultLogger.FatalWF(ctx, logCat, status, fields)
}
