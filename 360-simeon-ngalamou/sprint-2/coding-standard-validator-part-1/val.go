package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/rivo/uniseg"
)

func main() {
	readMeVal()
	licenseVal()
	commentsVal()
	charactersVal()
}

func readMeVal() {
	path, _ := os.Getwd() // allows me to get the working path instead of having to type it in
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var pass int   // created a variable that allows me to print for a final verdict on whether or not the file was found
	for _, file := range files { // searches the file directory and compares the file names to see if there is a readme file
		if file.Name() == "README.md" {
			pass = 1	// this variable tracks if the "Readme.md" was found in the path
		} else if file.Name() == "Readme.md" {
			pass = 1
		} else {
			pass = 0
		}
	}
	if pass == 0 {
		fmt.Println("Your Code Passes The \"README.md\" Validation Test")
	} else if pass == 0 {
		fmt.Println("Your Code Fails The \"README.md\" Validation Test")
	}
}

func licenseVal() {
	path, _ := os.Getwd() // allows me to get the working path instead of having to type it in
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var pass int
	for _, file := range files { // searches the file directort and compares the file names to see if there is a license file
		if file.Name() == "LICENSE.md" {
			pass = 1
		} else if file.Name() == "License.md" {
			pass = 1
		} else {
			pass = 0
		}
	}
	if pass == 0 {
		fmt.Println("Your Code Passes The \"LICENSE.md\" Validation Test")
	} else if pass == 0 {
		fmt.Println("Your Code Fails The \"LICENSE.md\" Validation Test")
	}
}

func commentsVal() {
	body, err := ioutil.ReadFile("val.go") // reads the inserted file name and puts it in the 'body' variable
	if err != nil {                        // catches an error if there is one
		log.Fatalf("unable to read file: %v", err)
	}

	if strings.Contains(string(body), "//") { // searches the 'body' variable for the string '//' because that indicates that there are comments in the code
		fmt.Println("Your Code Passes The \"Source Code Must Be Commented\" Validation Test")
	} else {
		fmt.Println("Your Code Fails The \"Source Code Must Be Commented\" Validation Test")
	}
}

func charactersVal() {
	// 7.  Lines should be wrapped at 100 characters
	// 13. File names should be meaningful and consistently follow either a CamelCase or dash-case naming convention
	file, err := os.Open("val.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var pass int
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		gr :=uniseg.GraphemeClusterCount(text)
		if int(gr) <= 100 {
			pass = 1
		} else {
			pass = 0
			
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	if pass == 1 {
		fmt.Println("Your Code Passes The \"100 Characters Per Line\" Validation Test")
	} else if pass == 0 {
		fmt.Println("Your Code Fails The \"100 Characters Per Line\" Validation Test")
	}
}
