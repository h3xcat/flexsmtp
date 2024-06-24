package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadLineCRLF(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		limit      int
		want       string
		wantErr    bool
		wantErrStr string
	}{
		{"Simple line", "Hello, World!\r\n", 20, "Hello, World!\r\n", false, ""},
		{"Exact limit", "Hello, World!\r\n", 15, "Hello, World!\r\n", false, ""},
		{"Over limit", "Hello, World!\r\n", 14, "", true, "short buffer"},
		{"No CRLF", "Hello, World!", 20, "Hello, World!", true, "EOF"},
		{"Empty input", "", 20, "", true, "EOF"},
		{"Only CRLF", "\r\n", 20, "\r\n", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			got, err := readLineCRLF(reader, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLineCRLF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantErrStr {
				t.Errorf("readLineCRLF() error = %v, wantErrStr %v", err, tt.wantErrStr)
				return
			}
			if got != tt.want {
				t.Errorf("readLineCRLF() = %v, want %v", got, tt.want)
			}
		})
	}
}
