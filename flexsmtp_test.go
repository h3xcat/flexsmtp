package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestSMTPServer(t *testing.T) {
	os.Setenv("FLEXSMTP_BIND_ADDR", "127.0.0.1:2525")
	os.Setenv("FLEXSMTP_HOSTNAME", "smtp.example.com")
	go func() {
		main() // Start the SMTP server in a goroutine
	}()
	time.Sleep(1 * time.Second) // Give the server a second to start

	tests := []struct {
		input    string
		expected string
	}{
		{"HELO example.com\r\n", "250 smtp.example.com greets example.com\r\n"},
		{"MAIL FROM:<test@example.com>\r\n", "250 OK\r\n"},
		{"RCPT TO:<recipient@example.com>\r\n", "250 OK\r\n"},
		{"DATA\r\n", "354 End data with <CR><LF>.<CR><LF>\r\n"},
		{"QUIT\r\n", "221 smtp.example.com Service closing transmission channel\r\n"},
	}

	for _, test := range tests {
		log.Println("Running test:", test.input)
		conn, err := net.Dial("tcp", "127.0.0.1:2525")
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)

		var msg string
		msg, err = reader.ReadString('\n') // Read the welcome message
		if err != nil {
			t.Fatalf("Failed to read from server: %v", err)
		}
		t.Logf("S: %s", msg)
		_, err = writer.WriteString(test.input)
		if err != nil {
			t.Fatalf("Failed to write to server: %v", err)
		}
		writer.Flush()
		t.Logf("C: %s", test.input)

		response, err := reader.ReadString('\n')
		if err != nil {
			t.Fatalf("Failed to read from server: %v", err)
		}
		t.Logf("S: %s", response)
		if response != test.expected {
			t.Errorf("Expected response %q, got %q", test.expected, response)
		}

		conn.Close()
	}
}
