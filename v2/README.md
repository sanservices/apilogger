# apilogger v2

Simple api logger for golang with the purpose of log basic information like api-key, clinet ip, request-id.

# Installation

Run `go get github.com/san-services/apilogger/v2`

# Changes

1. Added context as first argument to log functions. (thus we can pass some additional values to log, as now is implemented for execution time tracking)
2. Made logger global

# Usage

```go
package main

import (
	"context"
	"errors"
	"time"

	"github.com/sanservices/apilogger/v2"
)

func main() {
	ctx := context.TODO()

	ctx = context.WithValue(ctx, apilogger.APIKEY, "apikey1")
	ctx = context.WithValue(ctx, apilogger.RequestIDKey, "requestIdKey1")
	ctx = context.WithValue(ctx, apilogger.RemoteAddrKey, "remoteaddrKey1")
	ctx = context.WithValue(ctx, apilogger.SessionIDKey, "sessionIdKey1")
	ctx = context.WithValue(ctx, apilogger.StartTime, time.Now())

	l := apilogger.New(ctx)

	l.Info(ctx, apilogger.LogCatDebug, "This is an info message")
	l.Warn(ctx, apilogger.LogCatDebug, "This is a warning message")
	l.Error(ctx, apilogger.LogCatDebug, "This is an error message")

	l.Infof(ctx, apilogger.LogCatDebug, "This is an info message with %s", "some variable")
	l.Warnf(ctx, apilogger.LogCatDebug, "This is a warning message with format, some number: %d", 100)
	l.Errorf(ctx, apilogger.LogCatDebug, "This is an error message with format, %v", errors.New("an error"))

	l.InfoWF(ctx, apilogger.LogCatDebug, &apilogger.Fields{"message": "my message"})
	l.WarnWF(ctx, apilogger.LogCatDebug, &apilogger.Fields{"warning": "my warning", "other": "another message"})
	l.ErrorWF(ctx, apilogger.LogCatDebug, &apilogger.Fields{"error": errors.New("my error message")})

}
```

output is

```shell
INFO 2021/03/23 23:37:40 location="main.go:21", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.062139", function="main.main", code="DBG001", type="debug", message="This is an info message"
WARN 2021/03/23 23:37:40 location="main.go:22", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.150727", function="main.main", code="DBG001", type="debug", message="This is a warning message"
ERROR 2021/03/23 23:37:40 location="main.go:23", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.174835", function="main.main", code="DBG001", type="debug", message="This is an error message"
INFO 2021/03/23 23:37:40 location="main.go:25", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.190661", function="main.main", code="DBG001", type="debug", message="This is an info message with some variable"
WARN 2021/03/23 23:37:40 location="main.go:26", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.204906", function="main.main", code="DBG001", type="debug", message="This is a warning message with format, some number: 100"
ERROR 2021/03/23 23:37:40 location="main.go:27", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.219298", function="main.main", code="DBG001", type="debug", message="This is an error message with format, An error"
INFO 2021/03/23 23:37:40 location="main.go:29", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.252828", function="main.main", code="DBG001", type="debug", message="my message"
WARN 2021/03/23 23:37:40 location="main.go:30", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.274832", function="main.main", code="DBG001", type="debug", other="another message", warning="my warning"
ERROR 2021/03/23 23:37:40 location="main.go:31", requestId="requestIdKey1", clientIp="", apiKey="apikey1", sessionId="sessionIdKey1", ms="0.287675", function="main.main", code="DBG001", type="debug", error="my error message"
```

or with global logger initialization

```go 
package main

import (
	"context"
	"errors"
	"time"

	"github.com/sanservices/apilogger/v2"
)

func main() {
    ctx := context.TODO()

    ctx = context.WithValue(ctx, apilogger.APIKEY, "apikey1")
    ctx = context.WithValue(ctx, apilogger.RequestIDKey, "requestIdKey1")
    ctx = context.WithValue(ctx, apilogger.RemoteAddrKey, "remoteaddrKey1")
    ctx = context.WithValue(ctx, apilogger.SessionIDKey, "sessionIdKey1")
    ctx = context.WithValue(ctx, apilogger.StartTime, time.Now())

    apilogger.New(ctx) // actually initializes global logger

    apilogger.Info(ctx, apilogger.LogCatDebug, "This is an info message")
    apilogger.Warn(ctx, apilogger.LogCatDebug, "This is a warning message")
    apilogger.Error(ctx, apilogger.LogCatDebug, "This is an error message")

    apilogger.Infof(ctx, apilogger.LogCatDebug, "This is an info message with %s", "some variable")
    apilogger.Warnf(ctx, apilogger.LogCatDebug, "This is a warning message with format, some number: %d", 100)
    apilogger.Errorf(ctx, apilogger.LogCatDebug, "This is an error message with format, %v", errors.New("an error"))

    apilogger.InfoWF(ctx, apilogger.LogCatDebug, &apilogger.Fields{"message": "my message"})
    apilogger.WarnWF(ctx, apilogger.LogCatDebug, &apilogger.Fields{"warning": "my warning", "other": "another message"})
    apilogger.ErrorWF(ctx, apilogger.LogCatDebug, &apilogger.Fields{"error": errors.New("my error message")})
}
```

output is the same as above
