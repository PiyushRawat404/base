package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int  {
		i++
		return i

	}
}
	func add( num3 , num4 int ) int {
		return num3 + num4
	}
func main() {
	// var a, b int
	// var op string
	// fmt.Println("Enter the num1 ")
	// fmt.Scan(&a)
	// fmt.Println("Enter the operator")
	// fmt.Scan(&op)
	// fmt.Println("Enter the num2 ")
	// fmt.Scan(&b)
	// switch op {
	// case "+":
	// 	fmt.Println("num1 + num2 = ", a+b)
	// case "-":
	// 	fmt.Println("num1 - num2 = ", a-b)
	// case "*":
	// 	fmt.Println("num1 * num2 = ", a*b)
	// case "/":
	// 	fmt.Println("num1 / num2 = ", a/b)
	// case "%":
	// 	fmt.Println("num1 % num2 = ", a%b)
	// default:
	// 	fmt.Println("Enter valid numbers")

	// }

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextINTS := intSeq()
	fmt.Println(nextINTS())
	fmt.Println(nextINTS())
	fmt.Println(nextINTS())


 r := add(2,7)
 fmt.Println(r)


}
