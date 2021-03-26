package apilogger

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	apiKey := "r12d3f4"
	requestID := "1234"
	ip := "127.0.0.1"
	session := "b011157f-a97b-4090-8d56-6d4bafb4f60c"
	path := "/test"

	rq, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Error(err)
	}

	rq = rq.WithContext(context.WithValue(
		rq.Context(), APIKEY, apiKey))
	rq = rq.WithContext(context.WithValue(
		rq.Context(), RequestIDKey, requestID))
	rq = rq.WithContext(context.WithValue(
		rq.Context(), RemoteAddrKey, ip))
	rq = rq.WithContext(context.WithValue(
		rq.Context(), SessionIDKey, session))

	logger := New(rq.Context())

	assertEquals(t, logger.apiKey, apiKey)
	assertEquals(t, logger.remoteAddr, ip)
	assertEquals(t, logger.requestID, requestID)
	assertEquals(t, logger.session, session)
}

func TestFuncName(t *testing.T) {
	expected := "apilogger.TestFuncName"

	// mimics call stack depth
	func1 := func() string { return funcName() }
	func2 := func() string { return func1() }
	func3 := func() string { return func2() }
	func4 := func() string { return func3() }

	output := func4()

	assertEquals(t, output, expected)
}

func TestFormatIPAddr(t *testing.T) {
	expected := "127.0.0.1"
	output := formatIPAddr("127.0.0.1")

	assertEquals(t, output, expected)
}

func TestBaseMessage(t *testing.T) {
	rq, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Error(err)
	}

	logger := New(rq.Context())
	logCat := LogCatStartUp

	// mimics call stack depth
	func1 := func() string { return baseMessage(logger, logCat, time.Now()) }
	func2 := func() string { return func1() }
	func3 := func() string { return func2() }
	func4 := func() string { return func3() }

	output := func4()

	if !strings.Contains(output, "requestId") {
		t.Errorf("Output insufficient - [%s]", output)
	}
}

func TestFinalMessage(t *testing.T) {
	rq, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Error(err)
	}

	logger := New(rq.Context())
	logCat := LogCatStartUp
	output := finalMessage(logger, logCat, time.Now(), "hello test")

	assertStrContains(t, output, "hello test")
	assertStrContains(t, output, " code=\""+logCat.Code+"\"")
	assertStrContains(t, output, " type=\""+logCat.Type+"\"")
}

func TestFinalMessagef(t *testing.T) {
	rq, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Error(err)
	}

	logger := New(rq.Context())
	logCat := LogCatStartUp
	output := finalMessagef(logger, logCat, time.Now(), "%s", "hello test")

	assertStrContains(t, output, " message=\"hello test\"")
	assertStrContains(t, output, " code=\""+logCat.Code+"\"")
	assertStrContains(t, output, " type=\""+logCat.Type+"\"")
}

func TestFinalMessageWF(t *testing.T) {
	rq, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Error(err)
	}

	logger := New(rq.Context())
	logCat := LogCatStartUp
	output := finalMessageWF(logger, logCat, time.Now(), &Fields{"message": "hello test"})

	assertStrContains(t, output, " message=\"hello test\"")
	assertStrContains(t, output, " code=\""+logCat.Code+"\"")
	assertStrContains(t, output, " type=\""+logCat.Type+"\"")
}

func assertEquals(
	t *testing.T, output interface{}, expected interface{}) {

	if output != expected {
		t.Errorf("Output [%v] not equal to expected [%v]",
			output, expected)
	}
}

func assertStrContains(t *testing.T, output string, expected string) {
	if !strings.Contains(output, expected) {
		t.Errorf("Output insufficient - [%s]", output)
	}
}
