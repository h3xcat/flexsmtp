package cmdhandlers

import (
	"flexsmtp/types"
)

type MAILCommand struct{}

func (cmd MAILCommand) Execute(client *types.SMTPClient, arg string) {
	if !client.HasHelo {
		client.Writer.WriteString("503 Bad sequence of commands\r\n")
		client.Writer.Flush()
		return
	}
}
