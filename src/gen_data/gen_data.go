package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"math/rand"
	"strconv"
)

const OUT_PUT_PATH = "./output/"

// Generating n random non-negative pesudo-random integer between [0, m).
// @param n
// @param m
// @return []string
func gen_all(n int, m int) []string {
	// mp := make(map[string]string)

	// list := [n]string throws error: non-constant array bound
	list := make([]string, n)

	for i := 0; i < n; i++ {
		list[i] = strconv.Itoa(gen(m))
	}

	return list
}

// Generating one random non-negative pesudo-random integer between [0, m).
// @param m
// @return int
func gen(m int) int {
	r := rand.Intn(m)
	fmt.Printf("Generate new integer: %d\n", r)

	return r
}

// Write generated data into csv file.
// @param data
// @param fn
func write_to_csv(data []string, fn string) {
	fpath := OUT_PUT_PATH + fn
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		os.MkdirAll(OUT_PUT_PATH, 0777)
	}

	csvf, err := os.Create(fpath)
	if nil != err {
		fmt.Printf("Issue creating csv file: %s\n", err)
		return
	}

	defer csvf.Close()

	csvw := csv.NewWriter(csvf)
	
	// csv writer accept []string or [][]string
	for _, e := range data {
		fmt.Printf("Writing %s\n", e)
	}
	
	err = csvw.Write(data)
	if (nil != err) {
		fmt.Printf("Issue writing csv file: %s\n", err)
		return
	}
	
	csvw.Flush()
}

// Random data generator.
func main() {
	if (len(os.Args) < 4) {
		fmt.Println("Expect 3 arguments. Use case: '$GOPATH/bin/gen [numberOfRecords] [upperBound] [fileName]'\n")
		return
	}

	// Get all arguments except the program name (the 1st argument).
	args := os.Args[1:]
	n, err := strconv.ParseInt(args[0], 10, 64)
	m, err := strconv.ParseInt(args[1], 10, 64)
	fn := args[2]

	if nil != err {
		fmt.Printf("Invalid argument: %s\n", err)
	}
	
	fmt.Printf("Creating csv file: %s with %d random integer between [0, %d).\n", fn, n, m)
	write_to_csv(gen_all(int(n), int(m)), fn)
}