package otaks

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/api"
	"github.com/tma5/otaks/app"
	"github.com/tma5/otaks/chat"
)

// Server defines the runtime state of the otaks service
type Server struct {
	Logger  *logrus.Logger
	Config  *Config
	running bool

	appServer  *app.Server
	apiServer  *api.Server
	chatServer *chat.Server
}

func NewServer(config *Config) (*Server, error) {
	s := new(Server)
	s.Config = config

	log.SetOutput(os.Stdout)
	logLevel, err := log.ParseLevel(s.Config.Server.Logging.Level)
	if err != nil {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logLevel)
	}

	s.running = true

	return s, nil
}

// Shutdown gracefully shuts down the otaks server
func (s *Server) Shutdown() {
	s.running = false
}

func signalHandler() {
	sigc := make(chan os.Signal)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGHUP,
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
			case syscall.SIGTERM:
				log.Infof("SIGTERM received, stopping...")
				exit <- 0
			default:
				log.Errorf("Signal received, unknown type")
			}
		}
	}()

	code := <-exit
	os.Exit(code)
}

// Run starts the otaks server
func (s *Server) Run() error {
	log.Tracef("Initializing app server")
	s.appServer = app.NewServer()

	log.Tracef("Initializing api server")
	s.apiServer = api.NewServer()

	log.Tracef("Initializing chat server")
	s.chatServer = chat.NewServer()

	var g run.Group
	g.Add(func() error {
		log.Tracef("Starting app server")
		return s.appServer.Run()
	}, func(err error) {
		log.Error(err)
	})

	g.Add(func() error {
		log.Tracef("Starting api server")
		return s.apiServer.Run()
	}, func(err error) {
		log.Error(err)
	})

	g.Add(func() error {
		log.Tracef("Starting chat server")
		return s.chatServer.Run()
	}, func(err error) {
		log.Error(err)
	})

	return g.Run()
}
