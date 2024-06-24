package main

import (
	"flexsmtp/cmdhandlers"
	"flexsmtp/types"
)

func GetCommand(commandName string) types.SMTPCommand {
	switch commandName {
	case "HELO":
		return cmdhandlers.HELOCommand{}
	case "EHLO":
		return cmdhandlers.EHLOCommand{}
	case "MAIL":
		return cmdhandlers.MAILCommand{}
	case "QUIT":
		return cmdhandlers.QUITCommand{}
	case "HELP":
		return cmdhandlers.HELPCommand{}
	case "RCPT":
		return cmdhandlers.RCPTCommand{}
	case "DATA":
		return cmdhandlers.DATACommand{}
	case "RSET":
		return cmdhandlers.RSETCommand{}
	case "AUTH":
		return cmdhandlers.AUTHCommand{}
	case "BDAT":
		return cmdhandlers.BDATCommand{}
	default:
		return nil
	}
}
