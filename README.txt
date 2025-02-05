# README.txt

## A Library for Others

### Overview
This project is a CSV parsing library in Go that demonstrates key concepts such as interfaces, file handling, and error management. The library provides an interface for reading CSV files line by line, extracting specific fields, and handling errors gracefully.

### Project Structure
```
a-library-for-other/
│── csvlib/
│   │── error.go        # Defines custom error messages
│   │── csvlib.go       # Implements CSV parsing logic
│── main.go             # Demonstrates library usage
│── go.mod              # Go module file
│── example.csv         # Sample CSV file for testing
│── sample.csv          # Another test CSV file
```

### Features
- **ReadLine**: Reads a single line from a CSV file while handling different newline terminators.
- **GetField**: Retrieves the nth field from the last-read line, handling quoted fields correctly.
- **GetNumberOfFields**: Returns the total number of fields in the last-read line.
- **Error Handling**: Detects missing or excess quotes and incorrect field counts, returning appropriate errors.

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/a-library-for-other.git
   ```
2. Navigate to the project directory:
   ```sh
   cd a-library-for-other
   ```
3. Initialize Go modules:
   ```sh
   go mod tidy
   ```

### Usage
Example usage of the CSV parser:
```go
file, err := os.Open("example.csv")
if err != nil {
    log.Fatal("Error opening file:", err)
}
defer file.Close()

var parser CSVParser = YourCSVParser{}
for {
    line, err := parser.ReadLine(file)
    if err != nil {
        if err == io.EOF {
            break
        }
        log.Println("Error reading line:", err)
        continue
    }
    fmt.Println("Read line:", line)
}
```

### Notes
- Ensure the code follows `gofumpt` formatting.
- Only `io.Reader` is allowed for file reading.
- The program must not panic or exit unexpectedly.
- Test the library with various CSV formats to verify robustness.

### License
This project is licensed under the MIT License.

