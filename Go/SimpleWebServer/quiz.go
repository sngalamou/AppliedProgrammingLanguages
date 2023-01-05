package main

import (
	"fmt"
	"os"
	"text/template"
	"io/ioutil"
	"log"
)

type Quiz struct {
	id string    
	q1 string 
	a1 string 
	/*Q2 string `json:"q2"`
	A2 string `json:"a2"`
	Q3 string `json:"q3"`
	A3 string `json:"a3"`*/
}


func main() {
	q := Quiz{"1", "What's 1+1","2"}
	body, _ := ioutil.ReadFile("quiz.tmpl")
	fmt.Println(body)
	tmpl, err := template.New("quiz").Parse(string(body))
	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(os.Stdout, q)
}
