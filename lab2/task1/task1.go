package main

import "fmt"

func main() {
    // Обчислюємо мінімальне значення
    resultMin := getMinValue(0.1, -7, 2)
    fmt.Println("Мінімальне значення:", resultMin)

    // Обчислюємо середнє значення
    avg := calculateMean(14, 102, 501)
    fmt.Println("Середнє значення:", avg)

    // Розв'язуємо рівняння першого порядку
    equation := firstOrderEquation(6, 10, 5)
    fmt.Println("Рівняння першого порядку: y=kx+b\nРозв'язок:", equation)
}

// Функція для знаходження мінімального значення з трьох
func getMinValue(val1, val2, val3 float32) float32 {
    var minVal float32

    if val1 < val2 && val1 < val3 {
        minVal = val1
    } else if val2 < val1 && val2 < val3 {
        minVal = val2
    } else {
        minVal = val3
    }

    return minVal
}

// Функція для обчислення середнього значення
func calculateMean(num1, num2, num3 float32) float32 {
    return (num1 + num2 + num3) / 3
}

// Функція для розв'язку рівняння першого порядку: y = kx + b
func firstOrderEquation(k, x, b float32) float32 {
    return k*x + b
}
