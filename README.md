# Go CSV Parser

## Table of Contents

- [Go CSV Parser](#go-csv-parser)
  - [Table of Contents](#table-of-contents)
  - [About ](#about-)
  - [Project Structure ](#project-structure-)
  - [Getting Started ](#getting-started-)
    - [Prerequisites](#prerequisites)
    - [Installing](#installing)
  - [Usage ](#usage-)

## About <a name = "about"></a>

This project is a lightweight CSV parser library written in Go, designed to handle CSV files with support for quoted fields, escaped quotes, and basic validation. It provides a simple interface for reading lines, accessing fields, and managing errors, making it suitable for applications needing custom CSV parsing without relying on heavy dependencies.

The parser ensures consistent field counts across rows, skips carriage returns for compatibility, and includes custom error types for issues like unmatched quotes or invalid indices. 

## Project Structure <a name = "project_structure"></a>

```
.
├── csvlib/
│   ├── csvlib.go   → Core parser logic (MyCSVParser, ReadLine, etc.)
│   └── error.go    → Custom error types (ErrQuote, ErrFieldCount, etc.)
├── example.csv     → Sample CSV file for testing
├── go.mod          → Go module definition
├── main.go         → Example usage of the parser
└── README.md       → This file
```

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

```
- Go (version 1.18 or later): https://go.dev/dl/
```

### Installing

```bash
git clone https://github.com/SULTE4/csv-parser.git
```

```bash
cd a-library-for-others
go mod tidy
```

Run the demo:

```bash
go run main.go
```

Output will show parsed fields from `example.csv`.

## Usage <a name = "usage"></a>

Import the `csvlib` package and use `MyCSVParser` to read CSV data.

```go
import (
    "fmt"
    "io"
    "os"
    "github.com/SULTE4/csv-parser/csvlib"
)

func main() {
    file, _ := os.Open("example.csv")
    defer file.Close()

    parser := &csvlib.MyCSVParser{}

    for {
        line, err := parser.ReadLine(file)
        if err == io.EOF { break }
        if err != nil { fmt.Println("Error:", err); return }

        for i := 0; i < parser.GetNumberOfFields(); i++ {
            field, _ := parser.GetField(i)
            fmt.Printf("Field %d: %s\n", i+1, field)
        }
    }
}
```
