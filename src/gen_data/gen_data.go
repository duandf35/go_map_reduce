package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"math/rand"
	"strconv"
)

const INPUT_PATH = "./input/"
const OUTPUT_PATH = "./output/"

// Generating random temperature between [0, max) for all cities.
// @param cities
// @param max
// @return [][]string
func genAll(cities [][]string, max int) [][]string {
	// mp := make(map[string]string)

	// list := [n]string throws error: non-constant array bound
	// list := make([]string, n)

	for row := range cities {
		cities[row] = append(cities[row], strconv.Itoa(genTmp(max)))
	}

	return cities
}

// Generating one random non-negative pesudo-random integer between [0, m).
// @param m
// @return int
func genTmp(m int) int {
	r := rand.Intn(m)
	fmt.Printf("Generate new integer: %d\n", r)

	return r
}

func readFromCSV(fn string) [][]string {
	fpath := INPUT_PATH + fn;
	csvf, err := os.Open(fpath)
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
func writeToCSV(data [][]string, fn string) {
	fpath := OUTPUT_PATH + fn
	checkDir(fpath)

	csvf, err := os.Create(fpath)
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
func checkDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
}

// Random data generator.
func main() {
	if (len(os.Args) < 4) {
		fmt.Println("Expect 3 arguments. Use case: '$GOPATH/bin/gen_data [upperBound] [intputFile] [outputFile]'\n")
		return
	}

	// Get all arguments except the program name (the 1st argument)
	args := os.Args[1:]
	m, err := strconv.ParseInt(args[0], 10, 64)
	input := args[1]
	output := args[2]

	if nil != err {
		fmt.Printf("Invalid argument: %s\n", err)
	}
	
	fmt.Printf("Creating csv file: %s from %s with temerature between [0, %d).\n", output, input, m)
	writeToCSV(genAll(readFromCSV(input), int(m)), output)
}
