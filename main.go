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
	userName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	userName = strings.TrimSpace(userName)

	fmt.Print("How old are you? ")
	ageAsString, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading age:", err)
		return // Exit on error
	}
	ageAsString = strings.TrimSpace(ageAsString)

	// // process
	ageAsNumber, err = strconv.Atoi(ageAsString)
	if err != nil {
		fmt.Println("Error: Please enter a valid number for your age. Conversion failed.")
		return
	}
	ageFiveYearsAgo = ageAsNumber - 5

	// // output
	fmt.Println()
	fmt.Println("Hello, " + userName + "!")
	fmt.Println("You are " + strconv.Itoa(ageAsNumber) + "years old.")
	fmt.Printf("Five years ago, you were %d years old.\n", ageFiveYearsAgo)

	fmt.Println("\nDone.")
}
