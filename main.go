// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-11-17
// Fileoverview: This program asks the user for their name and age 
// 			and then returns that information.

package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var userName string
	var ageAsString string
	var ageAsNumber int
	var ageFiveYearsAgo int
	
	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("What is your name? ")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Print("How old are you? ")
	ageAsString, _ = reader.ReadString('\n')
	ageAsString = strings.TrimSpace(ageAsString)

	// process
	ageAsNumber, _ = strconv.Atoi(ageAsString)
	ageFiveYearsAgo = ageAsNumber - 5

	// output
	fmt.Println()
	fmt.Printf("Hello, " + userName + "!\n")
	fmt.Printf("Five years ago, you were " + strconv.Itoa(ageFiveYearsAgo) + " years old.")

	fmt.Println("\nDone.")
}
