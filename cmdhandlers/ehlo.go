package cmdhandlers

import (
	"flexsmtp/patterns"
	"flexsmtp/types"
)

type EHLOCommand struct{}

func (cmd EHLOCommand) Execute(client *types.SMTPClient, arg string) {
	if arg[0] == '[' && arg[len(arg)-1] != ']' {
		if !patterns.ValidateAddressLiteral(arg) {
			client.Writer.WriteString("501 Invalid address literal\r\n")
			client.Writer.Flush()
			return
		}
	} else {
		if len(arg) > 255 {
			client.Writer.WriteString("501 Domain name too long\r\n")
			client.Writer.Flush()
			return
		}
		if !patterns.ValidateDomain(arg) {
			client.Writer.WriteString("501 Invalid domain\r\n")
			client.Writer.Flush()
			return
		}
	}
	client.Name = arg
	client.HasHelo = true
	client.Writer.WriteString("250 " + client.Server.Name + " greets " + client.Name + "\r\n")
	client.Writer.Flush()
}
