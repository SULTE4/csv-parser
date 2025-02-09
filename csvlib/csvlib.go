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
	headLine    []string
	row         []string
	fieldNumber int
}

type fields []byte

// func nextQuote()	for i := range f {
// 		if f[i] == ' ' {
// 			continue
// 		}
// 		if f[i] == '"' {
// 			return i
// 		} else {
// 			return -1
// 		}
// 	}
// 	return -1
// // }
// func (f fields) DivideFields() ([]string, error) {
// 	flds := []string{}
// 	field := []byte{}
// 	inQuote := false
// 	afterComma := false

// 	for i := 0; i < len(f); i++ {
// 		elementnue
// 		}

// 		switch element {
// 		case '"':
// 			if nextQuote(f[i+1:]) != -1 {
// 				field = append(field, '"')
// 				afterComma = false
// 				i++
// 			} else {
// 				inQuote = !inQuote
// 			}
// 		case ',':
// 			if inQuote {
// 				field = append(field, element)
// 			}
//             else {
//                 if len(field) == 1 && field[0] == '"' {
//                     flds = append(flds, "")
//                 }else{
// 				    flds = append(flds, string(field))
//                 }
// 				afterComma = true
// 				field = nil
// 			}
// 		default:
// 			field = append(field, element)
// 			if afterComma && element == ' ' {
// 				continue
// 			}
// 			afterComma= false

// 		}
// 	}

// 	if afterComma {
// 		return nil, ErrFieldEmpty
// 	}

//     if len(field) == 1 && field[0] == '"' {
//         flds = append(flds, "")
// 	}else if len(string(field)) != 0 {
// 		flds = append(flds, string(field))
// 	}

// 	return flds, nil
// }

func (c *MyCSVParser) ReadLine(r io.Reader) (string, error) {
	var line fields
	buf := make([]byte, 1)
	inQuote := false
	hasContent := false

	for {
		count, err := r.Read(buf)
		if count > 0 {
			element := buf[0]

			if element == '\r' {
				continue
			}
			if element == '"' {
				inQuote = !inQuote
			}

			if !inQuote && element == '\n' {
				break
			}

			line = append(line, element)
			hasContent = true
		}

		if err == io.EOF {
			if inQuote {
				return "", ErrQuote
			}
			if !hasContent {
				return "", io.EOF
			}
			break
		} else if err != nil {
			return "", err
		}
	}
	tmp, err := line.DivideFields()
	if err != nil {
		return "", err
	}

	if len(c.headLine) == 0 {
		c.headLine = tmp
	}

	if len(c.headLine) != len(tmp) {
		return "", ErrFieldCount
	}

	c.row = tmp
	c.fieldNumber = len(c.row)
	return string(line), nil
}

func (c *MyCSVParser) GetNumberOfFields() int {
	return c.fieldNumber
}

func (c *MyCSVParser) GetField(n int) (string, error) {
	if n < 0 || n >= c.fieldNumber {
		return "", ErrFieldIndex
	}
	return c.row[n], nil
}
