package logger

import (
	"context"

	"github.com/mochammadshenna/saza-aluminium/state"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

var lf loggerField

func Init() {
	lf = newLoggerField()

	// logPath := "util/logger/logging.log"


	// Logger.SetOutput(mw)
}

type loggerField struct {
	// Custom field
	RequestId string `json:"requestId"`

	// Field handle by logger
	Message        string `json:"message"`
	Severity       string `json:"severity"`
	Timestamp      string `json:"timestamp"`
	SourceLocation string `json:"sourceLocation"`
}

func newLoggerField() loggerField {
	return loggerField{
		RequestId:      "requestId",
		Message:        "message",
		Severity:       "severity",
		Timestamp:      "timestamp",
		SourceLocation: "sourceLocation",
	}
}

func LoggerField() loggerField {
	return lf
}

func Info(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Info(args...)
		return
	}
	Logger.Info(args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Infof(format, args...)
		return
	}
	Logger.Infof(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Warn(args...)
		return
	}
	Logger.Warn(args...)
}

func Error(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Error(args...)
		return
	}
	Logger.Error(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Errorf(format, args...)
		return
	}
	Logger.Errorf(format, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Panic(args...)
		return
	}
	Logger.Panic(args...)
}

func withFields(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{}

	requestId := ctx.Value(state.HttpHeaders().RequestId)
	if requestId != "" {
		fields[LoggerField().RequestId] = requestId
	}

	return fields
}
