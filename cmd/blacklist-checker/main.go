package main

import (
	"github.com/ilijamt/blacklist_checker/cmd/blacklist-checker/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var err error
	if err = cmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
