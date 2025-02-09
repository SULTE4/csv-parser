package csvlib

type CsvError struct {
	Message string
}

func (e *CsvError) Error() string {
	return e.Message
}

var (
	ErrQuote      = &CsvError{Message: "excess or missing \" in quoted-field"}
	ErrFieldCount = &CsvError{Message: "wrong number of fields"}
	ErrFieldIndex = &CsvError{Message: "invalid field index"}
	ErrFieldEmpty = &CsvError{Message: "field is empty"}
)

func NewCsvError(message string) error {
	return &CsvError{Message: message}
}
