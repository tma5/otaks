package chat

import "net"

const (
	MULTICAST_CHAT_ADDR = "224.10.10.1:17012"
)

type ChatServer struct {
}

func NewChatServer() ChatServer {

}

func (srv *ChatServer) init() {

}

func (srv *ChatServer) Run() error {

}

func (srv *ChatServer) ListenAndServe() error {
	log.Tracef("Starting chat server on %s", MULTICAST_CHAT_ADDR)
	ln, err := net.Listen("udp", MULTICAST_CHAT_ADDR)
	net.ListenMulticastUDP()
}
