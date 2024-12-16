package t4

import "testing"

// Тест для перевірки мінімального значення серед трьох чисел
func TestFindMinimum(t *testing.T) {
    result := findMinimum(1, 2, 3)
    expected := float32(1)
    if result != expected {
        t.Errorf("Тест не пройдено. Результат - %f, а вірний - %f", result, expected)
    }
}

// Тест для перевірки середнього значення трьох чисел
func TestCalculateAverage(t *testing.T) {
    result := calculateAverage(1, 2, 3)
    expected := float32(2)
    if result != expected {
        t.Errorf("Тест не пройдено. Результат - %f, а вірний - %f", result, expected)
    }
}

// Тест для перевірки рівняння першого порядку: y = kx + b
func TestFirstOrderEquation(t *testing.T) {
    result := firstOrderEquation(1, 2, 3)
    expected := float32(5)
    if result != expected {
        t.Errorf("Тест не пройдено. Результат - %f, а вірний - %f", result, expected)
    }
}
