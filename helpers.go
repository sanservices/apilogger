package apilogger

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// location returns the location of the log call
// as a file/line combination, mostly advantageous
// for dev purposes as this would act as a hyperlink
// to the line of code in question in most IDEs/editors
func location() string {
	workDir, _ := os.Getwd()
	_, fn, line, _ := runtime.Caller(depth)
	fn = strings.TrimPrefix(fn, workDir+"/")
	return fmt.Sprintf("%s:%d", fn, line)
}

// returns the name of caller function.
func funcName() string {
	// Skip 2 levels to get the caller.
	pc, _, _, ok := runtime.Caller(depth)
	if !ok {
		fmt.Println("MSG: NO CALLER")
		return ""
	}

	// get the function caller.
	caller := runtime.FuncForPC(pc)
	if caller == nil {
		fmt.Println("MSG CALLER WAS NIL")
	}

	// remove extra file path characters.
	r := regexp.MustCompile(`[^\/]+$`)
	return fmt.Sprintf("%s", r.FindString(caller.Name()))
}

// returns ip format.
func formatIPAddr(addr string) string {
	r := regexp.MustCompile(`[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}`)
	return fmt.Sprintf("%s", r.FindString(addr))
}

// builds standard information.
func baseMessage(l *Logger, logCat LogCat) string {
	elapsed := time.Since(l.startTime)
	msElapsed := float64(elapsed.Nanoseconds()) / float64(time.Millisecond)

	return fmt.Sprintf(
		`location="%s", requestId="%s", clientIp="%s", apiKey="%s", sessionId="%s", ms="%f", function="%s", code="%s", type="%s"`,
		location(),
		l.requestID,
		l.remoteAddr,
		l.apiKey,
		l.session,
		msElapsed,
		funcName(),
		logCat.Code,
		logCat.Type)
}

// formats and finalizes the log content
func finalMessageWF(l *Logger, logCat LogCat, fields *Fields) string {
	base := baseMessage(l, logCat)
	msg := ""

	for k, v := range *fields {
		msg += fmt.Sprintf(`%s="%v", `, k, v)
	}

	msg = strings.TrimRight(msg, ", ")

	return base + ", " + msg
}

// formats and finalizes the log content
func finalMessage(l *Logger, logCat LogCat, v ...interface{}) string {
	base := baseMessage(l, logCat)
	msg := fmt.Sprint(v...)
	wrappedMsg := fmt.Sprintf(`message="%s"`, msg)

	return base + ", " + wrappedMsg
}

// formats and finalizes the log content
func finalMessagef(l *Logger, logCat LogCat, format string, v ...interface{}) string {
	base := baseMessage(l, logCat)
	msg := fmt.Sprintf(format, v...)
	wrappedMsg := fmt.Sprintf(`message="%s"`, msg)

	return base + ", " + wrappedMsg
}
