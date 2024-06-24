package types

type SMTPCommand interface {
	Execute(client *SMTPClient, arg string)
}
