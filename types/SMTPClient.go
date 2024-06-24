package types

import (
	"bufio"
	"net"
)

type SMTPClient struct {
	Conn    net.Conn
	Reader  *bufio.Reader
	Writer  *bufio.Writer
	Name    string
	Server  *SMTPServer
	HasQuit bool
	HasHelo bool
}
