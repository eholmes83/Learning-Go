package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER to continue..."

func main() {
	// one way - declare then assign (two steps)
	//var firstNumber int
	//firstNumber = 2

	// another way - declare type and name and assign value (doing all this way)

	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

	
	// one step variable - declare name, assign value, and let Go infer the type
	// subtraction := 7
	
	
	// ----------- GUESS THE NUMBER GAME --------------	
	
	// display welcome/instructions
	fmt.Println("Welcome to the Guess the Number Game!")
	fmt.Println("=====================================")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer at the end
	answer = (firstNumber * secondNumber) - subtraction
	fmt.Println("The answer is", answer)
}

