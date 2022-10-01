package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

type logEntryContextKey struct{}

var (
	logEntryKey = logEntryContextKey{}
)

// FromContext Gets the logger from context
func FromContext(ctx context.Context) logrus.FieldLogger {
	if ctx == nil {
		panic("nil context")
	}
	v, ok := ctx.Value(logEntryKey).(logrus.FieldLogger)
	if !ok {
		return logrus.WithContext(ctx)
	}
	return v
}
