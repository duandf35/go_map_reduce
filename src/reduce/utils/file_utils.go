package utils

import (
	"os"
	"fmt"
)

const INPUT_PATH = "/input/"
const OUTPUT_PATH = "/output/"

// Get the default data input path
// @param fn
// @return dir
func InputPath(fn string) string {
	dir := os.Getenv("GOPATH") + INPUT_PATH + fn;
	fmt.Printf("Input path: %d\n", dir)

	return dir
}

// Get the default data output path
// @param fn
// @return dir
func OutputPath(fn string) string {
	dir := checkDir(os.Getenv("GOPATH") + OUTPUT_PATH) + fn
	fmt.Printf("Output path: %d\n", dir)

	return dir
}

// Create directory if not exists.
// @param dir
func checkDir(dir string) string {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	return dir
}