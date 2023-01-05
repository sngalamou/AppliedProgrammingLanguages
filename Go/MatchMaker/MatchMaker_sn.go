package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	intro()
	weightedScore := question1()
	fmt.Printf("\nYour compatability score so far is %d percent", (100 - weightedScore))
	weightedScore += question2()
	fmt.Printf("\nYour compatability score so far is %d percent", (100 - weightedScore))
	weightedScore += question3()
	fmt.Printf("\nYour compatability score so far is %d percent", (100 - weightedScore))
	weightedScore += question4()
	fmt.Printf("\nYour compatability score so far is %d percent", (100 - weightedScore))
	weightedScore += question5()
	outro(weightedScore)
}

func intro() {
	fmt.Println("")
	fmt.Println("Welcome to the Go MatchMaker")
	fmt.Println("In this program, you will respond with a number from 1-5")
	fmt.Println("1 for strongly disagree, 3 for neutral & 5 for strongly agree")
}

func question1() int {
	inputVar := ""
	trueLove := 5
	weight := 3
	fmt.Println("\nHow Much Do You Like R&B?")
	fmt.Scanln(&inputVar) 		// scans the input from the console
	cont := validate(inputVar)
	if cont == false {
		fmt.Println("You asnwer the question correctly, remember your response must be a number between 1-5")
		fmt.Scanln(&inputVar)	
	} 
	intVar, _ := strconv.Atoi(inputVar)                                                                                                          // convert the input from the console to an integer
	compatibilityScore := int(math.Abs(float64(intVar - trueLove)))                                                                              // complete absolute value equation
	weightedScore := (weight * compatibilityScore)                                                                                               // calculate the weighted question compatibility scores
	fmt.Printf("The question compatibility score is %d, and the weighted question compatibility score is %d", compatibilityScore, weightedScore) // displays the compatibility score and weighted compatibility score
	return (weightedScore)                                                                                                                       // return the weighted question compatibility scores
}

func question2() int {
	fmt.Println("")
	inputVar := ""
	trueLove := 4
	weight := 2
	fmt.Println("\nHow Much Do You Like Video Games?")
	fmt.Scanln(&inputVar) 		// scans the input from the console
	cont := validate(inputVar)
	if cont == false {
		fmt.Println("You asnwer the question correctly, remember your response must be a number between 1-5")
		fmt.Scanln(&inputVar)	
	} 
	intVar, _ := strconv.Atoi(inputVar)                                                                                                           // convert the input from the console to an integer
	compatibilityScore := int(math.Abs(float64(intVar - trueLove)))                                                                              // complete absolute value equation
	weightedScore := (weight * compatibilityScore)                                                                                               // calculate the weighted question compatibility scores
	fmt.Printf("The question compatibility score is %d, and the weighted question compatibility score is %d", compatibilityScore, weightedScore) // displays the compatibility score and weighted compatibility score
	return (weightedScore)                                                                                                                       // return the weighted question compatibility scores
}

func question3() int {
	fmt.Println("")
	inputVar := ""
	trueLove := 4
	weight := 3
	fmt.Println("\nHow Much Do You Like Exercising?")
	fmt.Scanln(&inputVar) 		// scans the input from the console
	cont := validate(inputVar)
	if cont == false {
		fmt.Println("You asnwer the question correctly, remember your response must be a number between 1-5")
		fmt.Scanln(&inputVar)	
	} 
	intVar, _ := strconv.Atoi(inputVar)                                                                                                           // convert the input from the console to an integer
	compatibilityScore := int(math.Abs(float64(intVar - trueLove)))                                                                              // complete absolute value equation
	weightedScore := (weight * compatibilityScore)                                                                                               // calculate the weighted question compatibility scores
	fmt.Printf("The question compatibility score is %d, and the weighted question compatibility score is %d", compatibilityScore, weightedScore) // displays the compatibility score and weighted compatibility score
	return (weightedScore)                                                                                                                       // return the weighted question compatibility scores
}

func question4() int {
	fmt.Println("")
	inputVar := ""
	trueLove := 5
	weight := 2
	fmt.Println("\nHow Much Do You Like Trying New Things?")
	fmt.Scanln(&inputVar) 		// scans the input from the console
	cont := validate(inputVar)
	if cont == false {
		fmt.Println("You asnwer the question correctly, remember your response must be a number between 1-5")
		fmt.Scanln(&inputVar)	
	} 
	intVar, _ := strconv.Atoi(inputVar)                                                                                                           // convert the input from the console to an integer
	compatibilityScore := int(math.Abs(float64(intVar - trueLove)))                                                                              // complete absolute value equation
	weightedScore := (weight * compatibilityScore)                                                                                               // calculate the weighted question compatibility scores
	fmt.Printf("The question compatibility score is %d, and the weighted question compatibility score is %d", compatibilityScore, weightedScore) // displays the compatibility score and weighted compatibility score
	return (weightedScore)                                                                                                                       // return the weighted question compatibility scores
}

func question5() int {
	fmt.Println("")
	inputVar := ""
	trueLove := 3
	weight := 2
	fmt.Println("\nHow Much Do You Like To Read?")
	fmt.Scanln(&inputVar) 		// scans the input from the console
	cont := validate(inputVar)
	if cont == false {
		fmt.Println("You asnwer the question correctly, remember your response must be a number between 1-5")
		fmt.Scanln(&inputVar)	
	} 
	intVar, _ := strconv.Atoi(inputVar)                                                                                                           // convert the input from the console to an integer
	compatibilityScore := int(math.Abs(float64(intVar - trueLove)))                                                                              // complete absolute value equation
	weightedScore := (weight * compatibilityScore)                                                                                               // calculate the weighted question compatibility scores
	fmt.Printf("The question compatibility score is %d, and the weighted question compatibility score is %d", compatibilityScore, weightedScore) // displays the compatibility score and weighted compatibility score
	return (weightedScore)                                                                                                                       // return the weighted question compatibility scores
}

func outro(a int) {
	fmt.Println("")
	compScore := (100 - a)
	fmt.Println("\nThank you for answering the questions!")
	fmt.Println("We will now asess your responses & determine your compatability")
	fmt.Printf("Your final compatability score is %d percent", compScore)
	if compScore >= 80 {
		fmt.Println("\nYou two are made for each other, love at first sight")
	} else if compScore >= 65 {
		fmt.Println("\nYou two are compatible but it will take effort to make this thing work")
	} else {
		fmt.Println("\nDon't even bother trying, there is no future together for you two")
	}
	fmt.Println("")
}

func validate(s string) bool {
	options := [5]string{"1", "2", "3", "4", "5"}
	for _, v := range options {
		if v == s {
			return true
		}
	}
	return false
}
