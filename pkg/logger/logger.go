package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

type Config interface {
}

const loggerTimeFormat = "2006-01-02T15:04:05.000000Z07:00"

func Init(cfg Config) {
	zerolog.TimeFieldFormat = loggerTimeFormat
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.NoColor = false
		w.TimeFormat = loggerTimeFormat
	})

	log.Logger = zerolog.New(logger).Level(zerolog.DebugLevel).
		With().Stack().Timestamp().Logger()
}
