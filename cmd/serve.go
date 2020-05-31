package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tma5/otaks/config"
	"github.com/tma5/otaks/otaks"
)

var (
	serveCmd = &cobra.Command{
		Use:     "serve",
		Short:   "otaks serve",
		Aliases: []string{"s"},
		Run:     serve,
	}
)

func init() {

}

func serve(cmd *cobra.Command, args []string) {
	config, err := config.NewConfig(configLocation)
	if err != nil {
		log.Fatal(err)
	}

	if host != defaultHost {
		config.Server.Host = host
	}

	if port != defaultPort {
		config.Server.App.Port = port
	}

	if logLevel != defaultLogLevel {
		config.Server.Logging.Level = logLevel
	}

	server, err := otaks.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal("serve:", server.Run())
}
