package csvlib

import (
	"errors"
	"io"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

type MyCSVParser struct {
	fieldNumber int
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

func (c *MyCSVParser) ReadLine(r io.Reader) (string, error) {
	var result string

	return result, nil
}

func (c *MyCSVParser) GetField(n int) (string, error) {
}
