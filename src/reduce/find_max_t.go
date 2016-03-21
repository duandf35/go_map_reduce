package main

import(
	"fmt"
	"os"
	"strconv"
	"reduce/gen"
	"reduce/utils"
)

// Default result file name
const RESULT = "result.csv"

func do(input string, output string) {

}

func main() {
	if (len(os.Args) < 4) {
		fmt.Println("Expect 3 arguments. Use case: '$GOPATH/bin/reduce [upperBound] [intputFile].csv [outputFile].csv'\n")
		return
	}

	// Get all arguments except the program name (the 1st argument)
	args := os.Args[1:]
	m, err := strconv.ParseInt(args[0], 10, 64)
	input := utils.InputPath(args[1])
	output := utils.OutputPath(args[2])

	if nil != err {
		fmt.Printf("Invalid argument: %s\n", err)
	}

	gen.Do(int(m), output, input)
}