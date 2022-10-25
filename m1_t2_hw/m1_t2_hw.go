/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var dig int      //трехзначное число
	var firstNum int // первая цифра
	var secNum int   // вторая цифра
	var thirdNum int // третья цифра

	var invertDig string // реверсивная запись числа

	fmt.Println("Введите трехзначное целое число:")
	fmt.Scanf("%d\n", &dig)

	if dig > 99 && dig < 1000 {
		firstNum = dig / 100              	// сотни
		secNum = (dig - dig/100*100) / 10 	// десятки
		thirdNum = dig % 10               	//единицы

		invertDig = strconv.Itoa(thirdNum) + strconv.Itoa(secNum) + strconv.Itoa(firstNum)
		fmt.Println("Риверсивная запись числа: ", invertDig)
	} else {
		fmt.Println("Операция не поддерживается.")
	}
}
