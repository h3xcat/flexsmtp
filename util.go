package main

import (
	"bufio"
	"io"
	"strings"
)

func readLineCRLF(reader *bufio.Reader, limit int) (string, error) {
	var buffer strings.Builder
	var lastChar byte
	for {
		char, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				return buffer.String(), io.EOF
			}
			return "", err
		}

		if buffer.Len() == limit {
			return "", io.ErrShortBuffer
		}

		buffer.WriteByte(char)

		if lastChar == '\r' && char == '\n' {
			return buffer.String(), nil
		}
		lastChar = char
	}
}
