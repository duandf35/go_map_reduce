package gen

import(
	"fmt"
	"math/rand"
	"strconv"
	"reduce/utils"
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
		if row == 0 {
			cities[row] = append(cities[row], "Temperature")
		} else {
			cities[row] = append(cities[row], strconv.Itoa(genTmp(max)))
		}
	}

	return cities
}

// Generating one random non-negative pesudo-random integer between [0, m).
// @param m
// @return int
func genTmp(m int) int {
	r := rand.Intn(m)
	// fmt.Printf("Generate new integer: %d\n", r)

	return r
}

// Random data generator.
// @param max the max temperature
// @param input
func Do(max int, input string) {
	output := utils.DefaultGenPath()
	bucket := utils.ReadFromCSV(input)

	utils.WriteToCSV(genTmps(bucket, max), output)

	fmt.Printf("Creating csv file: %s from %s with temerature between [0, %d).\n", output, input, max)
}
