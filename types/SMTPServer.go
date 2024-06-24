package types

import "net"

type SMTPServer struct {
	Address  string
	Name     string
	Listener net.Listener
}
