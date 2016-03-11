package utils

import(
	"encoding/csv"
	"fmt"
	"os"
)

const INPUT_PATH = "/input/"
const OUTPUT_PATH = "/output/"

// IMPORTANT: Make the first letter of function name UPPER CASE to EXPORT!

// Read data from csv file
// @param fn
// @return data
func ReadFromCSV(fn string) [][]string {
	f := os.Getenv("GOPATH") + INPUT_PATH + fn;
	csvf, err := os.Open(f)
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
// @param fn
func WriteToCSV(data [][]string, fn string) {
	f := checkDir(os.Getenv("GOPATH") + OUTPUT_PATH) + fn
	
	csvf, err := os.Create(f)
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

// Create directory if not exists.
// @param dir
func checkDir(dir string) string {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	return dir
}
