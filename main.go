package main

import (
	"LabsGo/labs/laba1"
	"fmt"
)

func main() {
	var lab, task int

	fmt.Println("Введите номер лабораторной работы:")
	fmt.Scanln(&lab)

	fmt.Println("Введите номер задания в этой лабораторной работе:")
	fmt.Scanln(&task)

	switch lab {
	case 1:
		switch task {
		case 1:
			laba1.Laba11()
		case 2:
			laba1.Laba123()
		case 3:
			laba1.Laba123()
		case 4:
			var a, b int
			fmt.Println("Введите первое число")
			fmt.Scanln(&a)
			fmt.Println("Введите второе число")
			fmt.Scanln(&b)
			laba1.Laba14(a, b)
		case 5:
			var a, b float64
			fmt.Println("Введите первое число")
			fmt.Scanln(&a)
			fmt.Println("Введите второе число")
			fmt.Scanln(&b)
			laba1.Laba15(a, b)
		case 6:
			var a, b, c float64
			fmt.Println("Введите первое число")
			fmt.Scanln(&a)
			fmt.Println("Введите второе число")
			fmt.Scanln(&b)
			fmt.Println("Введите третье число")
			fmt.Scanln(&c)
			laba1.Laba16(a, b, c)
		default:
			fmt.Println("Неверный номер задания для первой лабораторной работы.")
		}
	/*case 2:
	switch task {
	case 1:
		laba2.Laba21()
	case 2:
		laba2.Laba22()
	case 3:
		laba2.Laba23()
	default:
		fmt.Println("Неверный номер задания для второй лабораторной работы.")
	}
	*/
	default:
		fmt.Println("Неверный номер лабораторной работы.")
	}
}
