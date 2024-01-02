package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

func sumFirstLastDigits(filePath string) int {
	totalSum := 0

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return totalSum
	}
	defer file.Close()

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error getting absolute file path:", err)
		return totalSum
	}
	fmt.Println("File Path:", absFilePath)

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		trimmedLine := strings.TrimFunc(string(line), func(r rune) bool {
			return !unicode.IsDigit(r)
		})

		if len(trimmedLine) > 0 {
			firstDigit := string(trimmedLine[0])
			lastDigit := string(trimmedLine[len(trimmedLine)-1])
			finalNum := firstDigit + lastDigit

			finalInt, err := strconv.Atoi(finalNum)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			totalSum += finalInt
		}
	}

	return totalSum
}

func main() {
	fmt.Print("Enter the path to your text file: ")
	var filePath string
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := sumFirstLastDigits(filePath)
	fmt.Println("Sum of first and last numbers:", result)
}
