package main

import (
	"fmt"
	"io"
	"os"

	"a-library-for-others/csvlib"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser csvlib.CSVParser = &csvlib.MyCSVParser{}
	// fmt.Println("Number of Fields: ", csvparser.GetNumberOfFields())
	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Println("=====================================")
		fmt.Println(line)

		for i := 0; i < csvparser.GetNumberOfFields(); i++ {
			field, err := csvparser.GetField(i)
			if err != nil {
				fmt.Println("Error to get field: ", err)
			} else {
				fmt.Printf("Field %d: %s\n", i+1, field)
			}
		}
		// fmt.Println("Number of Fields: ", csvparser.GetNumberOfFields())
		fmt.Println("=====================================")
	}
}
