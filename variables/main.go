package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

const prompt = "and press ENTER to continue..."

func main() {
	// one way - declare then assign (two steps)
	//var firstNumber int
	//firstNumber = 2

	// another way - declare type and name and assign value (doing all this way)

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = (firstNumber * secondNumber) - subtraction

	guessTheNumberGame(firstNumber, secondNumber, subtraction, answer)

	
	// one step variable - declare name, assign value, and let Go infer the type
	// subtraction := 7
}

func guessTheNumberGame(firstNumber, secondNumber, subtraction, answer int) {

	reader := bufio.NewReader(os.Stdin)
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
	fmt.Println("The answer is", answer)
}

