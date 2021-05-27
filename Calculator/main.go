package main

import "fmt"

func Add(a int, b int) int {
	sum := a + b
	return sum
}

func Subtract(a int, b int) int {
	subtract := a - b
	return subtract
}

func Multiply(a int, b int) int {
	multiply := a * b
	return multiply
}

func Divide(a int, b int) float64 {
	aTemp := float64(a)
	bTemp := float64(b)
	divide := aTemp / bTemp
	return divide
}

func addInput() {
	var a, b int
	fmt.Println("Please enter first integer a")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Please enter second integer b")
	fmt.Scanf("%d\n", &b)
	sum := Add(a, b)
	fmt.Printf("The sum is %d\n", sum)
}

func subtractInput() {
	var a, b int
	fmt.Println("Please enter first integer a")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Please enter second integer b")
	fmt.Scanf("%d\n", &b)
	subtract := Subtract(a, b)
	fmt.Printf("The subtraction is %d\n", subtract)
}

func multiplyInput() {
	var a, b int
	fmt.Println("Please enter first integer a")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Please enter second integer b")
	fmt.Scanf("%d\n", &b)
	mutiply := Multiply(a, b)
	fmt.Printf("The multiplication is %d\n", mutiply)
}

func divideInput() {
	var a, b int
	fmt.Println("Please enter first integer a")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Please enter second integer b")
	fmt.Scanf("%d\n", &b)
	divide := Divide(a, b)
	fmt.Printf("The division is %.2f\n", divide)
}

func main() {
	var mainInput int
	for {
		fmt.Println("1.) Add two integer")
		fmt.Println("2.) Subtract two integer")
		fmt.Println("3.) Multiply two integer")
		fmt.Println("4.) Divide two integer")
		fmt.Scanf("%d\n", &mainInput)

		switch mainInput {
		case 1:
			addInput()
		case 2:
			subtractInput()
		case 3:
			multiplyInput()
		case 4:
			divideInput()
		default:
			fmt.Println("Invalid Input!")
		}
	}
}
