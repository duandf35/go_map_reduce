package main

import(
	"fmt"
	"os"
	"strconv"
	"unicode"
	"reduce/gen"
	"reduce/utils"
)

// TODO: How to throw exception ?

// Find max temperature of EACH city and write the result into file.
// Assumption: City is in the second column and the last column is the temperature
// @param src
// @return result
func findMax(src [][]string) [][]string {
	// Map to track max temperature of each city
	mapping := make(map[string]int)
	
	fmt.Printf("Finding highest temperature of each city...")

	// TODO: More flexibility
	for row := range src {
		// fmt.Printf("Processing %d row: %s, length: %d\n", row, src[row], len(src[row]))

		if row != 0 {
			city := src[row][1]
			if isValidCity(city) {
				rawTemperature, err := strconv.ParseInt(src[row][7], 10, 64)

				if nil != err {
					// fmt.Printf("Invalid temperature: %d, skip.\n", src[row][9])
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
				// fmt.Printf("Skip invalid city name: %s\n", city)
			}
		}
	}

	return copyResult(mapping)
}

// Copy result into [][]string map which can be picked up by the CSV Writer.
// @param mapping
// @return result
func copyResult(mapping map[string]int) [][]string {
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

// Validate city name. 
// City name is invalid if the string contains any digit rune
// @param city
// @return true if valid
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
	if (len(os.Args) < 3) {
		fmt.Println("Expect 2 arguments. Use case: '$GOPATH/bin/reduce [upperBound] [intputFile].csv'\n")
		return
	}

	// Get all arguments except the program name (the 1st argument)
	args := os.Args[1:]
	m, err := strconv.Atoi(args[0])
	input := utils.InputPath(args[1])
	
	if nil != err {
		fmt.Printf("Invalid argument: %s\n", err)
		return
	} 

	// Generate testing data
	gen.Do(m, input)

	genPath := utils.DefaultGenPath()
	resultPath := utils.DefaultResultPath()

	bucket := utils.ReadFromCSV(genPath)
	result := findMax(bucket)
	utils.WriteToCSV(result, resultPath)

	fmt.Printf("Result has written in path: %s\n", resultPath)
}
