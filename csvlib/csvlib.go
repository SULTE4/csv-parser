package csvlib

import (
	"io"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

type MyCSVParser struct {
	row         []string
	fieldNumber int
}

type fields []byte

func (f fields) DivideFields() []string {
	flds := []string{}
	field := []byte{}
	inQuote := false
	for i := 0; i < len(f); i++ {
		element := f[i]

		if element == '"' {
			if inQuote && i+1 < len(f) && f[i+1] == '"' {
				i++
			} else {
				inQuote = !inQuote
			}
		} else if element == ',' && !inQuote {
			flds = append(flds, string(field))
			field = []byte{}
		} else {
			field = append(field, element)
		}
	}

	flds = append(flds, string(field))

	return flds
}

func (c *MyCSVParser) ReadLine(r io.Reader) (string, error) {
	var line fields
	buf := make([]byte, 1)
	inQuote := false

	for {
		count, err := r.Read(buf)

		if count == 0 {
			if err == io.EOF && len(line) > 0 {
				c.row = line.DivideFields()
				c.fieldNumber = len(c.row)
				return string(line), nil
			}
			return "", err
		}

		element := buf[0]

		if element == '"' {
			inQuote = !inQuote
		}

		if !inQuote && element == '\n' {
			break
		}
		if !inQuote && element == '\r' {
			count, _ = r.Read(buf)
			if count > 0 && buf[0] != '\n' {
				line = append(line, element)
			}
			break
		}

		line = append(line, element)
	}

	if inQuote {
		return "", ErrQuote
	}

	c.row = line.DivideFields()
	c.fieldNumber = len(c.row)
	return string(line), nil
}

func (c *MyCSVParser) GetNumberOfFields() int {
	return c.fieldNumber
}

func (c *MyCSVParser) GetField(n int) (string, error) {
	if n < 0 || n >= c.fieldNumber {
		return "", ErrFieldCount
	}
	return c.row[n], nil
}
