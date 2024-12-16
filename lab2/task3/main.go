package main

import "fmt"
import "lab2/task3/t3"

func main() {
	var min = t3.MinValue(0.1, -7, 2)
	fmt.Println("Мінімальне значення", min)
	var mean = t3.MeanValue(14, 102, 501)
	fmt.Println("Середнє значення", mean)
	var eq = t3.EquationOfFist(6, 10, 5)
	fmt.Println("Рівняння першого порядку: y=kx+b. \ny=", eq)
}
