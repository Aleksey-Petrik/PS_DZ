package main

import "fmt"

func input() string {
	var line string
	fmt.Print("Введите значение: ")
	fmt.Scan(&line)
	return line
}

func calc(a, b, c float64) float64 {
	return 0
}

func main() {
	const UsdToEur = 0.92305
	const UsdToRub = 0.011346
	input()
	fmt.Println("EurToRub -", UsdToEur*UsdToRub)
}
