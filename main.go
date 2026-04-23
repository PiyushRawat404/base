package main


import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// func main(){
// 	var num1 float64
// 	var num2 float64
// 	var operator string

// 	fmt.Println("Enter the first number")
// 	fmt.Scan(&num1)
// 	fmt.Println("Enter the second Number")
// 	fmt.Scan(&num2)
// 	fmt.Println("Choose the oprator + - / *")
// 	fmt.Scan(&operator)

// 	switch operator{
// 	case "+": 
// 	fmt.Println(num1 + num2)
// 	case "-":
// 	fmt.Println(num1-num2)
// 	case "*":
// 	fmt.Println(num1*num2)
// 	case "/":
// 	if(num2!=0){
// 		fmt.Println(num1/num2)
// 	}else{
// 		fmt.Println("Cannot determine")
// 	}
// 	default:fmt.Println("choose from operator")
// 	}
	
// }
func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	if num > 0 {
		fmt.Println("Number is Positive")
	} else if num < 0 {
		fmt.Println("Number is Negative")
	} else {
		fmt.Println("Number is Zero")
	}

	if isPrime(num) {
		fmt.Println("Number is Prime number")
	} else {
		fmt.Println("Number is Not a prime number")
	}

}
