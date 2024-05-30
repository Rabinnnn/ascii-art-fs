package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-fs/functions"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2{
		fmt.Println("Only two or one arguments are accepted after main.go.\nUsage1: go run . [STRING] [BANNER]\nEX1: go run . something standard\nUsage2: go run . [STRING]\nEX2: go run . something")
		return
	}

	input := os.Args[1] // Reading the argument entered
	// fmt.Println(input)

	// handling a case where an empty string or \n only has been entered as argument
	if input == "" {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	inputFile := "standard.txt"

	// reading standard.txt and handling it's error
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error openning", inputFile, err)
		return
	}
	var fileLine []string

	// slicing the file into an array of string based on new line
	if inputFile == "thinkertoy.txt" {
		fileLine = strings.Split(string(file), "\r\n")
	} else {
		fileLine = strings.Split(string(file), "\n")
	}

	// adding link to the file in the link variable
	link := ""
	switch inputFile {
	case "standard.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	case "shadow.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	case "thinkertoy.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	default:
		fmt.Println("The file", inputFile, "is not valid for this program")
		return
	}

	if len(fileLine) != 856 {
		fmt.Println("The file", inputFile, "has been tampered with, please use the version from ", link, "!!!")
		return
	}

	// functions.AsciiArt(input, string(file))
	fmt.Print(functions.AsciiArt(input, fileLine))
}
