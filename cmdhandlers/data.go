package cmdhandlers

import (
	"flexsmtp/types"
)

type DATACommand struct{}

func (cmd DATACommand) Execute(client *types.SMTPClient, arg string) {
	if !client.HasHelo {
		client.Writer.WriteString("503 Bad sequence of commands\r\n")
		client.Writer.Flush()
		return
	}
	// TODO: Implement BDAT command
}
