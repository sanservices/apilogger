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
	startTime  time.Time
	requestID  string
	apiKey     string
	remoteAddr string
	session    string
}

type Fields map[string]interface{}

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
	if ctx != nil {
		id, _ = ctx.Value(RequestIDKey).(string)
		apiKey, _ = ctx.Value(APIKEY).(string)
		addr, _ = ctx.Value(RemoteAddrKey).(string)
		session, _ = ctx.Value(SessionIDKey).(string)
	}

	return &Logger{
		startTime:  time.Now(),
		requestID:  id,
		apiKey:     apiKey,
		remoteAddr: formatIPAddr(addr),
		session:    session,
		output:     output,
		errOutput:  errOutput,
	}
}

// prints message.
func (l *Logger) printlnWF(
	logger *log.Logger, logCat LogCat, fields *Fields) {

	logger.Println(finalMessageWF(l, logCat, fields))
}

func (l *Logger) println(logger *log.Logger, logCat LogCat, v ...interface{}) {
	logger.Println(finalMessage(l, logCat, v...))
}

func (l *Logger) printlnf(logger *log.Logger, logCat LogCat, format string, v ...interface{}) {
	logger.Println(finalMessagef(l, logCat, format, v...))
}

func (l *Logger) Info(logCat LogCat, v ...interface{}) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	l.println(l.infoLog, logCat, v...)
}

// Infof prints a message using the specified format.
func (l *Logger) Infof(logCat LogCat, format string, v ...interface{}) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	l.printlnf(l.infoLog, logCat, format, v...)
}

// InfoWF prints message using Fields struct to pass multiple key=value pairs.
func (l *Logger) InfoWF(logCat LogCat, fields *Fields) {
	if l.infoLog == nil {
		l.infoLog = log.New(l.output, prefixInfo, log.Ldate|log.Ltime)
	}

	l.printlnWF(l.infoLog, logCat, fields)
}

func (l *Logger) Warn(logCat LogCat, v ...interface{}) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	l.println(l.warningLog, logCat, v...)
}

func (l *Logger) Warnf(logCat LogCat, format string, v ...interface{}) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}

	l.printlnf(l.warningLog, logCat, format, v...)
}

// WarnWF prints message with fields to use multiple key=value pairs.
func (l *Logger) WarnWF(logCat LogCat, fields *Fields) {
	if l.warningLog == nil {
		l.warningLog = log.New(l.output, prefixWarn, log.Ldate|log.Ltime)
	}
	l.printlnWF(l.warningLog, logCat, fields)
}

// Error prints message at error level.
func (l *Logger) Error(logCat LogCat, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	l.println(l.errorLog, logCat, v...)
}

// Errorf prints message at error level.
func (l *Logger) Errorf(logCat LogCat, format string, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	l.printlnf(l.errorLog, logCat, format, v...)
}

// ErrorWF prints message at error level using Fields with multiple key=value pairs.
func (l *Logger) ErrorWF(logCat LogCat, fields *Fields) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, prefixError, log.Ldate|log.Ltime)
	}

	l.printlnWF(l.errorLog, logCat, fields)
}

// Fatal prints and calls os.exit(1).
func (l *Logger) Fatal(logCat LogCat, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessage(l, logCat, v...))
}

// Fatalf prints and calls os.exit(1).
func (l *Logger) Fatalf(logCat LogCat, format string, v ...interface{}) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessagef(l, logCat, format, v...))
}

// FatalWF prints and calls os.exit(1) with multiple key=value pairs.
func (l *Logger) FatalWF(logCat LogCat, fields *Fields) {
	if l.errorLog == nil {
		l.errorLog = log.New(l.errOutput, "", log.Ldate|log.Ltime)
	}

	l.errorLog.SetPrefix(prefixFatal)
	l.errorLog.Fatal(finalMessageWF(l, logCat, fields))
}
