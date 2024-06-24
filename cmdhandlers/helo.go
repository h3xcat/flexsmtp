package cmdhandlers

import (
	"flexsmtp/patterns"
	"flexsmtp/types"
)

type HELOCommand struct{}

func (cmd HELOCommand) Execute(client *types.SMTPClient, arg string) {
	if !patterns.ValidateDomain(arg) {
		client.Writer.WriteString("501 Invalid domain\r\n")
		client.Writer.Flush()
		return
	}
	if len(arg) > 255 {
		client.Writer.WriteString("501 Domain name too long\r\n")
		client.Writer.Flush()
		return
	}
	client.Name = arg
	client.HasHelo = true
	client.Writer.WriteString("250 " + client.Server.Name + " greets " + client.Name + "\r\n")
	client.Writer.Flush()
}
