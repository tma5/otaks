package otaks

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// Server defines the runtime state of the otaks service
type Server struct {
	mutex sync.RWMutex

	started bool

	Logger    *logrus.Logger
	Config    *Config
	
	cotServer  *CotServer
}

func NewServer(config *Config) (*Server, error) {
	s := new(Server)
	s.Config = config
	s.bootstrap()

	return s, nil
}

// InitServer() initializes a otaks service definition
func (s *Server) bootstrap() {
	log.SetOutput(os.Stdout)
	logLevel, err := log.ParseLevel(s.Config.Server.Logging.Level)
	if err != nil {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logLevel)
	}

	host := s.Config.Server.Host
	port := s.Config.Server.Port

	log.Tracef("Initializing COT Server")
	s.cotServer = &CotServer{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Otaks: s,
	}
	log.Infof("Initialized COT Server")
}

// Shutdown gracefully shuts down the otaks server
func (s *Server) Shutdown() {
	s.cotServer.Shutdown()
}

// ListenAndServe starts the otaks server
func (s *Server) ListenAndServe() error {
	go s.cotServer.ListenAndServe()

	sigc := make(chan os.Signal)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		//syscall.SIGINFO, // TODO: SIGINFO doesn't exist in *nix, just BSD?
		syscall.SIGTERM,
	)

	exit := make(chan int)
	go func() {
		for {
			sigv := <-sigc
			switch sigv {
			case syscall.SIGINT:
				log.Infof("SIGINT received, stopping...")
				exit <- 0
			case syscall.SIGHUP:
				// TODO: Implement SIGHUP
				log.Infof("SIGHUP received, reloading...")
				log.Warnf("Not yet implemented!")
			case syscall.SIGUSR1:
				// TODO: Implement SIGUSR1
				log.Infof("SIGUSR1 received, reloading...")
				log.Warnf("Not yet implemented!")
			case syscall.SIGUSR2:
				// TODO: Implement SIGUSR2
				log.Infof("SIGUSR1 received, reloading...")
				log.Warnf("Not yet implemented!")
			// case syscall.SIGINFO:
			// 	s.printInfo()
			case syscall.SIGTERM:
				log.Infof("SIGTERM received, stopping...")
				s.Shutdown()
				exit <- 0
			default:
				log.Errorf("Signal received, unknown type")
			}
		}
	}()

	code := <-exit
	os.Exit(code)

	return nil
}