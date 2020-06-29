package main

import (
	"flag"

	"github.com/KairoBoni/sales/backend/server"
	"github.com/rs/zerolog/log"
)

func main() {
	flag.Parse()

	s, err := server.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create API Server")
	}

	if err = s.Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to start API Server")
	}
}
