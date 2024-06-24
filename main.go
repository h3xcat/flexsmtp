package main

import (
	"bufio"
	"flexsmtp/types"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func handleConnection(client *types.SMTPClient) {
	defer client.Conn.Close()

	client.Reader = bufio.NewReader(client.Conn)
	client.Writer = bufio.NewWriter(client.Conn)

	client.Writer.WriteString(fmt.Sprintf("220 %s FlexSMTP Service ready at %s\r\n", client.Server.Name, time.Now().Format(time.RFC1123Z)))
	client.Writer.Flush()
	for !client.HasQuit {
		line, err := readLineCRLF(client.Reader, 512)
		if err != nil {
			if err.Error() == "short buffer" {
				client.Writer.WriteString("500 Line too long\r\n")
				client.Writer.Flush()
				continue
			}
			fmt.Println("Error reading from connection:", err.Error())
			return
		}

		cmdargs := strings.SplitN(line, " ", 2)
		args := ""
		if len(cmdargs) > 1 {
			args = cmdargs[1]
		}
		commandName := strings.ToUpper(cmdargs[0])
		command := GetCommand(commandName)
		if command != nil {
			command.Execute(client, args)
		} else {
			client.Writer.WriteString("500 Unrecognized command\r\n")
			client.Writer.Flush()
		}
	}
}

func main() {
	var server types.SMTPServer
	var err error
	if hostname := os.Getenv("FLEXSMTP_HOSTNAME"); hostname != "" {
		server.Name = hostname
	} else {
		server.Name, err = os.Hostname()
		if err != nil {
			fmt.Println("Error getting Hostname:", err.Error())
			return
		}
	}
	if addr := os.Getenv("FLEXSMTP_BIND_ADDR"); addr != "" {
		server.Address = addr
	} else {
		server.Address = ":25"
	}

	fmt.Println("Server name is", server.Name)

	listener, err := net.Listen("tcp", server.Address)
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", server.Address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		client := types.SMTPClient{
			Conn:   conn,
			Server: &server,
		}
		go handleConnection(&client)
	}
}
