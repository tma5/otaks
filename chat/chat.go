package chat

import "github.com/tma5/otaks/state"

const (
	// MulticastChatAddress provides the multicast chat endpoint
	MulticastChatAddress = "224.10.10.1:17012"
)

// Server contains the chat server state
type Server struct {
	state *state.State
}

// NewServer returns a new instance of chat server
func NewServer(state *state.State) *Server {
	chat := Server{
		state: state,
	}
	return &chat
}

func (srv *Server) init() {

}

// Run begins the chat server
func (srv *Server) Run() error {
	return nil
}
