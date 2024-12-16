package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first) // false
	fmt.Println("first за замовчуванням false")
	fmt.Println("second = ", second) // false
	fmt.Println("second за замовчуванням false")
	fmt.Println("third  = ", third) // true
	fmt.Println("third задали значення true")
	fmt.Println("fourth = ", fourth) // false
	fmt.Println("fourth буде false,бо зазначили,що значення не може бути як у third,тобто true")
	fmt.Println("fifth  = ", fifth) // true
	fmt.Println("fifth задали значення true\n")

	fmt.Println("!true  = ", !true) // false
	fmt.Println("Значення не true - це завжди false")
	fmt.Println("!false = ", !false) // true
	fmt.Println("Значення не false- це завжди true\n")

	fmt.Println("true && true   = ", true && true) // true
	fmt.Println("true І true буде true")
	fmt.Println("true && false  = ", true && false) // false
	fmt.Println("true І false буде false")
	fmt.Println("false && false = ", false && false) // false
	fmt.Println("false І false буде false\n")

	fmt.Println("true || true   = ", true || true) // true
	fmt.Println("true АБО true буде true")
	fmt.Println("true || false  = ", true || false) // true
	fmt.Println("true АБО false буде true")
	fmt.Println("false || false = ", false || false) // false
	fmt.Println("false АБО false буде false\n")

	fmt.Println("2 < 3  = ", 2 < 3) // true
	fmt.Println("Тому що 2 меньше 3, тому true")
	fmt.Println("2 > 3  = ", 2 > 3) // false
	fmt.Println("Тому що 2 не може бути більше 3, тому false")
	fmt.Println("3 < 3  = ", 3 < 3) // false
	fmt.Println("Тому що 3 не може бути більше або меньше 3, тому false")
	fmt.Println("3 <= 3 = ", 3 <= 3) // true
	fmt.Println("Тому що 3 рівне 3, тому true")
	fmt.Println("3 > 3  = ", 3 > 3) // false
	fmt.Println("Тому що 3 не може бути більше або меньше 3, тому false")
	fmt.Println("3 >= 3 = ", 3 >= 3) // true
	fmt.Println("Тому що 3 рівне 3, тому true")
	fmt.Println("2 == 3 = ", 2 == 3) // false
	fmt.Println("Тому що 2 не рівне 3, тому false")
	fmt.Println("3 == 3 = ", 3 == 3) // true
	fmt.Println("Тому що 3 рівне 3, тому true")
	fmt.Println("2 != 3 = ", 2 != 3) // true
	fmt.Println("Тому що 2 не рівне 3, тому true")
	fmt.Println("3 != 3 = ", 3 != 3) // false
	fmt.Println("Тому що 3 рівне 3, тому false")
	//Задание.
	//1. Пояснити результати операцій
	// Пояснення вище
}
