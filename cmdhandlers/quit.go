package cmdhandlers

import (
	"flexsmtp/types"
)

type QUITCommand struct{}

func (cmd QUITCommand) Execute(client *types.SMTPClient, arg string) {
	if len(arg) > 0 {
		client.Writer.WriteString("501 Invalid arguments\r\n")
		client.Writer.Flush()
		return
	}
	client.Writer.WriteString("221 " + client.Server.Name + " Service closing transmission channel\r\n")
	client.Writer.Flush()
	client.HasQuit = true
}
