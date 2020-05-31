package chat

import "net"

const (
	MULTICAST_CHAT_ADDR = "224.10.10.1:17012"
)

type ChatServer struct {
}

func NewChatServer() *ChatServer {
	srv := ChatServer{}
	srv.init()
	return &srv
}

func (srv *ChatServer) init() {

}

func (srv *ChatServer) Run() error {
	return nil
}

func (srv *ChatServer) ListenAndServe() error {
	ln, err := net.Listen("udp", MULTICAST_CHAT_ADDR)
	defer ln.Close()
	if err != nil {
		return err
	}

	return nil
}
