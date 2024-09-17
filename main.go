package main

import (
	"LabsGo/labs/laba1"
	"LabsGo/labs/laba2"
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
	case 2:
	switch task {
	case 1:
		fmt.Println("Введите число для проверки!")
		var a int
		fmt.Scanln(&a)
		laba2.Laba21(a)
	case 2:
		fmt.Println("Введите число для проверки!")
		var a float64
		fmt.Scanln(&a)
		laba2.Laba22(a)
	case 3: 
	laba2.Laba23()
	case 4:
		fmt.Println("Введитес строку!")
		var str string
		fmt.Scanln(&str)
		fmt.Println("Кол-во символов: ", laba2.Laba24(str))
	case 5:
		var height, width float64
		fmt.Println("Введите высоту и ширину прямоугольника:")
		fmt.Scanln(&height, &width)

		rect := laba2.Rectangle{
			Height: height,
			Width:  width,
		}

		laba2.Laba25(rect)

	case 6: 
	 var a,b float64
	 fmt.Println("Введитес a и b!")
	 fmt.Scanln(&a,&b)
	 fmt.Println(laba2.Laba26(a,b))
	default:
		fmt.Println("Неверный номер задания для второй лабораторной работы.")
	}
	
	default:
		fmt.Println("Неверный номер лабораторной работы.")
	}
}
