package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/rivo/uniseg"
	"unicode"
	"regexp"
	"reflect"
)

func main() {
	/*ReadMeVal()
	LicenseVal()
	CommentsVal()
	CharactersVal()
	FileNamingVal()
	TabVal()*/
	FuncNamingVal()
}

func ReadMeVal() {
	path, _ := os.Getwd() // allows me to get the working path instead of having to type it in
	files, err := ioutil.ReadDir(path)
	if err != nil {
		CompileVal("f")
		log.Fatal(err)
	}
	var pass int	// created a variable that allows me to print for a final verdict on whether or not the file was found
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

func LicenseVal() {
	path, _ := os.Getwd() // allows me to get the working path instead of having to type it in
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		CompileVal("f")
	}
	var pass int
	for _, file := range files { // searches the file directory and compares the file names to see if there is a license file
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
	
func CommentsVal() {
	body, err := ioutil.ReadFile("val.go") // reads the inserted file name and puts it in the 'body' variable
	if err != nil {						// catches an error if there is one
		log.Fatalf("unable to read file: %v", err)
		CompileVal("f")
	}

	if strings.Contains(string(body), "//") { // searches the 'body' variable for the string '//' because that indicates that there are comments in the code
		fmt.Println("Your Code Passes The \"Source Code Must Be Commented\" Validation Test")
	} else {
		fmt.Println("Your Code Fails The \"Source Code Must Be Commented\" Validation Test")
	}
}

func CharactersVal() {
	file, err := os.Open("val.go")
	if err != nil {
		CompileVal("f")
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

	

func CompileVal(err string) {
	if err == "f" {		// handle error
		fmt.Println("Your Code Fails The \"No Errors or Warnings\" Validation Test")
	  } else if err == "p" {			// success
		fmt.Println("Your Code Passes The \"No Errors or Warnings\" Validation Test")
	  }
}

func FileNamingVal() {
	path, _ := os.Getwd() // allows me to get the working path instead of having to type it in
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		CompileVal("f")
	}
	var pass int
	for _, file := range files { // searches the file directory and compares the file names to see if there is a license file
		x := file.Name()
		CamelCase:= func() {
			for _,r:=range x {			
				if unicode.IsUpper(r) {
					pass = 1
					break
				} else {
					pass = 0
					break
				}
			}
		}
		if strings.Contains(string(x), "-") {
			pass = 1
		}else {
			pass = 0
			CamelCase()
			}
	}
	if pass == 1 {
		fmt.Println("Your Code Passes The \"File Names Should Be CamelCase or Dash-Case\" Validation Test")
	} else if pass == 0 {
		fmt.Println("Your Code Fails The \"File Names Should Be CamelCase or Dash-Case\" Validation Test")
	}
}

func FuncNamingVal() {			// I couldnt figure out how to make this function work exactly how I wanted, it would test for #10
	body, err := ioutil.ReadFile("val.go") // reads the inserted file name and puts it in the 'body' variable
	if err != nil {						// catches an error if there is one
		log.Fatalf("unable to read file: %v", err)
		CompileVal("f")
	}

	res := strings.SplitAfter(string(body),"func ")
	a := regexp.MustCompile(`func `)
	y:=0
	for x:=1; x<=10; x++ {
		var z string
		z = string(res[x])
		fmt.Println(a.Split(res[x], -1))
		for _,r:=range z {	
			fmt.Println(reflect.TypeOf(r))
			if unicode.IsUpper(r) {
				//pass = 1
				break
			} else {
				//pass = 0
				break
			}
		}
		y+= 1
	}
}

func TabVal() {
	body, err := ioutil.ReadFile("val.go") // reads the inserted file name and puts it in the 'body' variable
	if err != nil {						// catches an error if there is one
		log.Fatalf("unable to read file: %v", err)
		CompileVal("f")
	}
	y := strings.Repeat(" ", 4)
	if strings.Contains(string(body), y) { // searches the 'body' variable for the string '//' because that indicates that there are comments in the code
		fmt.Println("Your Code Fails The \"Tabs Should Be Used To Indent\" Validation Test")
		CompileVal("f")
	} else {
		fmt.Println("Your Code Passes The \"Tabs Should Be Used To Indent\" Validation Test")
		CompileVal("p")
	}
}
