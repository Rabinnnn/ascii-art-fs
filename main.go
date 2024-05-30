package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-fs/functions"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2{
		fmt.Println("Only one or two arguments are accepted after main.go.\nUsage1: go run . [STRING] [BANNER]\nEX1: go run . something standard\nUsage2: go run . [STRING]\nEX2: go run . something")
		return
	}

	StringInput := os.Args[1] // Reading the argument entered

	// handling a case where an empty string or \n only has been entered as argument
	if StringInput == "" {
		return
	} 

	BannerFile := "standard.txt"

	if len(os.Args) == 3{
		banner := strings.Replace(os.Args[2], ".txt", "", -1)
		if !(banner == "standard" || banner == "shadow" || banner == "thinkertoy"){
			fmt.Println("This banner name", os.Args[2],"is incorrect, the only acceptable banner names are standard, shadow or thinkertoy")
			return
		}
		BannerFile = banner + ".txt"
	}

	// reading standard.txt and handling it's error
	file, err := os.ReadFile(BannerFile)
	if err != nil {
		fmt.Println("Error openning", BannerFile, err)
		return
	}
	var fileLine []string

	// slicing the file into an array of string based on new line
	if BannerFile == "thinkertoy.txt" {
		fileLine = strings.Split(string(file), "\r\n")
	} else {
		fileLine = strings.Split(string(file), "\n")
	}

	// adding link to the file in the link variable
	link := ""
	switch BannerFile {
	case "standard.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	case "shadow.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	case "thinkertoy.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	default:
		fmt.Println("The file", BannerFile, "is not valid for this program")
		return
	}

	if len(fileLine) != 856 {
		fmt.Println("The file", BannerFile, "has been tampered with, please use the version from ", link, "!!!")
		return
	}

	fmt.Print(functions.AsciiArt(StringInput, fileLine))
}
