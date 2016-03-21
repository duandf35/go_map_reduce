package utils

import(
	"encoding/csv"
	"fmt"
	"os"
)

// IMPORTANT: Make the first letter of function name UPPER CASE to EXPORT!

// Read data from csv file
// @param input path
// @return data
func ReadFromCSV(input string) [][]string {
	csvf, err := os.Open(input)
	var csvd [][]string

	if nil != err {
		fmt.Printf("Issue opening csv file: %s\n", err)
	} else {
		csvr := csv.NewReader(csvf)

		data, err := csvr.ReadAll()
		if (nil != err) {
			fmt.Printf("Issue reading csv file: %s\n", err)
		} else {
			csvd = data
		}
	}

	return csvd
}

// Write generated data into csv file.
// @param data
// @param output path
func WriteToCSV(data [][]string, output string) {
	csvf, err := os.Create(output)
	if nil != err {
		fmt.Printf("Issue creating csv file: %s\n", err)
		return
	}

	defer csvf.Close()

	csvw := csv.NewWriter(csvf)
	
	for _, e := range data {
		fmt.Printf("Writing %s\n", e)
	}
	
	// csv writer accept []string or [][]string
	err = csvw.WriteAll(data)
	if (nil != err) {
		fmt.Printf("Issue writing csv file: %s\n", err)
		return
	}
	
	csvw.Flush()
}
