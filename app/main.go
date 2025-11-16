package main

import calculations "calculator/pkg/calculator"

var flag = true

func main() {
	for flag {
		flag = calculations.RunCalculatorSession()
	}
}
