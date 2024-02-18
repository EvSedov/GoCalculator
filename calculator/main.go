package main

import (
	calculator "github.com/evsedov/GoCalculator/calculator/server"
)

func main() {
	calculator := &calculator.Calculator{}
	calculator.Start()
}
