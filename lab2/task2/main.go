package main

import "fmt"

func main() {
	var min = minValue(0.1, -7, 2)
	fmt.Println("Мінімальне значення", min)
	var mean = meanValue(14, 102, 501)
	fmt.Println("Середнє значення", mean)
	var eq = equationOfFist(6, 10, 5)
	fmt.Println("Рівняння першого порядку: y=kx+b. \ny=", eq)
}
