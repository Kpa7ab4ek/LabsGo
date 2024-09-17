package laba2

import (
	"fmt"
	"unicode/utf8"
)

func Laba21(a int) {

	if a%2 == 0 {
		fmt.Println("Число:", a, " чётное")
	} else {
		fmt.Println("Число:", a, " нечётное")
	}
}

func Laba22(a float64) {
	if a > 0 {
		fmt.Println("Число:", a, " положительное")
	} else if a < 0 {
		fmt.Println("Число:", a, " отрицательное")
	} else {
		fmt.Println("Число:", a, " ноль")
	}
}

func Laba23() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func Laba24(str string) int {
	a := utf8.RuneCountInString(str)
	return a;

	/*
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Удаляем символ новой строки, который добавляется при вводе
	input = input[:len(input)-1]
	*/
}


type Rectangle struct{
	Height float64
	Width float64
}

func Laba25(rectangle Rectangle) {
	fmt.Println("Итоговая плозадь: ",rectangle.Width * rectangle.Height)
}

func Laba26(a,b float64) float64{
	return (a+b)/2
}
