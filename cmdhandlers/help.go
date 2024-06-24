package cmdhandlers

import (
	"flexsmtp/types"
)

type HELPCommand struct{}

func (cmd HELPCommand) Execute(client *types.SMTPClient, arg string) {
	client.Writer.WriteString("214-This server supports the following commands:\r\n")
	client.Writer.WriteString("214 HELO EHLO RCPT DATA RSET MAIL QUIT HELP AUTH BDAT\r\n")
	client.Writer.Flush()
	return
}
