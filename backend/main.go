package main

import (
	"flag"
	"log"
	"os"

	"github.com/KairoBoni/sales/server"
)

func main() {
	flag.Parse()

	err, s := server.NewServer(os.Getenv("CONFIG_FILEPATH"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start API Server")
	}

	s.Run()
}
