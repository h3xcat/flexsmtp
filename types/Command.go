package types

type Command interface {
	Execute(client *SMTPClient, arg string)
}
