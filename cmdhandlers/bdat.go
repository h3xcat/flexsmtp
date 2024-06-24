package cmdhandlers

import (
	"flexsmtp/types"
)

type BDATCommand struct{}

func (cmd BDATCommand) Execute(client *types.SMTPClient, arg string) {
	if !client.HasHelo {
		client.Writer.WriteString("503 Bad sequence of commands\r\n")
		client.Writer.Flush()
		return
	}
	// TODO: Implement BDAT command
}
