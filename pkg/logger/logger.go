package logger

import (
	"fmt"

	"github.com/rs/zerolog"
)

const (
	_logFieldComponent = "component"
)

type Field func(e *zerolog.Event)

type Logger interface {
	Info(msg string, fields ...Field)
	Infof(msg string, args ...any)
	Error(msg string, fields ...Field)
	Errorf(msg string, args ...any)
	Warn(msg string, fields ...Field)
	Warnf(msg string, args ...any)
	Debug(msg string, fields ...Field)
	Debugf(msg string, args ...any)
	Component(name string) Logger
	Field(key, name string) Logger
}

func Str(key, val string) Field {
	return func(e *zerolog.Event) {
		e.Str(key, val)
	}
}

func Err(err error) Field {
	return func(e *zerolog.Event) {
		e.Err(err)
	}
}

func Int(key string, val int) Field {
	return func(e *zerolog.Event) {
		e.Int(key, val)
	}
}

func Int64(key string, val int64) Field {
	return func(e *zerolog.Event) {
		e.Int64(key, val)
	}
}

func Any(key string, val any) Field {
	return func(e *zerolog.Event) {
		e.Any(key, val)
	}
}

type zeroLoggerWrapper struct {
	logger *zerolog.Logger
}

func NewLogger(logger *zerolog.Logger) Logger {
	return &zeroLoggerWrapper{logger: logger}
}

func (z *zeroLoggerWrapper) Info(msg string, fields ...Field) {
	event := z.logger.Info()
	for _, field := range fields {
		field(event)
	}

	event.Msg(msg)
}

func (z *zeroLoggerWrapper) Infof(msg string, args ...any) {
	z.logger.Info().Msg(fmt.Sprintf(msg, args...))
}

func (z *zeroLoggerWrapper) Error(msg string, fields ...Field) {
	event := z.logger.Error()
	for _, field := range fields {
		field(event)
	}

	event.Msg(msg)
}

func (z *zeroLoggerWrapper) Errorf(msg string, args ...any) {
	z.logger.Error().Msg(fmt.Sprintf(msg, args...))
}

func (z *zeroLoggerWrapper) Warn(msg string, fields ...Field) {
	event := z.logger.Warn()
	for _, field := range fields {
		field(event)
	}

	event.Msg(msg)
}

func (z *zeroLoggerWrapper) Warnf(msg string, args ...any) {
	z.logger.Warn().Msg(fmt.Sprintf(msg, args...))
}

func (z *zeroLoggerWrapper) Debug(msg string, fields ...Field) {
	event := z.logger.Debug()
	for _, field := range fields {
		field(event)
	}

	event.Msg(msg)
}

func (z *zeroLoggerWrapper) Debugf(msg string, args ...any) {
	z.logger.Debug().Msg(fmt.Sprintf(msg, args...))
}

func (z *zeroLoggerWrapper) Component(name string) Logger {
	l := z.logger.With().Str(_logFieldComponent, name).Logger()
	return &zeroLoggerWrapper{
		logger: &l,
	}
}

func (z *zeroLoggerWrapper) Field(field, value string) Logger {
	l := z.logger.With().Str(field, value).Logger()
	return &zeroLoggerWrapper{
		logger: &l,
	}
}

func (z *zeroLoggerWrapper) zeroLoggerWrapper() *zeroLoggerWrapper {
	return z
}
