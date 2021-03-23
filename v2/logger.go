package apilogger

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

const (
	// APIKEY is the context key used to
	// access the api-key request header value
	APIKEY string = "api-key"

	// RequestIDKey is the context key used to
	// access the x-request-id request header value
	RequestIDKey string = "x-request-id"

	// RemoteAddrKey is the context key used to
	// access the remote-address request header value
	RemoteAddrKey string = "remote-address"

	// SessionIDKey is the context key used to
	// access the session request header value
	SessionIDKey string = "session"

	// StartTime is the context key used to
	// access start time of transaction
	StartTime string = "start-time"

	// Depth of the callstack - needed to determine
	// the initial caller function, for example
	depth int = 5

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
	requestID  string
	apiKey     string
	remoteAddr string
	session    string
}

type Fields map[string]interface{}

var defaultLogger *Logger

// New returns a new Logger instance.
func New(ctx context.Context, outPath string) *Logger {
	var output io.Writer = os.Stdout
	var errOutput io.Writer = os.Stderr

	if len(outPath) != 0 {
		file, err := os.OpenFile(
			outPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			log.Println("Failed to open log file", err)
			return nil
		}

		output = io.MultiWriter(file, output)
		errOutput = io.MultiWriter(file, errOutput)
	}

	var id, apiKey, addr, session string

	id, _ = ctx.Value(RequestIDKey).(string)
	apiKey, _ = ctx.Value(APIKEY).(string)
	addr, _ = ctx.Value(RemoteAddrKey).(string)
	session, _ = ctx.Value(SessionIDKey).(string)

	defaultLogger = &Logger{
		requestID:  id,
		apiKey:     apiKey,
		remoteAddr: formatIPAddr(addr),
		session:    session,
		output:     output,
		errOutput:  errOutput,
	}

	return defaultLogger
}

// prints message.
func (l *Logger) printlnWF(logger *log.Logger, logCat LogCat, startTime time.Time, fields *Fields) {
	logger.Println(finalMessageWF(l, logCat, startTime, fields))
}

func (l *Logger) println(logger *log.Logger, logCat LogCat, startTime time.Time, v ...interface{}) {
	logger.Println(finalMessage(l, logCat, startTime, v...))
}

func (l *Logger) printlnf(logger *log.Logger, logCat LogCat, startTime time.Time, format string, v ...interface{}) {
	logger.Println(finalMessagef(l, logCat, startTime, format, v...))
}

func (l *Logger) Info(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.println(l.infoLog, logCat, startTime, v...)
}

func (l *Logger) Infof(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.printlnf(l.infoLog, logCat, startTime, format, v...)
}

func (l *Logger) InfoWF(ctx context.Context, logCat LogCat, fields *Fields) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.printlnWF(l.infoLog, logCat, startTime, fields)
}

func (l *Logger) Warn(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.println(l.warningLog, logCat, startTime, v...)
}

func (l *Logger) Warnf(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.printlnf(l.warningLog, logCat, startTime, format, v...)
}

func (l *Logger) WarnWF(ctx context.Context, logCat LogCat, fields *Fields) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.printlnWF(l.warningLog, logCat, startTime, fields)
}

func (l *Logger) Error(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.println(l.errorLog, logCat, startTime, v...)
}

func (l *Logger) Errorf(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.printlnf(l.errorLog, logCat, startTime, format, v...)
}

func (l *Logger) ErrorWF(ctx context.Context, logCat LogCat, fields *Fields) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.printlnWF(l.errorLog, logCat, startTime, fields)
}

func (l *Logger) Fatal(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessage(l, logCat, startTime, v...))
}

func (l *Logger) Fatalf(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessagef(l, logCat, startTime, format, v...))
}

func (l *Logger) FatalWF(ctx context.Context, logCat LogCat, fields *Fields) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	// Get start time from context
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessageWF(l, logCat, startTime, fields))
}

// Info prints message with logging level of info
func Info(ctx context.Context, logCat LogCat, v ...interface{}) {
	defaultLogger.Info(ctx, logCat, v...)
}

// Infof prints a message using the specified format.
func Infof(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	defaultLogger.Infof(ctx, logCat, format, v...)
}

// InfoWF prints message using Fields struct to pass multiple key=value pairs.
func InfoWF(ctx context.Context, logCat LogCat, fields *Fields) {
	defaultLogger.InfoWF(ctx, logCat, fields)
}

// Warn prints message with logging level of info
func Warn(ctx context.Context, logCat LogCat, v ...interface{}) {
	defaultLogger.Warn(ctx, logCat, v...)
}

// Warnf prints a message using the specified format.
func Warnf(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	defaultLogger.Warnf(ctx, logCat, format, v...)
}

// WarnWF prints message with fields to use multiple key=value pairs.
func WarnWF(ctx context.Context, logCat LogCat, fields *Fields) {
	defaultLogger.WarnWF(ctx, logCat, fields)
}

// Error prints message at error level.
func Error(ctx context.Context, logCat LogCat, v ...interface{}) {
	defaultLogger.Error(ctx, logCat, v...)
}

// Errorf prints message at error level.
func Errorf(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	defaultLogger.Errorf(ctx, logCat, format, v...)
}

// ErrorWF prints message at error level using Fields with multiple key=value pairs.
func ErrorWF(ctx context.Context, logCat LogCat, fields *Fields) {
	defaultLogger.ErrorWF(ctx, logCat, fields)
}

// Fatal prints and calls os.exit(1).
func Fatal(ctx context.Context, logCat LogCat, v ...interface{}) {
	defaultLogger.Fatal(ctx, logCat, v...)
}

// Fatalf prints and calls os.exit(1).
func Fatalf(ctx context.Context, logCat LogCat, format string, v ...interface{}) {
	defaultLogger.Fatalf(ctx, logCat, format, v...)
}

// FatalWF prints and calls os.exit(1) with multiple key=value pairs.
func FatalWF(ctx context.Context, logCat LogCat, fields *Fields) {
	defaultLogger.FatalWF(ctx, logCat, fields)
}
