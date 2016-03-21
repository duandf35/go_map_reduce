package main

import(
	"fmt"
	"os"
	"strconv"
	"unicode"
	"reduce/gen"
	"reduce/utils"
)

// Default result file name
const RESULT = "maxTemperature.csv"

// TODO: How to throw exception ?

// Find max temperature of EACH city and write the result into file.
// Assumption: City is in the second column and the last column is the temperature
// @param src
// @return result
func findMax(src [][]string) [][]string {
	// Map to track max temperature of each city
	mapping := make(map[string]int)
	
	// TODO: More flexibility
	for row := range src {
		// fmt.Printf("Processing %d row: %s, length: %d\n", row, src[row], len(src[row]))

		if row != 0 {
			city := src[row][1]
			if isValidCity(city) {
				rawTemperature, err := strconv.ParseInt(src[row][7], 10, 64)

				if nil != err {
					fmt.Printf("Invalid temperature: %d, skip.\n", src[row][9])
				} else {
					temperature := int(rawTemperature)
					existTemperature, ok := mapping[city]

					if ok {
						if existTemperature < temperature {
							// fmt.Printf("Find higher temperature: %d in city: %s\n", temperature, city)
							mapping[city] = temperature
						}
					} else {
						// fmt.Printf("Update temperature: %d in city: %s\n", temperature, city)
						mapping[city] = temperature
					}
				}
			} else {
				fmt.Printf("Skip invalid city name: %s\n", city)
			}
		}
	}

	// List to store final result
	result := make([][]string, 50)
	count := 0
	for k, v := range mapping {
		// fmt.Printf("Write result, city: %s, temperature: %d\n", k, v)

		if cap(result) == count {
			extension := make([][]string, cap(result) * 2)
			copy(extension, result)
			result = extension
		}

		result[count] = make([]string, 2)
		if count == 0 {
			result[count][0] = "City"
			result[count][1] = "Highest Temperature"
		} else {
			result[count][0] = k
			result[count][1] = strconv.Itoa(v) 
		}
		
		count++
	}
	
	return result
}

func isValidCity(city string) bool {
	isValid := true
	if _, err := strconv.Atoi(city); err == nil {
		isValid = false
	} else {
		for _, r := range city {
			if unicode.IsDigit(r) {
				isValid = false
				break
			}
		}
	}

	return isValid
}

// Find max temperature of EACH city and write the result into file.
func main() {
	if (len(os.Args) < 4) {
		fmt.Println("Expect 3 arguments. Use case: '$GOPATH/bin/reduce [upperBound] [intputFile].csv [outputFile].csv'\n")
		return
	}

	// Get all arguments except the program name (the 1st argument)
	args := os.Args[1:]
	m, err := strconv.Atoi(args[0])
	input := utils.InputPath(args[1])
	output := utils.OutputPath(args[2])

	if nil != err {
		fmt.Printf("Invalid argument: %s\n", err)
	}

	fmt.Printf("Generating data into %s ...\n", output)
	gen.Do(m, output, input)

	resultOutput := utils.OutputPath(RESULT)
	bucket := utils.ReadFromCSV(output)
	result := findMax(bucket)
	utils.WriteToCSV(result, resultOutput)

	fmt.Printf("Result has written in path: %s\n", resultOutput)
}