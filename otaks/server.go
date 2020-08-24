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
	"github.com/tma5/otaks/config"
	"github.com/tma5/otaks/state"
	"github.com/tma5/otaks/web"
)

// Server defines the runtime state of the otaks service
type Server struct {
	Logger  *logrus.Logger
	State   *state.State
	running bool

	appServer  *app.Server
	apiServer  *api.Server
	chatServer *chat.Server
	webServer  *web.Server
}

// NewServer returns a new server instance
func NewServer(config *config.Config) (*Server, error) {
	s := new(Server)
	s.State = state.NewState(config)
	log.Printf("%+v", s.State)
	log.SetOutput(os.Stdout)
	log.Printf("s.State.Config.Server.Logging.Level: %s")
	logLevel, err := log.ParseLevel(s.State.Config.Server.Logging.Level)
	if err != nil {
		log.Error("problem setting loglevel", err)
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
	s.appServer = app.NewServer(s.State)

	log.Tracef("Initializing api server")
	s.apiServer = api.NewServer(s.State)

	log.Tracef("Initializing web server")
	s.webServer = web.NewServer(s.State)

	var g run.Group
	g.Add(func() error {
		log.Tracef("Starting app server")
		return s.appServer.Run()
	}, func(err error) {
		if err != nil {
			log.Error(err)
		}
	})

	g.Add(func() error {
		log.Tracef("Starting api server")
		return s.apiServer.Run()
	}, func(err error) {
		if err != nil {
			log.Error(err)
		}
	})

	g.Add(func() error {
		log.Trace("Starting web server")
		return s.webServer.Run()
	}, func(err error) {
		if err != nil {
			log.Error(err)
		}
	})

	return g.Run()
}
