package main

import (
	"ascii-art-fs/functions"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAsciiArtStandard(t *testing.T) {

	inputFile := "standard.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\n")

	//create a map containing the input text as key and the path to its expected output as value
	testCases := map[string]string{
		"hello":         "testCases/expectedOutput1.txt",
		"you & me":      "testCases/expectedOutput2.txt",
		"Hello\\nThere": "testCases/expectedOutput3.txt",
		"{Hello There}": "testCases/expectedOutput4.txt",
		"1Hello 2There": "testCases/expectedOutput5.txt",
	}

	for input, expectedFilePath := range testCases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input, fileLine)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}

}

func TestAsciiArtShadow(t *testing.T) {

	inputFile := "shadow.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\n")

	//create a map containing the input text as key and the path to its expected output as value
	testCases := map[string]string{
		"hello world":                "testCases/expectedOutput6.txt",
		"123":                        "testCases/expectedOutput7.txt",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ": "testCases/expectedOutput8.txt",
	}

	for input, expectedFilePath := range testCases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input, fileLine)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}

}

func TestAsciiArtThinkertoy(t *testing.T) {

	inputFile := "thinkertoy.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\r\n")

	//create a map containing the input text as key and the path to its expected output as value
	testCases := map[string]string{
		"nice 2 meet you": "testCases/expectedOutput9.txt",
		"/(\")":           "testCases/expectedOutput10.txt",
		"\"#$%&/()*+,-./": "testCases/expectedOutput11.txt",
		"It's Working":    "testCases/expectedOutput12.txt",
	}

	for input, expectedFilePath := range testCases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input, fileLine)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}

}
