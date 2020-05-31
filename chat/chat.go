package chat

const (
	// MulticastChatAddress provides the multicast chat endpoint
	MulticastChatAddress = "224.10.10.1:17012"
)

// Server contains the chat server state
type Server struct {
}

// NewServer returns a new instance of chat server
func NewServer() *Server {
	chat := Server{}
	return &chat
}

func (srv *Server) init() {

}

// Run begins the chat server
func (srv *Server) Run() error {
	return nil
}
