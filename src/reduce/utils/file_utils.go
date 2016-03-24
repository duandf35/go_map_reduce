package utils

import (
	"os"
	"fmt"
)

const INPUT_PATH = "/input/"
const OUTPUT_PATH = "/output/"

// Default generate data file name
const DEF_GEN = "gen.csv"

// Default result file name
const DEF_RESULT = "result.csv"

// Get the default data input path.
// @param fn
// @return dir
func InputPath(fn string) string {
	dir := os.Getenv("GOPATH") + INPUT_PATH + fn;
	fmt.Printf("Input path: %s\n", dir)

	return dir
}

// Get the default data output path.
// @param fn
// @return dir
func OutputPath(fn string) string {
	dir := checkDir(os.Getenv("GOPATH") + OUTPUT_PATH) + fn
	fmt.Printf("Output path: %s\n", dir)

	return dir
}

// Get default data generation path.
// @return dir
func DefaultGenPath() string {
	return OutputPath(DEF_GEN)
}

// Get default result path.
// @return dir
func DefaultResultPath() string {
	return OutputPath(DEF_RESULT)
}

// Create directory if not exists.
// @param dir
// @return dir
func checkDir(dir string) string {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	return dir
}