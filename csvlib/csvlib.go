package csvlib

import (
	"errors"
	"fmt"
	"io"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

type fields string

type MyCSVParser struct {
	line        fields
	rows        []string
	fieldNumber int
}

func (f fields) DivideFields() []string {
	// result := []string{}

	return nil
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

func (c *MyCSVParser) ReadLine(r io.Reader) (string, error) {
	var line fields
	buf := make([]byte, 1)
	inQuote := false

	for {
		count, err := r.Read(buf)

		if buf[0] == '\n' {
			fmt.Print("----------->")
		}
		fmt.Println(string(buf[0]))

		if count == 0 {
			if err == io.EOF && len(line) > 0 {
				return string(line), nil
			}
			return "", err
		}

		element := buf[0]

		if element == '"' {
			inQuote = !inQuote
		}

		if !inQuote && (element == '\n' || element == '\r') {
			if element == '\r' {
				count, err = r.Read(buf)
				if count > 0 && buf[0] != '\n' {
					line += fields(element)
				}
			}
			break
		}

		line += fields(element)

	}

	if inQuote {
		return "", ErrQuote
	}

	return "", nil
}

func (c *MyCSVParser) GetNumberOfFields() int {
	return 1
}

func (c *MyCSVParser) GetField(n int) (string, error) {
	return "", nil
}
