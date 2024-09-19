package apilogger

import (
	"os"
	"strings"
	"testing"
	"time"

	assertion "github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	logger := New()

	assert := assertion.New(t)

	assert.Equal(
		&Logger{
			output:    os.Stdout,
			errOutput: os.Stderr,
		}, logger,
	)
}

func TestFuncName(t *testing.T) {
	expected := "v2.TestFuncName"

	// mimics call stack depth
	func1 := func() string { return funcName() }
	func2 := func() string { return func1() }
	func3 := func() string { return func2() }
	func4 := func() string { return func3() }

	output := func4()

	assertion.New(t).Equal(output, expected)
}

func TestBaseMessage(t *testing.T) {
	// mimics call stack depth
	func1 := func() string {
		return baseMessage(LogCatDebug, time.Now(), "UpdatePassword", "12345zw", StatusCatPending)
	}
	func2 := func() string { return func1() }
	func3 := func() string { return func2() }
	func4 := func() string { return func3() }
	func5 := func() string { return func4() }

	output := func5()

	if !strings.Contains(output, "taskName") {
		t.Errorf("Output insufficient - [%s]", output)
	}
}

func TestFinalMessage(t *testing.T) {
	logCat := LogCatStartUp
	output := finalMessage(logCat, time.Now(), "UpdatePassword", "12345zw", StatusCatPending, "hello test")
	assert := assertion.New(t)

	assert.Contains(output, "hello test")
	assert.Contains(output, " code=\""+logCat.Code+"\"")
	assert.Contains(output, " type=\""+logCat.Type+"\"")
}

func TestFinalMessagef(t *testing.T) {
	logCat := LogCatStartUp
	output := finalMessagef(logCat, time.Now(), "UpdatePassword", "12345zw", StatusCatPending, "%s", "hello test")
	assert := assertion.New(t)

	assert.Contains(output, " message=\"hello test\"")
	assert.Contains(output, " code=\""+logCat.Code+"\"")
	assert.Contains(output, " type=\""+logCat.Type+"\"")
}

func TestFinalMessageWF(t *testing.T) {
	logCat := LogCatStartUp
	output := finalMessageWF(logCat, time.Now(), "UpdatePassword", "12345zw", StatusCatPending, &Fields{"message": "hello test"})
	assert := assertion.New(t)

	assert.Contains(output, " message=\"hello test\"")
	assert.Contains(output, " code=\""+logCat.Code+"\"")
	assert.Contains(output, " type=\""+logCat.Type+"\"")
}
