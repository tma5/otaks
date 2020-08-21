package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tma5/otaks/config"
)

var (
	host           string
	port           int
	configLocation string
	logLevel       string

	helpFlag bool
)

const (
	defaultConfigLocation = "/etc/otaks/otaks.toml"
	defaultHost           = "0.0.0.0"
	defaultPort           = 8087
	defaultLogLevel       = "INFO"
)

var RootCmd = &cobra.Command{
	Use:   "otaks",
	Short: "otaks",
	Long:  ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func init() {
	// reassign help to free up `-h` for `--host`
	RootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help for otaks")

	// set version on root cmd to cheat cobra version into play
	RootCmd.Version = config.Version

	RootCmd.PersistentFlags().StringVarP(&configLocation, "config", "c", defaultConfigLocation, "configuration file")
	RootCmd.PersistentFlags().StringVarP(&host, "host", "h", defaultHost, "host to bind")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", defaultPort, "port")
	RootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", defaultLogLevel, "log level [ERROR|WARN|INFO|DEBUG|TRACE]")

	RootCmd.AddCommand(serveCmd)
}
