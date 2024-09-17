package laba1
import (
	"fmt"
	"time"
)

// 1) Написать программу, которая выводит текущее время и дату.
func Laba11() {
	currentTime := time.Now()
	fmt.Println("Текущее время: ", currentTime.String())
}

// 2) Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
// 3) Использовать краткую форму объявления переменных для создания и вывода переменных.
func Laba123() {
	fmt.Println("-----------------------------------")
	var numInt1 int = 5
	numInt2 := -5
	fmt.Println("Целые числа: ", numInt1, numInt2)

	var numFloat1 float64 = 12.12
	numFloat2 := 22.22
	fmt.Println("Вещественные числа: ", numFloat1, numFloat2)

	var string1 string = "String"
	string2 := "Тоже стринг"
	fmt.Println("Строки: ", string1, string2)

	var bool1 bool = true
	bool2 := false
	fmt.Println("Логический тип данный: ", bool1, bool2)
	fmt.Println("-----------------------------------")
}

// 4) Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
func Laba14(a, b int) {
	fmt.Println("Итого:", a+b)
}

// 5) Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
func Laba15(a, b float64) {
	fmt.Println("Итого: ", a+b)
}

// 6) Написать программу, которая вычисляет среднее значение трех чисел.
func Laba16(a, b, c float64) {
	fmt.Printf("Среднее значение: %.2f\n", (a+b+c)/3)
}
