# apilogger v3

Simple api logger for golang scheduled task with the purpose of log basic information like task name, uuid and start time.

# Installation

Run `go get github.com/sanservices/apilogger/v3`

# Changes

1. Apilogger modified to log shceduled tasks information , warnings and errors.

# Usage

```go
package main

import (
	"context"
	"errors"
	"time"
	"github.com/sanservices/apilogger/v3"
)

func main() {
	ctx := context.Background()

	var keyCtx apilogger.ContextKey = "key-name"

	contextData := apilogger.CtxKeys{
		TaskName:  "task-name",
		UUID:      "uuid",
		StartTime: time.Now(),
	}
	
	ctx = context.WithValue(ctx, keyCtx, contextData)

	l := apilogger.New(ctx)

  	l.Info(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an info message")
	l.Warn(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is a warning message")
	l.Error(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an error message")

	l.Infof(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an info message with %s", "some variable")
	l.Warnf(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is a warning message with format, some number: %d", 100)
	l.Errorf(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an error message with format, %v", errors.New("an error"))

	l.InfoWF(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, &apilogger.Fields{"message": "my message"})
	l.WarnWF(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, &apilogger.Fields{"warning": "my warning", "other": "another message"})
	l.ErrorWF(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, &apilogger.Fields{"error": errors.New("my error message")})

}
```

output is

```shell
INFO 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="1888.224446",  function="main.main", code="DBG001", type="debug", status="Debug", message="This is an info message"
WARN 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="60615.903550",  function="main.main", code="DBG001", type="debug", status="Debug", message="This is a warning message"
ERROR 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="1679.401089",  function="main.main", code="DBG001", type="debug", status="Debug", message="This is an error message"

INFO 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="1237.328450",  function="main.main", code="DBG001", type="debug", status="Debug", message="This is an info message with some variable"
WARN 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="18724.969123",  function="main.main", code="DBG001", type="debug", status="Debug", message="This is a warning message with format, some number: 100"
ERROR 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="35366.390624",  function="main.main", code="DBG001", type="debug", status="Debug", message="This is an error message with format, an error"

INFO 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="5253.215990",  function="main.main", code="DBG001", type="debug", status="Debug", message="my message"
WARN 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="33911.525828",  function="main.main", code="DBG001", type="debug", status="Debug", warning="my warning", other="another message"
ERROR 2024/09/23 11:29:55 uuid="20d989f8", taskName="Task-Name", location="main.go:19", ms="45049.412831",  function="main.main", code="DBG001", type="debug", status="Debug", error="my error message"
```

or with global logger initialization

```go 
package main

import (
	"context"
	"errors"
	"time"
	"github.com/sanservices/apilogger/v3"
)

func main() {
	ctx := context.Background()

	var keyCtx apilogger.ContextKey = "key-name"
	
	contextData := apilogger.CtxKeys{
		TaskName:  "task-name",
		UUID:      "uuid",
		StartTime: time.Now(),
	}
	
	ctx = context.WithValue(ctx, keyCtx, contextData)

	apilogger.New(ctx)// actually initializes global logger

  	apilogger.Info(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an info message")
	apilogger.Warn(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is a warning message")
	apilogger.Error(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an error message")

	apilogger.Infof(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an info message with %s", "some variable")
	apilogger.Warnf(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is a warning message with format, some number: %d", 100)
	apilogger.Errorf(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, "This is an error message with format, %v", errors.New("an error"))

	apilogger.InfoWF(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, &apilogger.Fields{"message": "my message"})
	apilogger.WarnWF(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, &apilogger.Fields{"warning": "my warning", "other": "another message"})
	apilogger.ErrorWF(ctx, apilogger.LogCatDebug, apilogger.StatusCatDebug, &apilogger.Fields{"error": errors.New("my error message")})

}
```

output is the same as above
