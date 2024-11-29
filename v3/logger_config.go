package apilogger

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func SetLoggerConfig(ctx context.Context, taskName string) context.Context {
	var keyCtx ContextKey = "context-data"
	contextData := CtxKeys{
		TaskName:  taskName,
		UUID:      uuid.New().String(),
		StartTime: time.Now(),
	}
	return context.WithValue(ctx, keyCtx, contextData)
}
