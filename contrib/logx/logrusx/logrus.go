package logrusx

import (
	"github.com/sirupsen/logrus"
	"github.com/sllt/tao/core/logx"
)

type LogrusWriter struct {
	logger *logrus.Logger
}

func NewLogrusWriter(opts ...func(logger *logrus.Logger)) logx.Writer {
	logger := logrus.New()
	for _, opt := range opts {
		opt(logger)
	}

	return &LogrusWriter{
		logger: logger,
	}
}

func (w *LogrusWriter) Alert(v any) {
	w.logger.Error(v)
}

func (w *LogrusWriter) Close() error {
	w.logger.Exit(0)
	return nil
}

func (w *LogrusWriter) Debug(v any, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Debug(v)
}

func (w *LogrusWriter) Error(v any, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Error(v)
}

func (w *LogrusWriter) Info(v any, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Info(v)
}

func (w *LogrusWriter) Severe(v any) {
	w.logger.Fatal(v)
}

func (w *LogrusWriter) Slow(v any, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Warn(v)
}

func (w *LogrusWriter) Stack(v any) {
	w.logger.Error(v)
}

func (w *LogrusWriter) Stat(v any, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Info(v)
}

func toLogrusFields(fields ...logx.LogField) logrus.Fields {
	logrusFields := make(logrus.Fields)
	for _, field := range fields {
		logrusFields[field.Key] = field.Value
	}
	return logrusFields
}
