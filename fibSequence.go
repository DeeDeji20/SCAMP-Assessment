package main

import(
	"fmt"
)

func main() {
	x:= fib(31)
	fmt.Println(x)
	}
	func fib(n int) int{
		if n <= 1{
			return 1
		}
		return fib(n - 1) + fib(n - 2)
	}

	//1.Create a function fib which takes in parameter 
	//	of type int and returns an Int
	//2.Using recurssion, return the fibonnacci 
	//	sequence's equivalent of the argument passed in 
