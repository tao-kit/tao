package zerologx

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/sllt/tao/core/logx"
)

type (
	ZeroLogWriter struct {
		logger zerolog.Logger
	}
)

func NewZeroLogWriter(logger zerolog.Logger) logx.Writer {
	return &ZeroLogWriter{logger: logger}
}

func (w *ZeroLogWriter) Alert(v any) {
	w.logger.Error().Msg(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Close() error {
	w.logger.Fatal().Msg("")
	return nil
}

func (w *ZeroLogWriter) Debug(v any, fields ...logx.LogField) {
	toZeroLogInterface(w.logger.Debug(), fields...).Msgf(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Error(v any, fields ...logx.LogField) {
	toZeroLogInterface(w.logger.Error(), fields...).Msgf(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Info(v any, fields ...logx.LogField) {
	toZeroLogInterface(w.logger.Info(), fields...).Msgf(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Severe(v any) {
	w.logger.Fatal().Msg(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Slow(v any, fields ...logx.LogField) {
	toZeroLogInterface(w.logger.Warn(), fields...).Msgf(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Stack(v any) {
	w.logger.Error().Msgf(fmt.Sprint(v))
}

func (w *ZeroLogWriter) Stat(v any, fields ...logx.LogField) {
	toZeroLogInterface(w.logger.Info(), fields...).Msgf(fmt.Sprint(v))
}

func toZeroLogInterface(event *zerolog.Event, fields ...logx.LogField) *zerolog.Event {
	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}
	return event
}
