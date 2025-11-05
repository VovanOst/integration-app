package logger

import (
	"fmt"
	"github.com/gookit/goutil/errorx"
	"github.com/rs/zerolog"
	"os"
)

func init() {
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return fmt.Sprintf("%+v", err)
	}
}

type Logger struct {
	log zerolog.Logger
}

func NewLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stdout).
		With().
		Logger()

	return &Logger{
		log: logger,
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log.Debug().Stack().Msgf(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log.Info().Stack().Msgf(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log.Warn().Stack().Msgf(msg, args...)
}

func (l *Logger) Error(msg string, err error, args ...interface{}) {
	e := errorx.Errorf(err.Error(), args...)
	l.log.Error().Stack().Err(e).Msg(msg)
}
