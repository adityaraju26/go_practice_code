package main 

import "fmt"

func main(){
	//Intro Information on User
	fmt.Println("Welcome to my Quiz!")
	fmt.Printf("Please enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hi %v! Welcome to the game.\n", name)

	fmt.Printf("Enter your age: ")
	var age uint //uint indicates that the value must not be negative
	fmt.Scan(&age)
	
	if age >= 18{ //checks if you are old enough to play the game
		fmt.Printf("\nYou are old enough for this quiz!\n")
	}else{
		fmt.Printf("\nSorry, you are too young for this quiz.")
		return //breaks out of function if user is too young
	}

	score := 0 //counter to keep track of points
	num_questions_total := 0 //total number of questions 

	//Question 1
	fmt.Printf("How many planets are in the universe? ")
	var numOfPlanets uint
	fmt.Scan(&numOfPlanets)
	if numOfPlanets != 8{
		fmt.Println("Incorrect")
	}else{
		fmt.Println("Correct! Let's move onto the next question")
		score++
	}

	//Question 2
	fmt.Printf("Which graphics card is better. RTX 3080 or RTX 3090? ")
	var graphicsCard string
	var graphisCardNumber string
	fmt.Scan(&graphicsCard, &graphisCardNumber)
	if graphicsCard + " " + graphisCardNumber == "RTX 3090" || graphicsCard + " " + graphisCardNumber == "rtx 3090" {
		fmt.Println("Nice Job! Let's move onto the next question")
		score++
	}else{
		fmt.Println("Incorrect.")
	}

	//Question 3
	fmt.Printf("How many hands does a human have? ")
	var numOfHands uint
	fmt.Scan(&numOfHands)
	if numOfHands != 2{
		fmt.Println("Incorrect")
	}else{
		fmt.Println("Correct! Let's move onto the next question")
		score++
	}

	num_questions_total = 3
	fmt.Printf("\n\nYou scored %v out of %v correct!", score, num_questions_total)
	percnet := (score/num_questions_total) * 100
	fmt.Printf("\n You scored %v percent", percent)

}