package main

import "fmt"

func main() {
	const UsdToEur = 0.92305
	const UsdToRub = 0.011346

	fmt.Println("EurToRub-", UsdToEur*UsdToRub)
}
