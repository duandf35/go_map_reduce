package gen_data

import(
	"fmt"
	"math/rand"
	"strconv"
	"find_max_t/utils"
)

// Generating random temperature between [0, max) for all cities.
// @param cities
// @param max
// @return [][]string
func genTmps(cities [][]string, max int) [][]string {
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

// Random data generator.
// @param max the max temperature
// @param output
// @param input
func Gen(max int, output string, input string) {
	// if (len(os.Args) < 4) {
	// 	fmt.Println("Expect 3 arguments. Use case: '$GOPATH/bin/gen_data [upperBound] [intputFile] [outputFile]'\n")
	// 	return
	// }

	// Get all arguments except the program name (the 1st argument)
	// args := os.Args[1:]
	// m, err := strconv.ParseInt(args[0], 10, 64)
	// input := args[1]
	// output := args[2]

	// if nil != err {
	// 	fmt.Printf("Invalid argument: %s\n", err)
	// }
	
	fmt.Printf("Creating csv file: %s from %s with temerature between [0, %d).\n", output, input, max)
	utils.WriteToCSV(genTmps(utils.ReadFromCSV(input), max), output)
}
