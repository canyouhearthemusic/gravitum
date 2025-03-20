package httpserver

import (
	"net"
)

type Configuration func(*Server)

func Port(port string) Configuration {
	return func(s *Server) {
		s.address = net.JoinHostPort("", port)
	}
}
