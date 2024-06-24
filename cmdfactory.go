package main

import (
	"flexsmtp/cmdhandlers"
	"flexsmtp/types"
)

func GetCommand(commandName string) types.Command {
	switch commandName {
	case "HELO":
		return cmdhandlers.HELOCommand{}
	case "EHLO":
		return cmdhandlers.EHLOCommand{}
	case "MAIL":
		return cmdhandlers.MAILCommand{}
	case "QUIT":
		return cmdhandlers.QUITCommand{}
	default:
		return nil
	}
}
