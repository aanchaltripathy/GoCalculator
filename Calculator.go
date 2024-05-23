package main

import "fmt"

func Add(a float64, b float64) float64 {

	return a + b
}
func Sub(a float64, b float64) float64 {
	if b > a {
		fmt.Println("Invalid input")
		return 0
	}
	return a - b
}
func Mul(a float64, b float64) float64 {
	return a * b
}
func Div(a float64, b float64) float64 {
	if b == 0 {
		fmt.Println("Invalid input")
		return 0
	}
	return a / b
}
func main() {
	var choice float64
	fmt.Println("--------------CALCULATOR---------------")
	fmt.Println("1. Addition")
	fmt.Println("2. Substraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("Enter choice")
	fmt.Scanln(&choice)
	var a, b float64
	switch choice {
	case 1:
		{

			fmt.Println("enter 1st number")
			fmt.Scanln(&a)
			fmt.Println("enter 2nd number")
			fmt.Scanln(&b)
			result := Add(a, b)
			fmt.Println("Result", result)
		}
	case 2:
		{
			fmt.Println("enter 1st number")
			fmt.Scanln(&a)
			fmt.Println("enter 2nd number")
			fmt.Scanln(&b)
			result := Sub(a, b)
			fmt.Println("Result", result)
		}
	case 3:
		{
			fmt.Println("enter 1st number")
			fmt.Scanln(&a)
			fmt.Println("enter 2nd number")
			fmt.Scanln(&b)
			result := Mul(a, b)
			fmt.Println("Result", result)
		}
	case 4:
		{
			fmt.Println("enter 1st number")
			fmt.Scanln(&a)
			fmt.Println("enter 2nd number")
			fmt.Scanln(&b)
			result := Div(a, b)
			fmt.Println("Result '", result, "'")
		}
	default:
		{
			fmt.Println("Enter valid choice")
		}
	}
}
